# Use a multi-stage build to keep the final image small
FROM golang:1.20.1-alpine3.17 as builder
# Set the working directory to /app
WORKDIR /app
# Copy the go.mod and go.sum files to the container
COPY go.mod go.sum ./
# Download the dependencies
RUN go mod download
# Copy the rest of the application to the container
COPY . .
# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -o app
# Use a minimal alpine image for the final image
FROM alpine:3.17
# Set the working directory to /app
WORKDIR /app
# Copy the built application from the builder image
COPY --from=builder /app/app .
# Use the ENTRYPOINT command to set the default command for the container
ENTRYPOINT ["/app/app"]
