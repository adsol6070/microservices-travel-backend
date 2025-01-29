# Use the official Golang image to build the Go app
FROM golang:1.23-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules manifests
COPY go.mod go.sum ./

# Download Go modules
RUN go mod download

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

# Copy the configuration files correctly
COPY ./config/flight-booking /usr/local/bin/config/flight-booking
COPY ./config/shared /usr/local/bin/config/shared

# Expose the port the service will run on
EXPOSE 9090

# Command to run the executable
CMD ["/usr/local/bin/flight-booking-service"]
