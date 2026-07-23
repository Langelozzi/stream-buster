# StreamBuster — Deployment Runbook

> Everything we currently know about how StreamBuster is deployed. Last verified: 2026-07-23.

## 1. Topology

Single **DigitalOcean droplet** — a shared box that also runs unrelated apps
(worlds-bucket-list, budget, etc.).

| | |
|---|---|
| **Host** | `146.190.120.67` |
| **Hostname** | `ubuntu-s-1vcpu-1gb-amd-sfo3-01` (1 vCPU / 1 GB, region `sfo3`) |
| **OS** | Ubuntu (kernel 6.11) |
| **Login** | `root@146.190.120.67`, SSH **key** auth (already trusted in `known_hosts`) |
| **Containers** | **None** — no Docker; processes run directly |
| **Git remote** | `github.com/Langelozzi/stream-buster` |

There is **no CamAppEngine involvement** for this box — `~/deploy.yaml` only manages
`m1` / pixel-game / qmetry, not StreamBuster.

```
                    ┌──────────────── nginx (TLS, certbot) ────────────────┐
 api.streambuster.xyz ─┼─▶ proxy_pass http://localhost:8080  (Go backend)   │
     streambuster.xyz ─┼─▶ static  /var/www/streambuster.xyz/.../dist       │
 dev.streambuster.xyz ─┼─▶ static  /var/www/dev.streambuster.xyz/.../dist   │
                    └───────────────────────────────────────────────────────┘
 Backend  : /root/stream-buster/backend  →  `go run main.go` in screen `goserver` (:8080)
 Database : Neon serverless Postgres (us-east-2)  [conn string in backend/.env]
```

## 2. The three clones on the server

The frontend and backend are deployed from **separate git clones**:

| Path | Role | Branch | Served by |
|---|---|---|---|
| `/root/stream-buster/backend` | Backend (running) | `development` | `go run` in screen `goserver` → nginx `api.streambuster.xyz` |
| `/var/www/streambuster.xyz/html/stream-buster/frontend/dist` | Prod frontend | `development` | nginx static |
| `/var/www/dev.streambuster.xyz/html/stream-buster/frontend/dist` | Dev frontend | `maintance` (sic) | nginx static |

> ⚠️ The backend clone carries a **long-standing uncommitted local edit to
> `middlewares/cors_middleware.go`**. Never `git reset --hard` or full-dir-rsync the
> backend — you'll wipe it. Only copy the specific files you changed.

## 3. Backend deployment

Prereqs on the box: **Go 1.23.2** is installed (`/usr/bin/go`).

1. **Stage** the changed/new `.go` files locally (tar preserves paths):
   ```bash
   cd backend
   tar -czf /tmp/sb-backend.tgz \
     middlewares/usage_tracking_middleware.go routes/router.go \
     utils/database/initialize_db.go  <...new files...>
   ```
   (macOS `tar` adds AppleDouble `._*` files — Go ignores them, but clean with
   `find . -name '._*' -delete` after extract.)
2. **Copy + compile-gate** (the running server is unaffected by editing source on disk,
   so build *before* restarting):
   ```bash
   scp /tmp/sb-backend.tgz root@146.190.120.67:/tmp/
   ssh root@146.190.120.67 '
     cd /root/stream-buster/backend
     # back up files you overwrite:
     cp -a middlewares/usage_tracking_middleware.go /tmp/sb-bak-$(date +%s)-usage.go   # etc.
     tar -xzf /tmp/sb-backend.tgz -C /root/stream-buster/backend
     find . -name "._*" -delete
     go build ./...           # ← must succeed before restarting
   '
   ```
3. **Restart** the process inside the `goserver` screen:
   ```bash
   ssh root@146.190.120.67 "
     screen -S goserver -X stuff \$'\003'                                   # Ctrl-C stops go run
     # wait for :8080 to free, then:
     screen -S goserver -X stuff \$'cd /root/stream-buster/backend && go run main.go\n'
   "
   ```
   Restart triggers `database.InitializeDb()` → GORM **AutoMigrate on the prod Neon DB**
   (additive; fatal-on-error, so a serving process means migration succeeded).
4. **Verify:** `curl -s -o /dev/null -w '%{http_code}' http://localhost:8080/api/v1/analytics/me`
   → `401` = new code live (`404` = old code still running).

> The backend is NOT a systemd service and NOT a compiled binary in prod — it's
> literally `go run main.go` in a detached `screen`. Restarting recompiles (~10–20 s
> of API downtime).

## 4. Frontend deployment

**node/npm are NOT installed on the server.** The build happens locally and the static
`dist/` is copied up. `dist` is **gitignored** (`.gitignore` line 28) — only
`dist/index.html` + `vite.svg` were force-added; `dist/assets/` (the hashed JS/CSS) is
untracked. **A git-pull frontend deploy therefore breaks** (index.html points at assets
that aren't in git). Always build locally and copy the whole `dist/`.

```bash
cd frontend
npm ci                                   # if deps changed
npx vite build                           # PROD  → dist/      (uses .env)
npx vite build --mode dev --outDir dist-dev --emptyOutDir   # DEV → dist-dev/ (uses .env.dev)

TS=$(date +%Y%m%d-%H%M%S)
P=/var/www/streambuster.xyz/html/stream-buster/frontend/dist
D=/var/www/dev.streambuster.xyz/html/stream-buster/frontend/dist

ssh root@146.190.120.67 "cp -a $P $P.bak-$TS"          # backup
rsync -az --delete -e "ssh" dist/     root@146.190.120.67:$P/
ssh root@146.190.120.67 "cp -a $D $D.bak-$TS"          # backup
rsync -az --delete -e "ssh" dist-dev/ root@146.190.120.67:$D/
```

`--delete` removes stale hashed assets. No nginx reload needed (static files).

## 5. Configuration / environment

**Frontend** (Vite, baked in at build time):
- `frontend/.env` → `VITE_API_URL=https://api.streambuster.xyz/api/v1` (prod build).
- `frontend/.env.dev` → same URL today; consumed by `vite build --mode dev` for the dev
  site. Repoint this to change what `dev.streambuster.xyz` talks to (there is only one
  backend, so a real split needs a second backend + `api.dev.streambuster.xyz` vhost).
- Local dev override (commented in `.env`): `http://localhost:8080/api/v1`.

**Backend** (`godotenv`, loaded from `backend/.env` at runtime):
- `DB_CONNECTION_STRING` — Neon serverless Postgres, region **us-east-2** (`…neon.tech/neondb`,
  `sslmode=require`). **Credentials live only in `backend/.env` on the server — not in git.**
- Backend listens on `:8080` (`router.Run(":8080")` in `main.go`).

## 6. Database & migrations

- **Neon serverless Postgres** (single DB, shared by prod). No migration tool — GORM
  `AutoMigrate` runs every backend start from `utils/database/initialize_db.go`
  (`InitializeDb`), plus post-deployment seeders (`runPostDeploymentScripts`:
  roles, admin user, endpoint records, etc.).
- Migrations are **additive** (create tables/columns; never drops). A failed migrate is
  `log.Fatalf` → the process won't serve, which is your signal.
- `main.go` also runs `CreateEndpointRecords` on boot, auto-registering any new routes
  into the `endpoints` table.

## 7. nginx / domains / TLS

- Config: `/etc/nginx/sites-enabled/{api.streambuster.xyz, streambuster.xyz, dev.streambuster.xyz}`.
- TLS via **Let's Encrypt / certbot** (`/etc/letsencrypt/live/<domain>/`). HTTP→HTTPS redirects.
- `api.streambuster.xyz` → `proxy_pass http://localhost:8080` with `X-Forwarded-*` headers.
- SPA vhosts use `try_files $uri /index.html`.
- Reload after config changes: `nginx -t && systemctl reload nginx`.

## 8. Verification (prove a real 200)

```bash
# login as the seeded admin (role 1) to get the token cookie, then hit the admin API
BASE=https://api.streambuster.xyz/api/v1
curl -s -c /tmp/ck -X POST $BASE/auth/login \
  -H 'Content-Type: application/x-www-form-urlencoded' \
  --data-urlencode 'email=admin@admin.com' --data-urlencode 'password=111'
curl -s -b /tmp/ck -o /dev/null -w '%{http_code}\n' $BASE/analytics/overview   # → 200
```
A `200` on `/analytics/overview` also confirms the `request_logs` table exists (its
window query hits it). Sites: `https://streambuster.xyz/` and `/dev.streambuster.xyz/`
should return `200` and reference the freshly-built `index-<hash>.js`.

> ⚠️ `admin@admin.com` / `111` is the **seeded default** from
> `backend/utils/database/post_deployment_functions/create_admin_user.go`. Rotate it in
> production.

## 9. Rollback

**Frontend** — backups sit beside each live dir:
```bash
rm -rf $P && mv $P.bak-<TS> $P        # prod   (same pattern for dev $D)
```
**Backend** — restore the overwritten files and drop the new ones, then restart:
```bash
cd /root/stream-buster/backend
git checkout -- middlewares/usage_tracking_middleware.go routes/router.go utils/database/initialize_db.go
rm -f controllers/analytics_controller.go daos/analytics_dao.go \
      daos/interfaces/analytics_dao_interface.go middlewares/admin_middleware.go \
      models/analytics.go models/request_log.go routes/api/v1/analytics_routes.go \
      services/analytics_service.go services/interfaces/analytics_service_interface.go \
      utils/dependency_injection/analytics_di.go
# restart goserver screen (see §3). The empty request_logs table is harmless;
# `DROP TABLE request_logs;` if you truly want it gone.
```

## 10. Gotchas

- **Shell is zsh** — an unquoted `$VAR` holding several `ssh -o` flags is **not**
  word-split; inline the flags or you'll get `keyword ... extra arguments`.
- **macOS `tar`/scp** emit AppleDouble `._*` files; Go ignores files starting with `.`/`_`
  but clean them up.
- **Analytics routes are not usage-tracked** (they live in the private group, not the
  usage-tracking group), so hitting them does not create `request_logs` rows.
- Only **one backend** (`:8080`) serves both `api.streambuster.xyz` and, indirectly, the
  dev frontend. There is no separate dev API.
- Other apps share this droplet (ports 3000 node, 8123 python, plus their own screens);
  don't disturb them.

## 11. App entry points

- Dashboard route (frontend): **`/dashboard`** (behind `PrivateRoute`). Admins (role 1)
  see the analytics bento view; everyone else sees their personal usage view.
- New analytics API: `GET /api/v1/analytics/overview?userId=` (admin-only) and
  `GET /api/v1/analytics/me` (self).
