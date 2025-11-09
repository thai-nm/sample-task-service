FROM golang:1.19-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o task-service .

FROM alpine:3.16

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/task-service .

EXPOSE 8080

CMD ["./task-service"]