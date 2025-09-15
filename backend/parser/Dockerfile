FROM golang:1.21-alpine

RUN apk add --no-cache git curl bash && \
    go install github.com/cosmtrek/air@latest

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

# No port needed unless you expose metrics
# CMD ["air", "-c", ".air.toml"]
CMD ["go", "run", "main.go"]