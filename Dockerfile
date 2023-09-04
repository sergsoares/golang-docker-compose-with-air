# https://hub.docker.com/_/golang/tags
FROM golang:1.21.0

# https://github.com/cosmtrek/air/releases
RUN go install github.com/cosmtrek/air@v1.45.0

WORKDIR /app
ENTRYPOINT ["air"]
