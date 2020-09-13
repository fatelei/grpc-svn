FROM golang:1.14

WORKDIR /go/src/app
COPY . .
COPY sources.list   /etc/apt/sources.list

RUN apt update
RUN apt install -y apt-transport-https ca-certificates
RUN apt install -y git-svn
RUN apt install -y protobuf-compiler
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go get -d -v ./...
RUN go install -v ./...

CMD ["app"]
