FROM golang:1.8.3-alpine3.6

ADD src /go/src

RUN go install app
RUN go test app

CMD app