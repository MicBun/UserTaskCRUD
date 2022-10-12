FROM golang:alpine

WORKDIR /UserSimpleCRUD

ADD . .

RUN go mod download

ENTRYPOINT go build  && ./UserSimpleCRUD