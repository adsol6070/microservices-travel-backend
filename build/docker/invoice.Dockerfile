#=====================================
# Stage 1: Base Image (Common for Dev & Prod)
#=====================================
FROM golang:1.23-alpine AS base

# Set environment variables
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    SERVICE_NAME=blog-service

# Set the working directory
WORKDIR /invoice-service

# Install system dependencies
RUN apk add --no-cache curl bash git make

# Copy only necessary dependency files to cache layers efficiently
COPY go.mod go.sum ./
RUN go mod download

# Copy the Makefile into the container
COPY Makefile ./
COPY misc/make /invoice-service/misc/make

# Install tools (migrate, air) using the Makefile
RUN make deps

#=====================================
# Stage 2: Development (Hot Reloading)
#=====================================
FROM base AS dev

# Copy only service-specific files
COPY  .air.toml /invoice-service/.air.toml
COPY  cmd/invoice-service /invoice-service/cmd/invoice-service
COPY  config/shared /invoice-service/config/shared
COPY  internal/invoice-service /invoice-service/internal/invoice-service
COPY  pkg/middlewares /invoice-service/pkg/middlewares
COPY  pkg/utils /invoice-service/pkg/utils
COPY  pkg/logger /invoice-service/pkg/logger
COPY  pkg/response /invoice-service/pkg/response
COPY  pkg/validation /invoice-service/pkg/validation

# Inject service name into the .air.toml file dynamically
RUN sed -i 's/\$SERVICE_NAME/invoice-service/' /invoice-service/.air.toml

# Expose the port the service listens on
EXPOSE 8200

# Run Air for hot reloading
CMD ["bin/air", "-c", "/invoice-service/.air.toml"]

#=====================================
# Stage 3: Build Application (Production)
#=====================================
FROM base AS builder

# Copy only necessary service-specific files
COPY  .air.toml /invoice-service/.air.toml
COPY  cmd/invoice-service /invoice-service/cmd/invoice-service
COPY  config/shared /invoice-service/config/shared
COPY  internal/invoice-service /invoice-service/internal/invoice-service
COPY  pkg/middlewares /invoice-service/pkg/middlewares
COPY  pkg/utils /invoice-service/pkg/utils
COPY  pkg/logger /invoice-service/pkg/logger
COPY  pkg/response /invoice-service/pkg/response
COPY  pkg/validation /invoice-service/pkg/validation

# Compile the Go application
RUN go build -o /invoice-service/bin/invoice-service ./cmd/invoice-service

#=====================================
# Stage 4: Production Ready Image
#=====================================
FROM alpine:latest AS prod 

# Create a non-root user for security
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

# Set the working directory
WORKDIR /invoice-service

# Copy only the compiled binary from the builder stage
COPY --from=builder /invoice-service/bin/invoice-service /invoice-service/bin/invoice-service

# Ensure correct permissions
RUN chown -R appuser:appgroup /invoice-service

# Set non-root user
USER appuser

# Expose the application port
EXPOSE 8200

# Run the compiled binary
CMD ["/invoice-service/bin/invoice-service"]



