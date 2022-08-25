FROM golang:1.18-alpine AS builder

WORKDIR /app

COPY ./go.mod /app
COPY ./go.sum /app

RUN go mod download

COPY ./ /app

RUN CGO_ENABLED=0 go build -o /app/brokerApp /app/cmd/api

RUN chmod +x /app/brokerApp

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/brokerApp /app

CMD ["/app/brokerApp"]