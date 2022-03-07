FROM golang:alpine AS builder

ENV GOOS linux
ENV GOPROXY https://goproxy.cn,direct

WORKDIR /webhooks

COPY . .

RUN go build -o webhooks

FROM alpine

RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime

RUN echo 'Asia/Shanghai' > /etc/timezone

WORKDIR /webhooks

COPY --from=builder /webhooks/webhooks /webhooks/webhooks

ENTRYPOINT ["./webhooks"]