# Use the official Golang image to build the Go user-service
FROM golang:1.23-alpine

# Set the working directory inside the container
WORKDIR /user-service

# Copy go.mod and go.sum first (to cache dependencies)
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Install Air for hot reloading
RUN go install github.com/air-verse/air@latest

# Copy the entire source code
COPY  .air.toml /user-service/.air.toml
COPY  cmd/user-service /user-service/cmd/user-service
COPY  config/user-service /user-service/config/user-service
COPY  config/shared /user-service/config/shared
COPY  internal/user-service /user-service/internal/user-service
COPY  pkg/middlewares /user-service/pkg/middlewares

# Inject service name into the .air.toml file dynamically
RUN sed -i 's/\$SERVICE_NAME/user-service/' /user-service/.air.toml

# Expose the port the service listens on
EXPOSE 5001

# Run Air for hot reloading
CMD ["air", "-c", "/user-service/.air.toml"]
