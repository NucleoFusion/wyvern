# Stage 1: Build the Go binary
FROM golang:1.24 AS builder

# Set working directory
WORKDIR /app

# Copy go mod files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the binary for Linux
RUN CGO_ENABLED=0 GOOS=linux go build -o server .

# Stage 2: Run the binary
FROM alpine:latest

# Install certificates for HTTPS if needed
RUN apk --no-cache add ca-certificates

# Set working directory in container
WORKDIR /root/

# Copy the binary from builder stage
COPY --from=builder /app/server .

# Expose port (optional)
EXPOSE 3000

# Run the binary
CMD ["./server"]

