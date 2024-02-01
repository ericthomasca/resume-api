# Use the official Golang image as the base image
FROM golang:1.21

# Set the working directory inside the container
WORKDIR /usr/src/app

# Pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# Copy the rest of the application files
COPY . .

# Build the application
RUN go build -v -o /usr/local/bin/app ./...

# Expose port 7777
EXPOSE 7777

# Set the command to run the application
CMD ["app"]
