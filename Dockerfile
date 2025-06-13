# syntax=docker/dockerfile:1
FROM golang:1.24-alpine

# Install git (needed for go mod), and set working dir
RUN apk add --no-cache git

WORKDIR /app

# Copy go.mod and go.sum first for better caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire project
COPY . .

# Build the Go binary
RUN go build -o /main ./cmd/main.go

# Run the binary
CMD ["/main"]
