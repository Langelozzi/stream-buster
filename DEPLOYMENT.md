# StreamBuster — Deployment Runbook

> How StreamBuster is deployed to the M1 via CamAppEngine.
> Last verified: 2026-07-24. The DigitalOcean droplet it replaces is documented
> in §9 and is still running as the rollback target.

## 1. Topology

Both halves run as Docker containers on a **Mac mini / MacBook (Apple Silicon)**,
deployed by **CamAppEngine**. The database did not move — it is still the same
Neon instance the droplet used.

| | |
|---|---|
| **Host** | `100.64.207.6` (Tailscale), `camerons-MacBook-Pro.local`, arm64 |
| **Login** | `camer@100.64.207.6`, key `~/.ssh/id_ed25519` |
| **Containers** | colima + Docker (`brew`), started by `homebrew.mxcl.colima` |
| **Reverse proxy** | Homebrew nginx on **:8080**, unprivileged, `homebrew.mxcl.nginx` |
| **Database** | Neon serverless Postgres, us-east-2 (unchanged) |
| **Deploy tool** | CamAppEngine, config `~/deploy.yaml`, project `stream-buster` |

```
   chickenwinsvm:443  ──▶  M1:8080  ──▶ nginx (Host-based vhosts)
                                          ├─ api.streambuster.xyz ─▶ 127.0.0.1:8790 ─▶ api  (Go, :8080)
                                          └─ streambuster.xyz     ─▶ 127.0.0.1:8791 ─▶ web  (nginx, :80)
                                                                                        │
   api ──────────────────────────────────────────────────────────────────────▶ Neon (us-east-2)
```

Both published ports bind **127.0.0.1 only** — nothing reaches a container
without passing through the host nginx first.

## 1a. Public ingress (edge VPS) — pending DNS

The M1 is Tailscale-only and its nginx is unprivileged, so it cannot bind
:80/:443. Public traffic arrives the same way `camfung.dev` does: through the
**edge VPS `chickenwinsvm`** at `178.128.151.171` (tailnet `100.118.152.28`),
which terminates TLS on a wildcard cert and proxies to `http://100.64.207.6:8080`
with `Host` preserved — the M1 does the per-host routing.

The proxy config is written and **staged, not active**, at
`/root/nginx-staged/tls-wildcard-streambuster-xyz.conf` on the edge. Installing
it before the cert exists makes `nginx -t` fail, which blocks every later reload
and would take `camfung.dev` down with it.

Remaining steps, in order:

1. **Move `streambuster.xyz` to Cloudflare.** It is still on Namecheap
   (`dns1/dns2.registrar-servers.com`); camfung.dev is on
   `sureena/harvey.ns.cloudflare.com`. Recreate the existing records in the new
   zone still pointing at the droplet — nothing changes for users yet. Confirm
   the API token in `/root/.secrets/cloudflare.ini` covers the new zone.
2. **Issue the wildcard cert** (DNS-01, so it works while DNS still points at
   the droplet — no downtime, no reachability requirement):
   ```bash
   certbot certonly --dns-cloudflare \
     --dns-cloudflare-credentials /root/.secrets/cloudflare.ini \
     -d streambuster.xyz -d '*.streambuster.xyz'
   ```
3. **Activate the proxy config:**
   ```bash
   mv /root/nginx-staged/tls-wildcard-streambuster-xyz.conf \
      /etc/nginx/camappengine/
   nginx -t && systemctl reload nginx
   ```
4. **Flip the A records** for `streambuster.xyz` and `api.streambuster.xyz` to
   `178.128.151.171`. Cert and route are already live, so the cutover is clean.

> ⚠️ **Leave `dev.streambuster.xyz` on the droplet.** The M1 has no route for it,
> and an unmatched `Host` there falls through to the stock "Welcome to nginx"
> page rather than 404ing — a silent wrong-content failure. The edge block
> matches `*.streambuster.xyz`, so only DNS keeps dev away from it.

Verified already, from the edge over the tailnet: `streambuster.xyz` → 200,
SPA fallback → 200, `api.streambuster.xyz` swagger → 200 and unauthenticated
→ 401. Only TLS and DNS remain.

## 2. How CamAppEngine deploys this

One command deploys both apps:

```bash
camappengine deploy -project stream-buster -env streambuster -server m1
```

Per app (`api`, then `web`) it:

1. **Syncs** `https://github.com/langelozzi/stream-buster` into
   `~/apps/stream-buster/<app>` — `git checkout development && git reset --hard
   origin/development`. **It deploys what is on GitHub, not your working tree.**
2. **Copies** that app's secrets file in as `.env` (see §3). Untracked, so the
   `reset --hard` on the next deploy leaves it alone.
3. **Runs** `docker compose up -d --build --remove-orphans` in the clone.
4. **Writes** `/opt/homebrew/etc/nginx/servers/<host>.conf` proxying to the
   app's port, runs `nginx -t`, and reloads. A failed config test removes the
   file rather than leaving a broken route.

Deploy one side only with `-app api` or `-app web`. Neither declares
`dependson`, so each deploys independently.

### One compose file, two apps

Both apps clone the *same repo*, so both get the same `docker-compose.yml`. The
copied-in `.env` decides which service actually starts:

| App | `COMPOSE_PROFILES` | `COMPOSE_PROJECT_NAME` | Host port |
|---|---|---|---|
| `api` | `api` | `streambuster-api` | 8790 |
| `web` | `web` | `streambuster-web` | 8791 |

Distinct project names keep the two stacks from adopting each other's
containers under `--remove-orphans`.

Run the whole stack locally with `COMPOSE_PROFILES=all docker compose up --build`.

## 3. Configuration and secrets

Secrets live **outside the repo**, on the machine that runs `camappengine`, and
are named in `~/deploy.yaml` as `secretsfile`:

| App | Local path |
|---|---|
| `api` | `~/deploy-secrets/stream-buster/api/.env` |
| `web` | `~/deploy-secrets/stream-buster/web/.env` |

`api/.env` carries the Compose wiring above plus everything `backend/.env` has
(`DB_CONNECTION_STRING`, `JWT_SECRET_TOKEN`, `DOMAIN`, the TMDB/CDN keys) and
one new variable:

- **`CORS_ALLOWED_ORIGINS`** — comma-separated exact origins. Credentialed
  routes rule out a wildcard, so every front end must be listed. Unset falls
  back to `http://localhost:5173`.

`web/.env` holds no secrets — it sets `VITE_API_URL`, which Vite **inlines at
image build time**. Changing the API URL requires a rebuild, not a restart.

Keep `api/.env` in sync with `backend/.env` when app config changes; nothing
enforces it.

## 4. Verification

```bash
M1=http://100.64.207.6:8080          # over Tailscale, before public ingress exists

# containers
ssh camer@100.64.207.6 'zsh -lc "docker ps --filter name=streambuster"'   # both (healthy)

# frontend
curl -s -o /dev/null -w '%{http_code}\n' -H 'Host: streambuster.xyz' $M1/          # 200
curl -s -o /dev/null -w '%{http_code}\n' -H 'Host: streambuster.xyz' $M1/watch/123 # 200, SPA fallback

# API: 401 -> login -> 200 with real data
H='Host: api.streambuster.xyz'
curl -s -o /dev/null -w '%{http_code}\n' -H "$H" $M1/api/v1/user/current           # 401
curl -s -c /tmp/ck -X POST -H "$H" -H 'Content-Type: application/x-www-form-urlencoded' \
  --data-urlencode 'email=admin@admin.com' --data-urlencode 'password=111' \
  $M1/api/v1/auth/login
curl -s -b /tmp/ck -o /dev/null -w '%{http_code}\n' -H "$H" $M1/api/v1/user/current # 200

# CORS
curl -s -i -X OPTIONS -H "$H" -H 'Origin: https://streambuster.xyz' \
  -H 'Access-Control-Request-Method: GET' $M1/api/v1/user/current | grep -i access-control
```

`rm /tmp/ck` afterwards — it holds a live session token.

> ⚠️ `admin@admin.com` / `111` is the seeded default from
> `backend/utils/database/post_deployment_functions/create_admin_user.go`.
> Rotate it.

## 5. Database and migrations

Unchanged by the move. GORM `AutoMigrate` runs on every backend start from
`utils/database/initialize_db.go`, followed by the idempotent post-deployment
seeders, and `main.go` registers new routes into the `endpoints` table.
Migrations are additive and `log.Fatalf` on failure — a container that reaches
`healthy` means the migration succeeded. The healthcheck allows a 40s
`start_period` for exactly this.

Because the droplet is still running against the **same Neon database**, both
backends share it during the transition. That is safe (additive migrations,
idempotent seeders) but means data written on one is immediately visible to the
other.

## 6. Rollback

DNS still points at the droplet, so during the transition rollback is "do
nothing". Once cut over:

- **Fast:** point `streambuster.xyz` / `api.streambuster.xyz` back at
  `146.190.120.67`. The droplet is untouched and still serving.
- **Per app on the M1:** `cd ~/apps/stream-buster/<app> && docker compose down`,
  then `rm /opt/homebrew/etc/nginx/servers/<host>.conf && nginx -s reload`.
- **Previous image:** deploys build in place; there is no image history to roll
  back to. Redeploy an older commit by pointing the environment's `branch` at
  it in `~/deploy.yaml` and deploying again.

## 7. Gotchas

- **Compose interpolates every service, including inactive profiles.** A `${VAR:?}`
  in the `web` service breaks the `api` deploy, whose `.env` has no `VITE_API_URL`.
  Build args need defaults, not required-markers.
- **Non-interactive ssh on macOS has no `/opt/homebrew/bin`.** CamAppEngine
  prefixes its own commands, but anything you run by hand needs `zsh -lc`.
- **nginx on the M1 listens on 8080, not 80** — it runs unprivileged as `camer`.
  Routing is by `Host`, so an edge proxy must preserve that header.
- **The deploy is from GitHub.** Uncommitted or unpushed work is invisible to it,
  and `development` is protected — changes land through a PR.
- **`go.sum` used to be gitignored.** Container builds start from a clean clone,
  so it must stay committed or `go mod download` has nothing to verify against.
- **`GetEnvVariable` used to `log.Fatalf` when `.env` was missing.** Containers
  ship no `.env`; it now warns once and reads the process environment.
- **Vite bakes `VITE_API_URL` into the bundle.** It is a build arg, not runtime
  config.
- Other apps share the M1 (`pixelizer` 8777, `qmetry` 8781, `demo` 3210, immich
  2283) — do not reuse those ports.

## 8. App entry points

- Frontend SPA at `/`, client-side routed (nginx `try_files` → `index.html`).
- API under `/api/v1`, Swagger UI at `/api/v1/swagger/index.html` — the only
  unauthenticated 200 the API serves, and what the container healthcheck uses.

## 9. Previous deployment (DigitalOcean droplet) — rollback target

Still running at **`146.190.120.67`** (`root@`, key auth), Ubuntu, no Docker:

- **Backend:** `/root/stream-buster/backend`, `go run main.go` in a `screen`
  named `goserver`, nginx proxying `api.streambuster.xyz` → `localhost:8080`.
  Restart with `screen -S goserver -X stuff $'\003'` then
  `stuff $'cd /root/stream-buster/backend && go run main.go\n'`.
- **Frontend:** node/npm are **not** installed there; `dist` was built locally
  and rsync'd to `/var/www/streambuster.xyz/html/stream-buster/frontend/dist`
  (dev site under `/var/www/dev.streambuster.xyz/...`).
- **TLS:** Let's Encrypt / certbot, configs in `/etc/nginx/sites-enabled/`.
- It carries a **long-standing uncommitted edit to
  `middlewares/cors_middleware.go`** — the origin list that made production
  work. That list is now `CORS_ALLOWED_ORIGINS` (§3), so the edit is obsolete,
  but do not `git reset --hard` there while the box is still the fallback.
- The box also runs unrelated apps (worlds-bucket-list, budget); leave them be.
