FROM golang:1.18-alpine as builder 

WORKDIR /app 

COPY ./go.mod /app
COPY ./go.sum /app

RUN go mod download

COPY ./ /app

RUN go build -o /app/addressManager /app/cmd/api

RUN chmod +x /app/addressManager

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/addressManager /app

CMD ["/app/addressManager"]