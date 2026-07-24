# syntax=docker/dockerfile:1

# ---- build ----------------------------------------------------------------
FROM node:20-alpine AS build

WORKDIR /app

COPY package.json package-lock.json ./
RUN npm ci

COPY . .

# Vite inlines this at build time, so it is baked into the bundle and cannot
# be changed by the running container — a different API URL needs a rebuild.
# .env is gitignored, so the value has to come in as a build arg.
ARG VITE_API_URL
RUN test -n "$VITE_API_URL" \
    || { echo "VITE_API_URL build arg is required" >&2; exit 1; } \
 && printf 'VITE_API_URL=%s\n' "$VITE_API_URL" > .env.production \
 && npm run build

# ---- serve ----------------------------------------------------------------
FROM nginx:alpine

COPY nginx.conf /etc/nginx/conf.d/default.conf
COPY --from=build /app/dist /usr/share/nginx/html

EXPOSE 80
