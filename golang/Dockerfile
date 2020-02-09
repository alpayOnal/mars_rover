FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY .. .

RUN apt-get update  -y && apt-get install -y --no-install-recommends apt-utils

RUN go get github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon --build="go build -o main" --command=./main