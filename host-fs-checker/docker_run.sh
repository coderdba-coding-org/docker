# no mount directory (not applicable for this project)
#docker run -d host-fs-checker

# mount one directory
#docker run -d --name host-fs-checker -v /tmp/fsmetrics:/fsmetrics host-fs-checker:latest  

# mount two directories
docker run -d --name host-fs-checker -v /tmp/fsmetrics:/fsmetrics -v /tmp/fsmetricsOut:/fsmetricsOut host-fs-checker:latest  
