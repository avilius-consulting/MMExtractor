# STAGE 1: Build using standard Debian-based Go image (Includes GCC natively)
FROM golang:1.22 AS builder

WORKDIR /app

# Copy and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy application source code
COPY . .

# Compile binary with CGO enabled for SQLite
ENV CGO_ENABLED=1
RUN go build -o server cmd/server/main.go

# STAGE 2: Lightweight Debian runtime execution image (Matches glibc!)
FROM debian:bookworm-slim
WORKDIR /root/

# Install essential runtime certificates so HTTPS image downloads work
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

# Copy the binary from the builder stage
COPY --from=builder /app/server .

EXPOSE 8080

CMD ["./server"]
