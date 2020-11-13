FROM alpine:latest
# VOLUME /var/appogs/live.wikifx.com:/logs
RUN mkdir /mybin
ADD ./* /mybin/
RUN chmod +x /mydir/wh110api
WORKDIR /mybin
LABEL version="1.0.16" description="描述说明"
EXPOSE 8001
ENTRYPOINT ["./wh110api"]