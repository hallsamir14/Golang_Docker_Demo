# syntax=docker/dockerfile:1

# Stage 1: Build
FROM golang:1.24 AS builder

# Set working directory
WORKDIR /src

# Copy go.mod and download dependencies
COPY ./src/go.mod ./
RUN go mod download

# Copy the source code to working directory
COPY ./src ./

# Build the Go binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main.bin main.go

# Stage 2: Run
FROM alpine:latest

# Set working directory
WORKDIR /root/

#Copy json samples to root
COPY ./samples /samples

# Copy the binary from the builder stage
COPY --from=builder /src/main.bin .

# Expose the port (if your app listens on a specific port, e.g., 8080)
# EXPOSE 8080

# Run the binary
CMD ["./main.bin"]