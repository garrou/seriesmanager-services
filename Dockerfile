FROM golang:alpine

ENV GIN_MODE=release

WORKDIR /app

RUN go build

COPY . .

EXPOSE 8080

ENTRYPOINT ["./app"]