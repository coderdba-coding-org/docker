=====================
WEB SERVER IN GOLANG
=====================

=====================
GOLANG COMPILER IMAGE
=====================

- Existing image on laptop:
$ docker images | grep golang
golang                                                          1.13          84125009cb55   14 months ago   803MB

$ docker run -ti golang:1.13 /bin/sh
# go version
go version go1.13.12 linux/amd64

- Pull latest version
https://hub.docker.com/_/golang?tab=tags&page=1&ordering=last_updated

$ docker pull golang:alpine3.14
$ docker images |grep golang
golang                                                          alpine3.14    4d3587ec7acf   8 days ago      315MB

$ docker run -ti golang:alpine3.14 /bin/sh
/go # go version
go version go1.17 linux/amd64


=========================================
SET UP GO MODULE
=========================================
- Initialize
$ go mod init goweb1
$ ls 
go.mod

- Create main.go
$ ls
go.mod
main.go

- Run the program once to download dependencies and create go.sum
$ go run .

$ ls 
go.mod
go.sum
main.go

=========================================
CREATE IMAGE WITH YOUR WEB SERVER PROGRAM
=========================================
https://docs.docker.com/language/golang/build-images/
- https://github.com/olliefr/docker-gs-ping (code)

- Create a Dockefile

- Build image
$ docker image build . -t image02-golang:1.0

$ docker images |grep golang
image02-golang                                                  1.0           2d5e04fe6d68   2 minutes ago   467MB

[+] Building 23.6s (12/12) FINISHED                                                                                                                                   
 => [internal] load build definition from Dockerfile                                                                                                             0.0s
 => => transferring dockerfile: 321B                                                                                                                             0.0s
 => [internal] load .dockerignore                                                                                                                                0.0s
 => => transferring context: 2B                                                                                                                                  0.0s
 => [internal] load metadata for docker.io/library/golang:alpine3.14                                                                                             0.0s
 => [1/7] FROM docker.io/library/golang:alpine3.14                                                                                                               0.0s
 => [internal] load build context                                                                                                                                0.0s
 => => transferring context: 6.63kB                                                                                                                              0.0s
 => [2/7] WORKDIR /app                                                                                                                                           0.0s
 => [3/7] COPY go.mod ./                                                                                                                                         0.0s
 => [4/7] COPY go.sum ./                                                                                                                                         0.0s
 => [5/7] RUN go mod download                                                                                                                                   19.1s
 => [6/7] COPY *.go ./                                                                                                                                           0.0s
 => [7/7] RUN go build -o /myapp                                                                                                                                 3.3s
 => exporting to image                                                                                                                                           1.0s
 => => exporting layers                                                                                                                                          1.0s
 => => writing image sha256:2d5e04fe6d686701288fc5eb04fddd1fec5e0cbc92a59c2019f6b165e895de50                                                                     0.0s
 => => naming to docker.io/library/image02-golang:1.0                                                                                                            0.0s

=================================
RUN THE IMAGE AS A CONTAINER
=================================
https://docs.docker.com/engine/reference/commandline/run/
Example: docker run -p 127.0.0.1:80:8080/tcp ubuntu bash --> This binds port 8080 of the container to TCP port 80 on 127.0.0.1 of the host machine. 

- Run the image (quick test)
$ docker run -d --rm image02-golang:1.0

$ docker ps -a
CONTAINER ID   IMAGE                 COMMAND                  CREATED          STATUS                        PORTS      NAMES
86d48f4f75c3   image02-golang:1.0    "/myapp"                 15 seconds ago   Up 15 seconds                 8081/tcp   recursing_snyder

$ docker stop 86d48f4f75c3
$ docker rm 86d48f4f75c3

- Run the image
$ docker run --name image02-golang1 -d --rm -p 8081:8081/tcp image02-golang:1.0 
b9446820177c8f6d229d30d2b07b39a51a0b9ee16f37f11f0b730da012773794

$ docker ps -a
CONTAINER ID   IMAGE                COMMAND    CREATED         STATUS         PORTS                                       NAMES
b9446820177c   image02-golang:1.0   "/myapp"   4 seconds ago   Up 2 seconds   0.0.0.0:8081->8081/tcp, :::8081->8081/tcp   image02-golang1

-- Verify
$ curl localhost:8081
{"message":"Welcome!"}



