# Use the official Golang image to build the Go app
FROM golang:1.23-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules manifests
COPY go.mod go.sum ./

# Download Go modules
RUN go mod download

# Copy the source code
COPY . .

# Build the Go app
RUN GOOS=linux GOARCH=amd64 go build -o /go/bin/user-service ./cmd/user-service

# Start a new stage from a smaller image to run the service
FROM alpine:latest  

# Install CA certificates to avoid SSL issues
RUN apk --no-cache add ca-certificates

# Copy the pre-built binary from the previous stage
COPY --from=builder /go/bin/user-service /usr/local/bin/user-service

# Copy the configuration files correctly
COPY ./config/user-service /usr/local/bin/config/user-service
COPY ./config/shared /usr/local/bin/config/shared

# Expose the port the service will run on
EXPOSE 5001

# Command to run the executable
CMD ["/usr/local/bin/user-service"]
