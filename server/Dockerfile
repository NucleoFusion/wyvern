FROM golang:1.24

# Set working directory
WORKDIR /app

# Install dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source files (only for initial run; volume mount will override)
COPY . .

# Expose app port
EXPOSE 3000

# Run directly with go run
CMD ["go", "run", "."]

