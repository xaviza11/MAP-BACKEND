# Use the official Golang image as the base image
FROM golang:1.23.4
# Set the Current Working Directory inside the container
WORKDIR /app
# Copy go.mod and go.sum files
COPY go.mod ./
# Install dependencies
RUN go mod download
# Copy the source code
COPY . .
# Build the Go app
RUN go build -v -o go-sqlite-backend .
# Expose port 4000 to the outside world
EXPOSE 4000
# Enable permisions
RUN chmod +x go-sqlite-backend
# Command to run the executable
CMD ["/app/go-sqlite-backend"]J