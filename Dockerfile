# Build stage
FROM golang:1.23-alpine AS builder

# Install ca-certificates for HTTPS
RUN apk add --no-cache ca-certificates

# Set working directory
WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o hello-go-api .

# Final stage
FROM scratch

# Copy ca-certificates from builder
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Copy the binary from builder
COPY --from=builder /app/hello-go-api /hello-go-api

# Expose port
EXPOSE 8888

# Run the binary
ENTRYPOINT ["/hello-go-api"]
