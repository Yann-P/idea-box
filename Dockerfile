FROM golang:alpine

RUN apk add --no-cache git \
    && go get -u github.com/go-pg/pg \
    && apk del git

ADD ./src /app
WORKDIR /app
CMD ["go", "run", "main.go"]
