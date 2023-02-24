FROM golang:1.18-alpine as builder

WORKDIR /go/src/forum
COPY . .

RUN go env -w GO111MODULE=on \
    && go env -w GOPROXY="https://goproxy.io,https://goproxy.cn,https://goproxy.io/zh/,https://proxy.golang.com.cn,direct" \
    && go env -w CGO_ENABLED=0 \
    && go env \
    && go mod tidy \
    && go build -o forum .

FROM alpine:latest

LABEL MAINTAINER="CherryQll@cherryqcsk@gmail.com"

WORKDIR /go/src/forum

COPY --from=0 /go/src/forum ./
COPY --from=0 /go/src/forum/resource ./resource/
COPY --from=0 /go/src/forum/config.docker.yaml ./

EXPOSE 8889
ENTRYPOINT ./forum -c config.docker.yaml
