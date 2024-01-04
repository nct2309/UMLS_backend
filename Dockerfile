ENV GOPROXY=""

# Use the official Golang 1.21 image as the base image
FROM golang:1.21

# Set the working directory inside the container
WORKDIR /app

# Copy the main.go file into the container at /app
COPY . /app

# Download and install any required dependencies
RUN go mod download

# Build the binary
RUN go build -o myapp .

# Expose the port that your application will run on
EXPOSE 8080

# Command to run the executable
CMD ["./myapp"]

