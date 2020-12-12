FROM golang:latest

RUN useradd -m problem3-app
RUN mkdir -p /go/src/github.com/tukeJonny/problem3-app
WORKDIR /go/src/github.com/tukeJonny/problem3-app
ADD main.go .

RUN go mod init
RUN go get -u -v github.com/julienschmidt/httprouter
RUN go get -u -v github.com/urfave/cli/v2

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

RUN go build -o /home/problem3-app/problem3-app main.go

EXPOSE 8000

ENTRYPOINT ["/home/problem3-app/problem3-app"]
