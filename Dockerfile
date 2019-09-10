FROM golang:1.13-alpine

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

LABEL maintainer="20scoops <dev.20scoops@gmail.com>"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]