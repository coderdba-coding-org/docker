================================
POD RESTARTER 
================================

To use in Kubernetes to watch a pod and restart it if it is not running

================================
BUILD THE IMAGE
================================
Create Dockerfile

$ docker image build . -t pod-restarter:1.0

================================
RUN THE IMAGE
================================
$ docker run -d --name pod-restarter pod-restarter:1.0 
aed43b383588ae8123945fcb00cd22b8f7fe07d488c6c0a1f3795d48cd5006da

C02Z10JHLVDQ:image03-admin-pod-restarter dbgsm0$ docker ps -a
CONTAINER ID   IMAGE                 COMMAND                  CREATED         STATUS         PORTS     NAMES
aed43b383588   pod-restarter:1.0     "/app/pod_restarter.…"   3 seconds ago   Up 2 seconds             pod-restarter

================================
VERIFY THE CONTAINER
================================

docker exec -ti pod-restarter /bin/sh
/ # 
--> then explore in the container if the kubeconfig, script files are copied and if /tmp has the restarter logs
