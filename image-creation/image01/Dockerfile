FROM alpine:3.14.1

LABEL MAINTAINER="coderdba"

ENV KUBE_LATEST_VERSION="v1.22"

RUN apk add --update ca-certificates \
    && apk add --update -t deps curl \
    && apk add --update gettext \
    && curl -L https://dl.k8s.io/release/v1.22.0/bin/linux/amd64/kubectl -o /usr/local/bin/kubectl \
    && chmod +x /usr/local/bin/kubectl \
    && apk del --purge deps \
    && rm /var/cache/apk/* \
    && mkdir /root/.kube

COPY admin.kubeconfig.ksn2 /root/.kube/.
COPY admin.kubeconfig.ksn3 /root/.kube/.

