# Use the official Golang image to build the app
FROM golang:1.22.5 AS builder


# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the Go Modules manifest
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o main .

EXPOSE 50054

# Command to run the executable
CMD ["./main"]
