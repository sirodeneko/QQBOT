FROM golang as build

ENV GOPROXY=https://goproxy.io

ADD . /QQBOT

WORKDIR /QQBOT

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o qq_bot

FROM alpine:3.7

RUN echo "http://mirrors.aliyun.com/alpine/v3.7/main/" > /etc/apk/repositories && \
    apk update && \
    apk add ca-certificates && \
    echo "hosts: files dns" > /etc/nsswitch.conf && \
    mkdir -p /www/conf

WORKDIR /www

COPY --from=build /QQBOT/qq_bot /usr/bin/qq_bot

RUN chmod +x /usr/bin/qq_bot

ENTRYPOINT ["qq_bot"]