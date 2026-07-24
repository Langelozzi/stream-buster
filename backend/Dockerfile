# syntax=docker/dockerfile:1

# ---- build ----------------------------------------------------------------
FROM golang:1.23-alpine AS build

WORKDIR /src

# Its own layer so editing source doesn't refetch the module cache.
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Static binary: the runtime stage has no libc beyond musl and nothing here
# needs cgo.
RUN CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o /out/server .

# ---- run ------------------------------------------------------------------
FROM alpine:3.20

# ca-certificates: outbound TLS to Neon, TMDB and the CDN APIs.
# tzdata: analytics windows bucket by wall-clock time.
RUN apk add --no-cache ca-certificates tzdata

WORKDIR /app
COPY --from=build /out/server /app/server

# Config arrives as real environment variables (see docker-compose.yml); there
# is deliberately no .env in the image.
EXPOSE 8080
ENTRYPOINT ["/app/server"]
