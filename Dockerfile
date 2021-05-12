FROM golang:1.16
WORKDIR /go/app
ADD . /go/app
RUN go build -o main .
ENTRYPOINT [ "./main" ]