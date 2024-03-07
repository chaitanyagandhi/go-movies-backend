# # Start from the official Golang image
# FROM golang:latest

# # Set the Current Working Directory inside the container
# WORKDIR /app

# # Copy the Go modules files
# COPY go.mod go.sum ./

# # Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
# RUN go mod download

# # Copy the entire project directory into the Working Directory inside the container
# COPY . .

# # Build the Go app
# RUN go build -o main ./cmd/api

# # Command to run the executable
# CMD ["./main"]


# syntax=docker/dockerfile:1

FROM golang:1.22.0

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/reference/dockerfile/#copy
COPY . .

# Build
# RUN go build -o /gomovies

# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/reference/dockerfile/#expose
EXPOSE 8080

# Run
# CMD ["/gomovies"]