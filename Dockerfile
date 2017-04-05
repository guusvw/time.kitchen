FROM alpine

EXPOSE 8080
ADD ./time.kitchen /usr/bin/time.machine

ENTRYPOINT ["/usr/bin/time.machine"]
