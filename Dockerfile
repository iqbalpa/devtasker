# Stage 1: Build stage
FROM golang:1.24 AS build

# Set working directory
WORKDIR /usr/src/app

# Copy and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the app
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
RUN go build -v -o /usr/local/bin/app ./cmd

# Stage 2: Final stage
FROM alpine:3.20

# Install curl for healthcheck
RUN apk add --no-cache curl

# Set working directory
WORKDIR /usr/src/app

# Copy binary from the build stage
COPY --from=build /usr/local/bin/app .

CMD [ "./app" ]