=======================
DOCKER IMAGE CREATION
=======================

===================
REFERENCES
===================

First Alpine Docker Image and Push it to DockerHub - https://dockerlabs.collabnix.com/beginners/building-your-first-alpine-container.html

===================
CAUTION
===================
Dont use Busybox 1.33 based base images - it does not resolve nslookup kubernetes.default 
- Instead use busybox 1.26 based images
- Maybe Alpine:3.6 has busybox 1.26

===================
IMAGE DESCRIPTIONS
===================
image01-kubectl
- Kubectl runner

image02-golang
- Simple web server with golang with / and /message
- Pushed to docker hub as coderdba/goweb1
-- Using this repo: docker-git-workflow01

image03-admin-pod-restarter
- Admin container - to keep an eye on a pod and restart it if it is down
