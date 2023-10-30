

[What is Container ?](#what-is-container)

[Where do containers live ?](#where-do-containers-live)

[After container ?](#after-container)

[What is container in technical manner ?](#what-is-container-in-technical-manner)

[Containers vs Virtual machine ?](#containers-vs-virtual-machine)

[Container port vs HOST port (port binding)](#container-port-vs-host-port)

[Docker Comments ?](#docker-comments)

[Docker run vs Ducker start](#docker-run-vs-ducker-start)

### What is container ?

- A way to package application with all the necessary dependencies and configuration.
- Portable artifact , easily shared and moved around
- Makes development and deployment more efficient

### Where do containers live ?

- container repository
- private repositories
- public repository for docker

### After container ?

- developers and operations work together to package the application in a container
- no environmental configuration needed on server - except Docker runtime.

### What is container in technical manner ?

- Layers of images

1. Base Image
   - Mostly Linux Base Image, coz small in size
2. Application image on top
   - psql , node, python anything

### Containers vs Virtual machine ?

- Both are virtualizations tools
- CONTAINERS is a running environment of IMAGE.
  - application image: postgres,redis,mongo
  - port Binded: talk to application running - inside of container
  - virtual file system
  - Docker
    - docker images much smaller
    - docker containers start and run much faster
    - Docker is os level
      1. Hardware
      2. Os Kernel - 1st layer
      3. Application - 2ed - layer
         - Different levels of abstractions
         - Why linux-based docker containers don't run in windows
- VM
  - virtualize the Application and os Kernel
  - vm have to download Application images along with kernel image
  - vm of any OS can run on any OS host
```bash
docker -v
   # to see docker version

docker version
    # To see more information about docker

docker run image_name:version
    # To get image from docker repo

```


### Container port vs HOST port ?
- Multiple containers can run on your host machine
- Your laptop has only certain ports available
- conflict when same port on host machine
- port binding the blow u can see
  host  
  laptop          port 5000           port 3000            port 3001
  container       port 5000           port 3000            port 3000
          
```bash
docker run -d -p6000:6379 redis
```



### Docker Comments ?
```bash
docker Version
 
docker images
#  to get downloaded images details
  
docker run image_name
#   tags:
         # - d pulls the image and starts container
         # -p_laptop_port:container_port
         # - p6000:3782
            #  this is called port binding

      #   -- name container_name_that_you want         example:

docker run -d -p6000:3782 --name my-first-container redis:4.0

docker ps
   # - ps(Process Status) lists all containers that are up and running. 
   tags:
      # - a  means all (both stopped and running) containers
docker stop container_id

docker start container_id

docker pull image_name

sudo docker run hello-world
# to test docker 

sudo docker images  
# see all the pulled image in system
```

           
### Debugging Containers ?

``` bash
docker logs container_id (or) container_name      /bin/bash

docker exec -it container_id (or) container_name      /bin/bash
   # - it(interactive terminal) 
# example:
sudo docker exec -it redis-older    /bin/bash
```

### Docker run vs Ducker start ?
- docker run creates the new container
- docker start 
   - your not working with images your are working with containers
          



### Reference

- [Docker Hub](https://hub.docker.com/)
