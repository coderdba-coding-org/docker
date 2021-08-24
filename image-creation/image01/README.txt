====================
KUBECTL IMAGE
====================
Reference: https://github.com/wayarmy/alpine-kubectl
Reference: https://github.com/lwolf/kubectl-deployer-docker/blob/master/Dockerfile

====================
PULL ALPINE IMAGE
====================
https://hub.docker.com/_/alpine
docker pull alpine:3.14.1

====================
Dockerfile
====================
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

======================================================
RUN WITH KUBECONFIG THAT IS COPIED OVER TO THE IMAGE
======================================================
$ docker run -it --rm alpine-kubectl:1.22 kubectl get pods --all-namespaces --kubeconfig=/root/.kube/admin.kubeconfig.ksn3
NAMESPACE       NAME                                        READY   STATUS    RESTARTS   AGE
default         web-79d88c97d6-h5kqp                        1/1     Running   1          2d6h
default         web2-5d47994f45-44xqs                       1/1     Running   1          2d6h
ingress-nginx   ingress-nginx-controller-5b97d5cd4b-v75q5   1/1     Running   1          2d8h
kube-system     calico-kube-controllers-86475544f5-5fvks    1/1     Running   3          3d5h
kube-system     calico-node-mdv66                           1/1     Running   3          3d5h
kube-system     coredns-8494f9c688-vcxrj                    1/1     Running   3          3d6h
kube-system     kube-apiserver-ksn3                         1/1     Running   4          3d11h
kube-system     kube-controller-manager-ksn3                1/1     Running   4          3d11h
kube-system     kube-proxy-jv8pc                            1/1     Running   3          3d6h
kube-system     kube-scheduler-ksn3                         1/1     Running   4          3d11h

====================
RUN - INTERACTIVE
====================
- This will remove the container afterwards
docker run -it --rm alpine-kubectl:1.22 /bin/sh

- This one will not remove/cleanup the container after running
docker run -it alpine-kubectl:1.22 /bin/sh
