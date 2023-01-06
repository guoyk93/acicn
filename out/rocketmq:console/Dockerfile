FROM ghcr.io/guoyk93/acicn/jdk:8-debian-11-rc
WORKDIR /opt/rocketmq
ADD console/minit.d /etc/minit.d
RUN curl -sSL -o console.jar https://github.com/guoyk93/acicn-assets/releases/download/v20221122/rocketmq-console-ng-2.0.0.jar
ADD banner.minit.txt /etc/banner.minit.txt