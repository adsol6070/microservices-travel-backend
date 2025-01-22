# Use the official Golang image as the build environment
FROM golang:1.20-alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the Go Modules and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod tidy

# Copy the source code into the container
COPY . .

# Build the Go app
RUN GOOS=linux GOARCH=amd64 go build -o /go/bin/flight-booking-service ./cmd/flight-booking-service

# Start a new stage from a smaller image to run the service
FROM alpine:latest  

# Install CA certificates to avoid SSL issues
RUN apk --no-cache add ca-certificates

# Copy the pre-built binary from the previous stage
COPY --from=builder /go/bin/flight-booking-service /usr/local/bin/flight-booking-service

# Expose the port the service will run on
EXPOSE 9090

# Command to run the executable
CMD ["flight-booking-service"]
