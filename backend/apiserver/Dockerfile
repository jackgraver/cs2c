# Use a small Go image
FROM golang:1.21-alpine

# Install Air for live reload
RUN apk add --no-cache git curl bash && \
    go install github.com/cosmtrek/air@latest

WORKDIR /app

# Copy go.mod/go.sum first for caching
COPY go.mod ./
RUN go mod download

# Copy rest of the code
COPY . .

# Expose API port
EXPOSE 8080

# Run Air (looks for .air.toml)
# CMD ["air", "-c", ".air.toml"]
CMD ["go", "run", "main.go"]