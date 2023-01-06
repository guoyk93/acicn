FROM ghcr.io/guoyk93/acicn/jdk:8-debian-11
RUN curl -sSLo rocketmq.zip "http://archive.apache.org/dist/rocketmq/4.6.1/rocketmq-all-4.6.1-bin-release.zip" && \
unzip rocketmq.zip && \
rm -f rocketmq.zip && \
mv rocketmq-all-4.6.1-bin-release /opt/rocketmq
WORKDIR /opt/rocketmq
ENV ROCKETMQ_HOME /opt/rocketmq
ADD namesrv/minit.d /etc/minit.d
ADD banner.minit.txt /etc/banner.minit.txt