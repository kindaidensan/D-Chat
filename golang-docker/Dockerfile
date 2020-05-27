FROM golang:1.14.3-buster

ENV GOPATH /.go

RUN apt-get -y update && \
    apt-get -y install git

RUN go get -v github.com/gin-gonic/gin

CMD ["go", "run", "main.go"]