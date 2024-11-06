# Start with the official Go image
FROM golang:1.23

# Set the working directory inside the container
WORKDIR /app

# Copy the entire contents of the user-service folder into the container
COPY ./ ./

# Install dependencies
RUN go mod tidy

# Build the Go application
RUN go build -o user-service

# Command to run the application
CMD ["./user-service"]
