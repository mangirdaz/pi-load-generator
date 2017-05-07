### PI load generator
Commands:
        `make build` - builds go binary
        `make build-image` - builds previous step and docker image
        `make publish` - build previous step and publish image to docker.io
        
        
To execute:
    `curl localhost:8080/api/v1/pi/300` where 300 is how many PI approximations will get executed. All of them will be in separate go thread/go routine.
    
Run:
    `docker run -it -p8080:8080 docker.io/mangirdas/pi-load-generator`