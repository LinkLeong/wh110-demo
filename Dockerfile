FROM alpine:latest
# VOLUME /var/appogs/live.wikifx.com:/logs
WORKDIR /
RUN mkdir /mydir
ADD ./* /mydir/
RUN chmod +x /mydir/wh110api
LABEL version="1.0.15" description="描述说明"
EXPOSE 8001
ENTRYPOINT ["/mydir/wh110api"]