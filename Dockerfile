
FROM golang:1.10 AS build-env

ADD . /go/src/github.com/guusvw/time.machine
WORKDIR /go/src/github.com/guusvw/time.machine

RUN  go build -o goapp

FROM scratch

EXPOSE 8080

COPY --from=build-env /go/src/github.com/guusvw/time.machine/goapp /usr/bin/time.machine

ENTRYPOINT ["/usr/bin/time.machine"]
