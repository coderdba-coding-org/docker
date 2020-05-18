FROM bitnami/kubectl:latest

RUN apk add --update dumb-init 
ADD kube-autofix /

ENTRYPOINT ["dumb-init", "/kube-autofixer"]
