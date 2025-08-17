
FROM golang:1.25 AS builder


WORKDIR /app


COPY go.mod go.sum ./

RUN go mod download

COPY . .


RUN CGO_ENABLED=0 GOOS=linux go build -o /app/exchange-rate-service ./main.go


FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/exchange-rate-service .

EXPOSE 8080


CMD ["./exchange-rate-service"]