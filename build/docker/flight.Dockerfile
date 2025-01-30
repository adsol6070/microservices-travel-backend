# Use the official Golang image to build the Go flight-booking-service
FROM golang:1.23-alpine

# Set the working directory inside the container
WORKDIR /flight-booking-service

# Copy go.mod and go.sum first (to cache dependencies)
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Install required system packages (curl, bash, git)
RUN apk add --no-cache curl bash git make

# Copy the Makefile into the container
COPY Makefile ./
COPY misc/make /flight-booking-service/misc/make

# Install tools (migrate, air) using the Makefile
RUN make deps

# Copy the entire source code
COPY  .air.toml /flight-booking-service/.air.toml
COPY  cmd/flight-booking-service /flight-booking-service/cmd/flight-booking-service
COPY  config/flight-booking /flight-booking-service/config/flight-booking
COPY  config/shared /flight-booking-service/config/shared
COPY  internal/flight-booking /flight-booking-service/internal/flight-booking
COPY  pkg/middlewares /flight-booking-service/pkg/middlewares

# Inject service name into the .air.toml file dynamically
RUN sed -i 's/\$SERVICE_NAME/flight-booking-service/' /flight-booking-service/.air.toml

# Expose the port the service listens on
EXPOSE 9090

# Run Air for hot reloading
CMD ["bin/air", "-c", "/flight-booking-service/.air.toml"]
