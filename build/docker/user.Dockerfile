#=====================================
# Stage 1: Base Image (Common for Dev & Prod)
#=====================================
FROM golang:1.23-alpine AS base

# Set environment variables
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    SERVICE_NAME=user-service

# Set the working directory
WORKDIR /user-service

# Install system dependencies
RUN apk add --no-cache curl bash git make

# Copy only necessary dependency files to cache layers efficiently
COPY go.mod go.sum ./
RUN go mod download

# Copy the Makefile into the container
COPY Makefile ./
COPY misc/make /user-service/misc/make

# Install tools (migrate, air) using the Makefile
RUN make deps

#=====================================
# Stage 2: Development (Hot Reloading)
#=====================================
FROM base AS dev

# Copy only service-specific files
COPY  .air.toml /user-service/.air.toml
COPY  cmd/user-service /user-service/cmd/user-service
COPY  config/user-service /user-service/config/user-service
COPY  config/shared /user-service/config/shared
COPY  internal/user-service /user-service/internal/user-service
COPY  internal/shared/rabbitmq /user-service/internal/shared/rabbitmq
COPY  pkg/middlewares /user-service/pkg/middlewares

# Inject service name into the .air.toml file dynamically
RUN sed -i 's/\$SERVICE_NAME/user-service/' /user-service/.air.toml

# Expose the port the service listens on
EXPOSE 7100

# Run Air for hot reloading
CMD ["bin/air", "-c", "/user-service/.air.toml"]

#=====================================
# Stage 3: Build Application (Production)
#=====================================
FROM base AS builder

# Copy only necessary service-specific files
COPY  cmd/user-service /user-service/cmd/user-service
COPY  config/user-service /user-service/config/user-service
COPY  config/shared /user-service/config/shared
COPY  internal/user-service /user-service/internal/user-service
COPY  internal/shared/rabbitmq /user-service/internal/shared/rabbitmq
COPY  pkg/middlewares /user-service/pkg/middlewares
COPY  pkg/security /user-service/pkg/security

# Compile the Go application
RUN go build -o /user-service/bin/user-service ./cmd/user-service

#=====================================
# Stage 4: Production Ready Image
#=====================================
FROM alpine:latest AS prod 

# Create a non-root user for security
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

# Set the working directory
WORKDIR /user-service

# Copy only the compiled binary from the builder stage
COPY --from=builder /user-service/bin/user-service /user-service/bin/user-service

# Ensure correct permissions
RUN chown -R appuser:appgroup /user-service

# Set non-root user
USER appuser

# Expose the application port
EXPOSE 7100

# Run the compiled binary
CMD ["/user-service/bin/user-service"]



