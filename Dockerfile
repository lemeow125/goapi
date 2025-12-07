# Stage 1: Build the Go binary
FROM golang:alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum (if exists) for dependency caching
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# Copy source code
COPY . .

# Build the application with optimizations
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build \
    -ldflags="-w -s" \
    -trimpath \
    -o /app/main \
    ./cmd

# Stage 2: Create minimal runtime image
FROM alpine:latest

RUN apk add --no-cache ca-certificates tzdata

# Create non-root user for security
RUN addgroup -g 1001 -S appgroup && \
    adduser -u 1001 -S appuser -G appgroup

WORKDIR /app

# Copy the binary from builder stage
COPY --from=builder --chown=appuser:appgroup /app/main /app/main

# Switch to non-root user
USER appuser

# Run the binary
ENTRYPOINT ["/app/main"]