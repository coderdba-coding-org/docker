#docker run -it -v /var/run/docker.sock:/var/run/docker.sock ubuntu:latest sh -c "apt-get update ; apt-get install docker.io -y ; bash"
docker run -it -v /var/run/docker.sock:/var/run/docker.sock ubuntu:18.04 sh -c "apt-get update ; apt-get install docker.io -y ; bash"
