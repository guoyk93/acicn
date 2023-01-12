FROM dev-build/acicn/alpine:3.16
RUN apk add --no-cache beanstalkd
WORKDIR /data
ADD minit.d /etc/minit.d
ENV BEANSTALKD_OPTS ""