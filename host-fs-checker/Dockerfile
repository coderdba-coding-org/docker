FROM alpine:latest

ADD scripts/* /scripts/

RUN apk add --no-cache --virtual .fetch-deps \
    bash \
    jq \
    curl

WORKDIR /scripts
ENTRYPOINT /scripts/fs.sh
