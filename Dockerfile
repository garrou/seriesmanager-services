FROM golang:alpine

WORKDIR /app

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

ENV GIN_MODE=release

RUN go build -o seriesmanager-services

EXPOSE 8080

CMD ["seriesmanager-services"]