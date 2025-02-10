#=====================================
# Stage 1: Base Image (Common for Dev & Prod)
#=====================================
FROM golang:1.23-alpine AS base

# Set environment variables
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    SERVICE_NAME=hotel-booking-service

# Set the working directory
WORKDIR /hotel-booking-service

# Install system dependencies
RUN apk add --no-cache curl bash git make

# Copy only necessary dependency files to cache layers efficiently
COPY go.mod go.sum ./
RUN go mod download

# Copy the Makefile into the container
COPY Makefile ./
COPY misc/make /hotel-booking-service/misc/make

# Install tools (migrate, air) using the Makefile
RUN make deps

#=====================================
# Stage 2: Development (Hot Reloading)
#=====================================
FROM base AS dev

# Copy only service-specific files
COPY  .air.toml /hotel-booking-service/.air.toml
COPY  cmd/hotel-booking-service /hotel-booking-service/cmd/hotel-booking-service
COPY  config/hotel-booking /hotel-booking-service/config/hotel-booking
COPY  config/shared /hotel-booking-service/config/shared
COPY  internal/hotel-booking /hotel-booking-service/internal/hotel-booking
COPY  internal/shared /hotel-booking-service/internal/shared
COPY  pkg/middlewares /hotel-booking-service/pkg/middlewares

# Inject service name into the .air.toml file dynamically
RUN sed -i 's/\$SERVICE_NAME/hotel-booking-service/' /hotel-booking-service/.air.toml

# Expose the port the service listens on
EXPOSE 5100

# Run Air for hot reloading
CMD ["bin/air", "-c", "/hotel-booking-service/.air.toml"]

#=====================================
# Stage 3: Build Application (Production)
#=====================================
FROM base AS builder

# Copy only necessary service-specific files
COPY  cmd/hotel-booking-service /hotel-booking-service/cmd/hotel-booking-service
COPY  config/hotel-booking /hotel-booking-service/config/hotel-booking
COPY  config/shared /hotel-booking-service/config/shared
COPY  internal/hotel-booking /hotel-booking-service/internal/hotel-booking
COPY  internal/shared /hotel-booking-service/internal/shared
COPY  pkg/middlewares /hotel-booking-service/pkg/middlewares

# Compile the Go application
RUN go build -o /hotel-booking-service/bin/hotel-booking-service ./cmd/hotel-booking-service

#=====================================
# Stage 4: Production Ready Image
#=====================================
FROM alpine:latest AS prod 

# Create a non-root user for security
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

# Set the working directory
WORKDIR /hotel-booking-service

# Copy only the compiled binary from the builder stage
COPY --from=builder /hotel-booking-service/bin/hotel-booking-service /hotel-booking-service/bin/hotel-booking-service

# Ensure correct permissions
RUN chown -R appuser:appgroup /hotel-booking-service

# Set non-root user
USER appuser

# Expose the application port
EXPOSE 5100

# Run the compiled binary
CMD ["/hotel-booking-service/bin/hotel-booking-service"]



