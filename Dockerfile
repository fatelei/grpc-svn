FROM golang:1.14

WORKDIR /go/src/app
COPY . .
COPY sources.list   /etc/apt/sources.list

RUN apt update
RUN apt install -y apt-transport-https ca-certificates git-svn protobuf-compiler vim
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN make server
CMD ["bin/git_service"]
