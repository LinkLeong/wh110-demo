FROM alpine:latest
# VOLUME /var/appogs/live.wikifx.com:/logs
WORKDIR /mybin
ADD ./* /
RUN chmod +x wh110api
LABEL version="1.0.17" description="描述说明"
EXPOSE 8001
ENTRYPOINT ["./wh110api"]