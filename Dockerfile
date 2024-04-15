FROM golang:1.21.1-alpine

WORKDIR /www

ADD . .

RUN go build

CMD ["./godzilla", "--config", "./godzilla.yaml"]
