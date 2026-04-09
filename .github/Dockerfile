# =============================================================================
#  Test Container for Vaious Go Versions
# =============================================================================
# Default minimum version
ARG VARIANT="1.24-alpine"

# -----------------------------------------------------------------------------
#  Main Stage
# -----------------------------------------------------------------------------
FROM golang:${VARIANT}

RUN apk add --no-cache \
    git \
    alpine-sdk \
    build-base

WORKDIR /workspaces

ENTRYPOINT go mod download && go test -race ./...
