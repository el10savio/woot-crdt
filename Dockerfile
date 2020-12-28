FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git 

RUN mkdir /woot

WORKDIR /woot

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -a -installsuffix cgo -o /go/bin/woot

ENTRYPOINT ["/go/bin/woot"]

EXPOSE 8080