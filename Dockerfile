FROM alpine

RUN apk add --no-cache ca-certificates

ADD release/quadlekBot /

EXPOSE 8000

ENTRYPOINT ["/quadlekBot"]