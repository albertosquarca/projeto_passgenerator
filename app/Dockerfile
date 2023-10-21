# Use the official Golang image as the base image
FROM golang:latest as build

# Set the working directory inside the container
WORKDIR /app

# Copie a pasta de modelos para o contêiner
COPY templates /app/templates

# Copy the Go module and Go sum files to the container
COPY go.mod go.sum ./

# Download the application dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

# Create a minimal image with just the binary
FROM alpine:latest
LABEL maintainer="albertosquarca@gmail.com"

# Copy the binary from the previous stage
COPY --from=build /app/app /app/app

# Expose the port your application will listen on
EXPOSE 8080

# Command to run your application
CMD ["/app/app"]
