FROM alpine
MAINTAINER clemenko@gmail.com
LABEL RUN="docker rum -d -v /tmp/redis:/data clemenko/redis --appendonly yes" \
  VERSION="3.0.2"
RUN apk update && apk add redis &&\
    sed -i -e 's/daemonize yes/daemonize no/g' /etc/redis.conf -e 's/# bind 0.0.0.0/bind 0.0.0.0/g' /etc/redis.conf &&\
    rm -rf /var/cache/apk/*
WORKDIR /data
ENTRYPOINT ["/usr/bin/redis-server"]
EXPOSE 6379
