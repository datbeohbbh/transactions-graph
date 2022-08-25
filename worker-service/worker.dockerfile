FROM golang:1.18-alpine AS builder

RUN apk add --no-cache gcc musl-dev linux-headers

WORKDIR /app

COPY ./go.mod /app
COPY ./go.sum /app
RUN cd /app && go mod download

COPY . /app
RUN go build -o /app/Worker /app/cmd/api
RUN chmod +x /app/Worker

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/Worker /app

CMD [ "/app/Worker" ]