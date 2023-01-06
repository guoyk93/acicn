FROM ghcr.io/guoyk93/acicn/alpine:3.16
RUN apk add --no-cache beanstalkd
WORKDIR /data
ADD minit.d /etc/minit.d
ENV BEANSTALKD_OPTS ""
ADD banner.minit.txt /etc/banner.minit.txt