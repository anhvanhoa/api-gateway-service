# Build stage
FROM golang:1.25-alpine AS builder

# Set working directory
WORKDIR /app

# Install git and ca-certificates (needed for go mod download)
RUN apk add --no-cache git ca-certificates

# Configure git for private repositories
RUN git config --global url."https://github.com/".insteadOf "git@github.com:" && \
    git config --global url."https://".insteadOf "git://"

# Set up authentication for private repositories using BuildKit secrets
ARG GOPRIVATE=github.com/anhvanhoa/*
ENV GOPRIVATE=${GOPRIVATE}
ENV GONOSUMDB=${GOPRIVATE}

# Prepare auth for private repos when available (BuildKit secret)
RUN --mount=type=secret,id=github_token \
    if [ -f /run/secrets/github_token ]; then \
        GITHUB_TOKEN=$(cat /run/secrets/github_token) && \
        printf "machine github.com\n  login %s\n  password x-oauth-basic\n" "$GITHUB_TOKEN" > ~/.netrc && \
        chmod 600 ~/.netrc; \
    fi

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN --mount=type=secret,id=github_token \
    if [ -f /run/secrets/github_token ]; then \
      GITHUB_TOKEN=$(cat /run/secrets/github_token); \
      printf "machine github.com\n  login %s\n  password x-oauth-basic\n" "$GITHUB_TOKEN" > ~/.netrc; \
      chmod 600 ~/.netrc; \
    fi; \
    go mod tidy

# Copy source code
COPY . .

# RUN go generate
# Ensure private module access also during generate
RUN --mount=type=secret,id=github_token \
    if [ -f /run/secrets/github_token ]; then \
      GITHUB_TOKEN=$(cat /run/secrets/github_token); \
      printf "machine github.com\n  login %s\n  password x-oauth-basic\n" "$GITHUB_TOKEN" > ~/.netrc; \
      chmod 600 ~/.netrc; \
    fi; \
    go run tools/generate.go

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./src/main.go

# Final stage
FROM alpine:latest

# Install ca-certificates for HTTPS requests
RUN apk --no-cache add ca-certificates

# Create non-root user
RUN adduser -D -s /bin/sh appuser

# Set working directory
WORKDIR /app

# Copy the binary from builder stage
COPY --from=builder /app/main .

# Copy swagger files
COPY --from=builder /app/swagger ./swagger

# Change ownership to appuser
RUN chown -R appuser:appuser /app

# Switch to non-root user
USER appuser

# Expose port
EXPOSE 8088

# Health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=5s --retries=3 \
    CMD wget --no-verbose --tries=1 --spider http://localhost:8088/health || exit 1

# Run the application
CMD ["./main"]
