# Example golang app with air

Example application using golang air with docker-compose for live reload.

```
$ docker-compose up --build
[+] Building 0.3s (7/7) FINISHED                                                                                           
 => [internal] load build definition from Dockerfile                                                                  0.0s
 => => transferring dockerfile: 221B                                                                                  0.0s
 => [internal] load .dockerignore                                                                                     0.0s
 => => transferring context: 2B                                                                                       0.0s
 => [internal] load metadata for docker.io/library/golang:1.21.0                                                      0.3s
 => [1/3] FROM docker.io/library/golang:1.21.0@sha256:b490ae1f0ece153648dd3c5d25be59a63f966b5f9e1311245c947de4506981  0.0s
 => CACHED [2/3] RUN go install github.com/cosmtrek/air@v1.45.0                                                       0.0s
 => CACHED [3/3] WORKDIR /app                                                                                         0.0s
 => exporting to image                                                                                                0.0s
 => => exporting layers                                                                                               0.0s
 => => writing image sha256:dac18fa0ed143861760f071f5e5f38d1988dfd38040b996d963bbbf93c62cb6f                          0.0s
 => => naming to docker.io/library/golang-docker-compose-with-air_app                                                 0.0s
[+] Running 1/0
 ⠿ Container golang-docker-compose-with-air-app-1  Running                                                            0.0s
Attaching to golang-docker-compose-with-air-app-1
```

Then it need be basic configuration to allow live reloading code.