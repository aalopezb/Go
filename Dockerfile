# Use official Go image
FROM golang:1.21-alpine

# Set working directory
WORKDIR /app

# Copy files and build the application
COPY main.go ./
RUN go build -o app main.go

# Expose port 8080
EXPOSE 8080

# Command to run the application
CMD ["./app"]
