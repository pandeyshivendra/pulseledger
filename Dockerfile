# ---- Build Stage ----
FROM golang:1.24.2-alpine AS builder

# Install CGO build dependencies (required by go-sqlite3)
RUN apk add --no-cache gcc musl-dev

WORKDIR /app

# Cache dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the Go app
RUN go build -o pulseledger main.go

# ---- Run Stage ----
FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/pulseledger .

# Expose the default Fiber port
EXPOSE 8080

# Command to run the executable
ENTRYPOINT ["./pulseledger"]