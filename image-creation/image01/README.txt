====================
KUBECTL IMAGE
====================

====================
PULL ALPINE IMAGE
====================
https://hub.docker.com/_/alpine
docker pull alpine:3.14.1

====================
Dockerfile
====================
Reference: https://github.com/lwolf/kubectl-deployer-docker/blob/master/Dockerfile
Change kubectl version and download location of kubectl as needed in Dockerfile

--> Change curl kubectl with latest URL from https://kubernetes.io/docs/tasks/tools/install-kubectl-linux/#install-kubectl-binary-with-curl-on-linux

--> For specific kubectl version
curl -LO https://dl.k8s.io/release/v1.22.0/bin/linux/amd64/kubectl

--> For latest kubectl version
curl -LO "https://dl.k8s.io/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl.sha256"

====================
BUILD IMAGE
====================
NOTE: 1.22 is the kubectl version - change it as needed

docker image build . -t alpine-kubectl:1.22

====================
RUN 
====================
docker run -it --rm alpine-kubectl:1.22 kubectl version

====================
RUN - INTERACTIVE
====================
- This will remove the container afterwards
docker run -it --rm alpine-kubectl:1.22 /bin/sh

- This one will not remove/cleanup the container after running
docker run -it alpine-kubectl:1.22 /bin/sh
