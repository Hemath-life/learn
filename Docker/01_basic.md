what is container :
  - A way to package application with all the necessary dependencies and configuration.
  - portable artifact , easily shared and moved around
  - makes development and deployment more efficient

where do containers live :
  - container repository
  - private repositories
  - public repository for docker

after container :
  - developers and operations work together to package the application in a container
  - no environmental configuration needed on server - except Docker runtime.

What is container in technical manner:
  - Layers of images
      1. Base Image
                -Mostly Linux Base Image, coz small in size
      2. Application image on top
                - psql , node anything

cmd :
  sudo docker run hello-world:
    - to test docker 

  sudo docker images:  
    - see all the pulled image in system

containers vs virtual machine :
  - Both are virtualizations tools
  - Docker :
    - Docker is os level
        1. Hardware
        2. Os Kernel - 1st layer
        3. Application - 2ed - layer
          - Different levels of abstractions
          - Why linux-based docker containers don't run in windows

Docker :
  - virtualize the Application level
  - docker images much smaller
  - docker containers start and run much faster

vm :
  - virtualize the Application and os Kernel
  - vm have to download Application images along with kernel image
  - vm of any OS can run on any OS host

Containers vs image :
  - CONTAINERS is a running environment of IMAGE
  - application image:
    - postgres,redis,mongo
    - port Binded: talk to application running inside of container
    - virtual file system

Version vs Tag :
    - TAG is basically the version of the image

Docker Commands :
    docker Version:
    docker images:
        - to get downloaded images details
    docker run image_name:
      - d __ 
      - pull the image and starts container
      - p_laptop_port:container_port 
          - p6000:3782
          - this is called port binding
          - -- name container_name_that_you want
      - example:
          docker run -d -p6000:3782 --name my-first-container redis:4.0
    docker ps -a:
      - ps(Process Status) ___ lists all containers that are up and running. 
      - a  ___ means all (both stopped and running) containers
    docker stop container_id:
    docker start container_id:
    docker pull image_name:
      - pull the image and starts container

CONTAINER port vs HOST port :
  - Multiple containers can run on your host machine
  - Your laptop has only certain ports available
  - conflict when same port on host machine
  - port binding the blow u can see
      host  
        laptop          port 5000           port 3000           port 3001

        container      port 5000           port 3000           port 3000
          
      docker run -d -p6000:6379 redis
           
Debugging Containers :
  docker logs container_id (or) container_name      /bin/bash:
  docker exec -it container_id (or) container_name      /bin/bash:
      - it(interactive terminal)
  examples:
      - sudo docker exec -it redis-older    /bin/bash

Docker run vs Ducker start:
  docker run creates the new container: 
  docker start :
    - your not working with images your are working with containers
            
links :
  - https://hub.docker.com/

cli :
  docker -v 
    - to see docker version
          
  docker version      
    - to see more information about docker

  docker run image_name:version ---- 
    -to get image from docker repo