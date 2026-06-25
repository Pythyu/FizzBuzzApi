FROM golang:alpine

WORKDIR /fizzbuzz
COPY . .

RUN go build -o ./bin/api ./cmd/api

CMD ["/fizzbuzz/bin/api"]
EXPOSE 8989