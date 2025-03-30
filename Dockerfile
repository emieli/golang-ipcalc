FROM golang:1.24.1-bookworm AS base
LABEL org.opencontainers.image.source=https://github.com/emieli/golang-ipcalc

# Builder stage
# =============================================================================
# Create a builder stage based on the "base" image
FROM base AS builder

# Move to working directory /build
WORKDIR /build

# Copy the go.mod and go.sum files to the /build directory
COPY go.mod go.sum ipcalc.go .

# Install dependencies
RUN go mod download

# Copy the entire source code into the container
COPY . .

# Build the application
# Turn off CGO to ensure static binaries
RUN CGO_ENABLED=0 go build -o main

# Production stage
# =============================================================================
# Create a production stage to run the application binary
FROM scratch AS production

# Move to working directory /prod
WORKDIR /prod

# Copy binary from builder stage
COPY --from=builder /build/main ./
COPY ./static ./static

# Document the port that may need to be published
EXPOSE 8000

# Start the application
CMD ["/prod/main"]
