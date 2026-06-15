# =============================================================================
#  Test Container for Vaious Go Versions
# =============================================================================
# Default minimum version
ARG VARIANT="1.25-alpine"

# -----------------------------------------------------------------------------
#  Main Stage
# -----------------------------------------------------------------------------
FROM golang:${VARIANT}

RUN apk add --no-cache \
    git \
    alpine-sdk \
    build-base

WORKDIR /workspaces

ENTRYPOINT ["sh", "-c", "go mod download && go test -race ./..."]
