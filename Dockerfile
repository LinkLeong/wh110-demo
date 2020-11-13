FROM alpine:latest
# VOLUME /var/appogs/live.wikifx.com:/logs
WORKDIR /
ADD / /
RUN chmod +x wh110api
LABEL version="1.0.19" description="描述说明"
EXPOSE 8001
ENTRYPOINT ["./wh110api"]