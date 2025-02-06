# Stage 1: Base Image (Common for Dev & Prod)
FROM golang:1.23-alpine AS base

# Set environment variables
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    SERVICE_NAME=email-service

# Set the working directory
WORKDIR /email-service

# Install system dependencies
RUN apk add --no-cache curl bash git make

# Copy only necessary dependency files to cache layers efficiently
COPY go.mod go.sum ./
RUN go mod download

# Copy the Makefile into the container
COPY Makefile ./
COPY misc/make /email-service/misc/make

# Install tools (migrate, air) using the Makefile
RUN make deps

# Stage 2: Development (Hot Reloading)
FROM base AS dev

# Copy only service-specific files
COPY .air.toml /email-service/.air.toml
COPY cmd/email-service /email-service/cmd/email-service
COPY internal/email-service/adapters/rabbitmq_consumer.go /email-service/internal/email-service/adapters/rabbitmq_consumer.go
COPY internal/email-service/services/email_service.go /email-service/internal/email-service/services/email_service.go

# Inject service name into the .air.toml file dynamically
RUN sed -i 's/\$SERVICE_NAME/email-service/' /email-service/.air.toml

# Expose the port the service listens on
EXPOSE 8100

# Run Air for hot reloading
CMD ["bin/air", "-c", "/email-service/.air.toml"]

# Stage 3: Build Application (Production)
FROM base AS builder

# Copy necessary service-specific files
COPY cmd/email-service /email-service/cmd/email-service
COPY adapters/rabbitmq_consumer.go /email-service/adapters/rabbitmq_consumer.go
COPY services/email_service.go /email-service/services/email_service.go

# Compile the Go application
RUN go build -o /email-service/bin/email-service ./cmd/email-service

# Stage 4: Production Ready Image
FROM alpine:latest AS prod

# Create a non-root user for security
RUN addgroup -S appgroup && adduser -S appuser -G appgroup

# Set the working directory
WORKDIR /email-service

# Copy the compiled binary from the builder stage
COPY --from=builder /email-service/bin/email-service /email-service/bin/email-service

# Ensure correct permissions
RUN chown -R appuser:appgroup /email-service

# Set non-root user
USER appuser

# Expose the application port
EXPOSE 8100

# Run the compiled binary
CMD ["/email-service/bin/email-service"]
