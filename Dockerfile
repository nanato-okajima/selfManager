FROM golang:1.17.5-alpine3.15 AS dev

WORKDIR /go/src
RUN apk update &&\
    apk upgrade &&\
    apk add git
COPY go.mod go.sum ./
RUN go mod download
RUN go install github.com/cosmtrek/air@latest
COPY .air.toml .

CMD ["air", "-c", ".air.toml"]



FROM golang:1.17.5-alpine3.15 AS builder

WORKDIR /go/src
COPY . .
RUN go mod download
ARG CGO_ENABLED=0
ARG GOOS=linux
ARG GOARCH=amd64
RUN go build \
    -o /go/bin/main



FROM scratch  AS prod
WORKDIR /go/src
COPY ./templates ./templates

COPY --from=builder /go/bin/main /go/bin/main

ENTRYPOINT [ "/go/bin/main" ]
