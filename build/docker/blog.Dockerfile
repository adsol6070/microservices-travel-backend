#=====================================
# Stage 1: Base Image (Common for Dev & Prod)
#=====================================
FROM golang:1.23-alpine AS base

# Set environment variables
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    SERVICE_NAME=blog-service

# Set the working directory
WORKDIR /blog-service

# Install system dependencies
RUN apk add --no-cache curl bash git make

# Copy only necessary dependency files to cache layers efficiently
COPY go.mod go.sum ./
RUN go mod download

# Copy the Makefile into the container
COPY Makefile ./
COPY misc/make /blog-service/misc/make

# Install tools (migrate, air) using the Makefile
RUN make deps

#=====================================
# Stage 2: Development (Hot Reloading)
#=====================================
FROM base AS dev

# Copy only service-specific files
COPY  .air.toml /blog-service/.air.toml
COPY  cmd/blog-service /blog-service/cmd/blog-service
COPY  config/shared /user-service/config/shared
COPY  internal/blog-service /blog-service/internal/blog-service
COPY  pkg/middlewares /blog-service/pkg/middlewares

# Inject service name into the .air.toml file dynamically
RUN sed -i 's/\$SERVICE_NAME/blog-service/' /blog-service/.air.toml

# Expose the port the service listens on
EXPOSE 7200

# Run Air for hot reloading
CMD ["bin/air", "-c", "/blog-service/.air.toml"]

#=====================================
# Stage 3: Build Application (Production)
#=====================================
FROM base AS builder

# Copy only necessary service-specific files
COPY  cmd/blog-service /blog-service/cmd/blog-service
COPY  cmd/blog-service /blog-service/cmd/blog-service
COPY  config/shared /user-service/config/shared
COPY  internal/blog-service /blog-service/internal/blog-service
COPY  pkg/middlewares /blog-service/pkg/middlewares
COPY  pkg/security /blog-service/pkg/security

# Compile the Go application
RUN go build -o /blog-service/bin/blog-service ./cmd/blog-service

#=====================================
# Stage 4: Production Ready Image
#=====================================
FROM alpine:latest AS prod 

# Create a non-root user for security
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

# Set the working directory
WORKDIR /blog-service

# Copy only the compiled binary from the builder stage
COPY --from=builder /blog-service/bin/blog-service /blog-service/bin/blog-service

# Ensure correct permissions
RUN chown -R appuser:appgroup /user-service

# Set non-root user
USER appuser

# Expose the application port
EXPOSE 7200

# Run the compiled binary
CMD ["/blog-service/bin/blog-service"]



