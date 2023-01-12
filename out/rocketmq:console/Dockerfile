FROM dev-build/acicn/jdk:8-debian-11
WORKDIR /opt/rocketmq
ADD console/minit.d /etc/minit.d
RUN curl -sSL -o console.jar https://github.com/guoyk93/acicn-assets/releases/download/v20221122/rocketmq-console-ng-2.0.0.jar