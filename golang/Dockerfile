FROM golang:1.13

WORKDIR /app

COPY . .

RUN apt-get update  -y && apt-get install -y --no-install-recommends apt-utils

RUN go get github.com/githubnemo/CompileDaemon

RUN go build .
RUN chmod +x mars_rover
CMD ["./mars_rover"]
#ENTRYPOINT CompileDaemon --build="go build -o main" --command=./main