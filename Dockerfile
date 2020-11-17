FROM alpine:latest
# VOLUME /var/appogs/live.wikifx.com:/logs
WORKDIR /
ADD / /
RUN chmod +x wh110api
LABEL version="1.1.0" description="增加线上配置内容"
EXPOSE 8001
ENTRYPOINT ["./wh110api"]