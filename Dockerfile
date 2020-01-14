#构建复杂， 需执行编译动作，体积小
FROM golang:alpine
WORKDIR /go/src
ADD . /go/src
RUN cd /go/src && go build -o cstwebapp
EXPOSE 8888
ENTRYPOINT ./cstwebapp
