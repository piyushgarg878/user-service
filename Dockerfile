FROM golang:1.21-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o user-service ./cmd/main.go

EXPOSE 8080

CMD ["./user-service"]