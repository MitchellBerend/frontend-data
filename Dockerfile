FROM golang:1.22.1-bookworm AS builder

WORKDIR /app

# Copy go.mod and go.sum files to the workspace
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go application
RUN GOOS=linux go build -o frontend-data cmd/api/main.go

# # Stage 2: Final Stage
# FROM gcr.io/distroless/base-debian12

# testing different base image
FROM golang:1.22.1-bookworm

# Copy the executable from the builder stage
COPY --from=builder /app/frontend-data /

EXPOSE 8080

# Command to run the executable
CMD ["/frontend-data"]
