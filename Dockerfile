FROM golang:1.16.3

WORKDIR /app

COPY ./ /app

RUN go mod download

ENTRYPOINT make build