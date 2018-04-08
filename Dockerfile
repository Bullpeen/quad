FROM alpine

RUN apk add --no-cache ca-certificates

ADD release/quadlekBot /
ADD assets /opt/quad-assets

EXPOSE 8000

ENTRYPOINT ["/quadlekBot"]