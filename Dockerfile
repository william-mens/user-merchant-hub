# Step 1: Build the Go application
FROM golang:1.22.0 as builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files first to leverage Docker cache for dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the Go app as a statically linked binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/main ./main.go

# Step 2: Create a lightweight image to run the binary
FROM alpine:latest

# Set working directory
WORKDIR /app

# Copy the binary from the builder stage
COPY --from=builder /app/main .

# Copy the .env file from the build context into the runtime container
COPY .env .env

# Expose port (if needed by your application)
EXPOSE 8080

# Run the application
CMD ["/app/main"]
