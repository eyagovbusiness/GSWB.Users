FROM golang:1.24.4-alpine

WORKDIR /app

# Install git (needed for go mod) and build deps
RUN apk add --no-cache git

# Copy go mod files and download deps early (build cache)
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the project
COPY . .

# Expose API port
EXPOSE 8080

CMD ["go", "run", "./cmd/server"]
