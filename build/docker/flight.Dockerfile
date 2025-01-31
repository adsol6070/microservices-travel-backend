#=====================================
# Stage 1: Base Image (Common for Dev & Prod)
#=====================================
FROM golang:1.23-alpine AS base

# Set environment variables
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    SERVICE_NAME=flight-booking-service

# Set the working directory
WORKDIR /flight-booking-service

# Install system dependencies
RUN apk add --no-cache curl bash git make

# Copy only necessary dependency files to cache layers efficiently
COPY go.mod go.sum ./
RUN go mod download

# Copy the Makefile into the container
COPY Makefile ./
COPY misc/make /flight-booking-service/misc/make

# Install tools (migrate, air) using the Makefile
RUN make deps

#=====================================
# Stage 2: Development (Hot Reloading)
#=====================================
FROM base AS dev

# Copy only service-specific files
COPY  .air.toml /flight-booking-service/.air.toml
COPY  cmd/flight-booking-service /flight-booking-service/cmd/flight-booking-service
COPY  config/flight-booking /flight-booking-service/config/flight-booking
COPY  config/shared /flight-booking-service/config/shared
COPY  internal/flight-booking /flight-booking-service/internal/flight-booking
COPY  pkg/middlewares /flight-booking-service/pkg/middlewares

# Inject service name into the .air.toml file dynamically
RUN sed -i 's/\$SERVICE_NAME/flight-booking-service/' /flight-booking-service/.air.toml

# Expose the port the service listens on
EXPOSE 6100

# Run Air for hot reloading
CMD ["bin/air", "-c", "/flight-booking-service/.air.toml"]

#=====================================
# Stage 3: Build Application (Production)
#=====================================
FROM base AS builder

# Copy only necessary service-specific files
COPY  cmd/flight-booking-service /flight-booking-service/cmd/flight-booking-service
COPY  config/flight-booking /flight-booking-service/config/flight-booking
COPY  config/shared /flight-booking-service/config/shared
COPY  internal/flight-booking /flight-booking-service/internal/flight-booking
COPY  pkg/middlewares /flight-booking-service/pkg/middlewares

# Compile the Go application
RUN go build -o /flight-booking-service/bin/flight-booking-service ./cmd/flight-booking-service

#=====================================
# Stage 4: Production Ready Image
#=====================================
FROM alpine:latest AS prod 

# Create a non-root user for security
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

# Set the working directory
WORKDIR /flight-booking-service

# Copy only the compiled binary from the builder stage
COPY --from=builder /flight-booking-service/bin/flight-booking-service /flight-booking-service/bin/flight-booking-service

# Ensure correct permissions
RUN chown -R appuser:appgroup /flight-booking-service

# Set non-root user
USER appuser

# Expose the application port
EXPOSE 6100

# Run the compiled binary
CMD ["/flight-booking-service/bin/flight-booking-service"]