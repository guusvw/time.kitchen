FROM alpine

apk add -U tzdata

EXPOSE 8080
ADD ./time.kitchen /usr/bin/time.machine

ENTRYPOINT ["/usr/bin/time.machine"]
