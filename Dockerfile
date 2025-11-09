FROM golang:1.19-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o task-service .

FROM alpine:3.16

WORKDIR /root/

COPY --from=builder /app/task-service .

EXPOSE 8080

CMD ["./task-service"]