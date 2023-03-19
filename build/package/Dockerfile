FROM alpine:3.14

ARG GOARCH=arm64

ADD ./main_linux_$GOARCH /app/service
WORKDIR /app

CMD ["/app/service", "-config-dir", "/app/configs/", "-secret-dir", "/app/secrets/"]