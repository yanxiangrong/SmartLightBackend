FROM golang:1.17.2

ENV GOPROXY https://goproxy.cn,direct
WORKDIR $GOPATH/src/SmartLightBackend/
COPY . $GOPATH/src/SmartLightBackend/
RUN go build .

EXPOSE 9527
EXPOSE 9528
ENTRYPOINT ["./SmartLightBackend"]
