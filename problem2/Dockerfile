FROM golang:latest
RUN apt-get update && apt-get install -y xinetd \
    && useradd -m problem2 \
    && chown -R root:root /home/problem2

RUN mkdir -p /go/src/github.com/tukeJonny/problem2
WORKDIR /go/src/github.com/tukeJonny/problem2
ADD main.go .

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

RUN go build -o /home/problem2/problem2 main.go

ADD xinetd.conf /etc/xinetd.d/problem2

RUN chmod 755 /home/problem2/problem2 \
    && chmod 644 /etc/xinetd.d/problem2

EXPOSE 8000

CMD ["/usr/sbin/xinetd", "-dontfork"]
