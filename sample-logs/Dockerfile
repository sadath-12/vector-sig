# Use a specific platform and select the Golang Alpine image as the builder stage
FROM --platform=linux/amd64 golang:1.22-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY *.go ./

# Build the Go application
RUN go build -o main .

# Expose port 8080 (optional, you might not need it in Alpine)
EXPOSE 8080

# Set the command to run the executable
CMD ["./main"]
