FROM golang:1.18-alpine as builder 

# RUN apk add --no-cache gcc musl-dev linux-headers

WORKDIR /app 

COPY ./go.mod /app
COPY ./go.sum /app

RUN go mod download

COPY ./ /app

RUN go build -o /app/graphService /app/cmd/api

RUN chmod +x /app/graphService

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/graphService /app

CMD ["/app/graphService"]