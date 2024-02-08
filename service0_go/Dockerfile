# Start from a base image with Go installed
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the Go files and any other necessary files
COPY . .

# Download dependencies
RUN go get -d -v ./...

# Build the application
RUN go build -o main .

# Expose the port the server listens on
EXPOSE 5001

# Run the binary
CMD ["./main"]
