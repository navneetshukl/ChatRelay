FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

FROM golang:1.23

WORKDIR /app

COPY --from=builder /app /app

CMD go run cmd/http/main.go && go run cmd/main.go
