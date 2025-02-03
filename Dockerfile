# Use official Golang image
FROM golang:1.23

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum, then download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the project
COPY . .

# Build the application
RUN go build -o main ./cmd/main.go

# Run the built binary
CMD ["./main"]
