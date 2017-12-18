FROM alpine

RUN apk add --no-cache --update ca-certificates

EXPOSE 8080
ADD ./time.kitchen /usr/bin/time.machine

ENTRYPOINT ["/usr/bin/time.machine"]
