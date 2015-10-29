FROM golang:1.5

ADD vendor/src /go/src
ADD src /go/src
RUN go install ./...

CMD ["/go/bin/bot"]
