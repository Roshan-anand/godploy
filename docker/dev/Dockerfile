# Stage 1: Frontend development with Bun in a single-package setup.
FROM --platform=$BUILDPLATFORM oven/bun:1.2.23-slim AS frontend-dev

ENV NODE_ENV=development

WORKDIR /app/frontend

EXPOSE 3000

CMD ["sh", "-c", "bun install --frozen-lockfile && bun run prepare && bun run dev --host 0.0.0.0 --port 3000"]

# Stage 2: Backend development image for the single backend module.
FROM --platform=$BUILDPLATFORM golang:1.26-trixie AS backend-dev

# Install development tools in a single layer
RUN apt-get update && apt-get install -y --no-install-recommends \
    ca-certificates \
    curl \
    git \
    tzdata && \
    rm -rf /var/lib/apt/lists/* && \
    go install github.com/air-verse/air@latest

WORKDIR /app

# Copy backend module files for dependency caching.
COPY backend/go.mod backend/go.sum ./backend/

# Download dependencies with cache mounts.
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    cd backend && go mod download && go mod verify

# Create runtime directories with proper permissions.
RUN mkdir -p /app/backend/data /app/backend/bin && \
    chmod 755 /app/backend/data /app/backend/bin

WORKDIR /app/backend

RUN mkdir -p frontend/dist && if [ ! -f frontend/dist/index.html ]; then printf '<!doctype html><html><body>dev placeholder</body></html>' > frontend/dist/index.html; fi;

EXPOSE 8080

ENTRYPOINT ["air"]
