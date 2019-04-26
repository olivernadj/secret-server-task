FROM golang:1.12
MAINTAINER Oliver Nadj <mr.oliver.nadj@gmail.com>

ADD src /go/src/goapi
ADD src /go/src/github.com/olivernadj/secret-server-task/goapi/src

EXPOSE 8080
WORKDIR /go/src/goapi
ENTRYPOINT [ "go", "run",  "main.go",  "-p",  "8080" ]
