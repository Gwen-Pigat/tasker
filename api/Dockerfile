# Stage 1: Build the Go application
FROM golang:1.23-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Install git and ssl certificates
RUN apk add --no-cache git ca-certificates

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the Go application
# CGO_ENABLED=0 is critical for running in static alpine/scratch environments
# -ldflags="-s -w" shrinks the binary size significantly by stripping debug information
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o server main.go

# Stage 2: Final minimal runtime image
FROM alpine:3.18

# Add certificates for secure external outgoing network requests
RUN apk add --no-cache ca-certificates tzdata

WORKDIR /app

# Copy the compiled binary from the builder stage
COPY --from=builder /app/server .

# Expose the API server port
EXPOSE 3000

# Set environment variables defaults
ENV PORT=3000

# Start the API server
CMD ["./server"]
