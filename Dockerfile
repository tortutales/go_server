# syntax=docker/dockerfile:1
FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /svc-feelify-colors

EXPOSE 5000

CMD [ "/svc-feelify-colors" ]

