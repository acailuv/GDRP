FROM golang:1.16.3

RUN mkdir -p /backend
COPY . /backend

WORKDIR /backend
RUN go get

EXPOSE 8080

# Hot Reload
RUN [ "go", "get", "github.com/githubnemo/CompileDaemon" ]
CMD CompileDaemon -log-prefix=false -build="go build" -command="./main"