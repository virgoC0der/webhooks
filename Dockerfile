FROM golang:latest

RUN go env -w GO111MODULE=on \
&& go env -w GOPROXY=https://goproxy.cn,direct

RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

RUN echo 'Asia/Shanghai' > /etc/timezone

WORKDIR /webhooks

ADD . /webhooks

RUN make build

ENTRYPOINT ["./webhooks"]