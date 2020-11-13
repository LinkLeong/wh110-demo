FROM alpine:latest
# VOLUME /var/appogs/live.wikifx.com:/logs
WORKDIR /
ADD / /
LABEL version="1.0.10" description="描述说明"
EXPOSE 8001
ENTRYPOINT ["./wh110api"]