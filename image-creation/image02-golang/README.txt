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


