FROM golang:1.17.5-alpine3.15

WORKDIR /go/src/github.com/

RUN apk update &&\
    apk upgrade &&\
    apk add git

RUN go install github.com/cosmtrek/air@latest
COPY .air.toml .

CMD ["air", "-c", ".air.toml"]
