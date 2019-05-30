FROM golang:latest

WORKDIR $GOPATH/src/github.com/Huangscar/gloudgo-io
ADD . $GOPATH/src/github.com/Huangscar/gloudgo-io
RUN go build .

EXPOSE 8080

ENTRYPOINT ["./main"]