FROM golang:1.18-alpine as builder

RUN apk add --update ca-certificates \
    && apk add curl git coreutils \
    && rm /var/cache/apk/*

ENV APP_PATH=/quadlek
RUN mkdir -p $APP_PATH
ADD . $APP_PATH
WORKDIR $APP_PATH
RUN go build -mod=vendor -o /build/quadlekBot


FROM alpine

RUN apk add --no-cache ca-certificates

COPY --from=builder /build/quadlekBot /
ADD assets /opt/quad-assets

EXPOSE 8000

ENTRYPOINT ["/quadlekBot"]
