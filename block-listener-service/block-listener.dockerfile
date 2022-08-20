FROM golang:1.18-alpine AS builder

RUN apk add --no-cache gcc musl-dev linux-headers

WORKDIR /app

COPY ./go.mod /app
COPY ./go.sum /app
RUN cd /app && go mod download

COPY . /app
RUN go build -o /app/blockListener /app/cmd/api
RUN chmod +x /app/blockListener

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/blockListener /app

RUN mkdir -p /app/db-data

CMD [ "/app/blockListener" ]