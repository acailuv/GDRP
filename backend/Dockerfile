FROM golang:1.16.3

RUN mkdir -p /backend
COPY . /backend

WORKDIR /backend
RUN go get

EXPOSE 8080
EXPOSE 5000

CMD [ "go", "run", "main" ]