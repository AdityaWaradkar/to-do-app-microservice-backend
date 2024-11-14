# Use the official Golang image as the base image
FROM golang:1.23

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files first (this helps cache dependencies)
COPY go.mod go.sum ./

# Install dependencies
RUN go mod tidy

# Copy the entire contents of the todo-service folder into the container
COPY ./ ./

# Build the Go application
RUN go build -o todo-service

# Expose port the application will run on
EXPOSE 8082

# Command to run the application
CMD ["./todo-service"]
