# Use the official Golang image as a base
FROM golang:alpine

# Set the working directory in the container
WORKDIR /usr/src/app

# Copy the Go modules to the working directory
COPY go.mod ./

# Install the dependencies
RUN go mod download

# Copy the rest of the code to the working directory
COPY . .

# Build the executable
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/main ./cmd

# Run the command when the container starts
CMD ["./build/main"]
