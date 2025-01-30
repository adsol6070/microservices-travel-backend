# Use the official Golang image to build the Go hotel-booking-service
FROM golang:1.23-alpine

# Set the working directory inside the container
WORKDIR /hotel-booking-service

# Copy go.mod and go.sum first (to cache dependencies)
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Install required system packages (curl, bash, git)
RUN apk add --no-cache curl bash git make

# Copy the Makefile into the container
COPY Makefile ./
COPY misc/make /hotel-booking-service/misc/make

# Install tools (migrate, air) using the Makefile
RUN make deps

# Copy the entire source code
COPY  .air.toml /hotel-booking-service/.air.toml
COPY  cmd/hotel-booking-service /hotel-booking-service/cmd/hotel-booking-service
COPY  config/hotel-booking /hotel-booking-service/config/hotel-booking
COPY  config/shared /hotel-booking-service/config/shared
COPY  internal/hotel-booking /hotel-booking-service/internal/hotel-booking
COPY  pkg/middlewares /hotel-booking-service/pkg/middlewares

# Inject service name into the .air.toml file dynamically
RUN sed -i 's/\$SERVICE_NAME/hotel-booking-service/' /hotel-booking-service/.air.toml

# Expose the port the service listens on
EXPOSE 5000

# Run Air for hot reloading
CMD ["bin/air", "-c", "/hotel-booking-service/.air.toml"]
