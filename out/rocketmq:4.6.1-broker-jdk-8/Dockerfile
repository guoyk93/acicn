FROM dev-build/acicn/jdk:8-debian-11
RUN curl -sSLo rocketmq.zip "http://archive.apache.org/dist/rocketmq/4.6.1/rocketmq-all-4.6.1-bin-release.zip" && \
unzip rocketmq.zip && \
rm -f rocketmq.zip && \
mv rocketmq-all-4.6.1-bin-release /opt/rocketmq
WORKDIR /opt/rocketmq
ENV ROCKETMQ_HOME /opt/rocketmq
ADD broker/minit.d /etc/minit.d
RUN mkdir -p /data && ln -sf /data /root/store
ADD broker.conf /opt/rocketmq/conf/broker.conf
ENV RMQBROKERCFG_brokerClusterName  DefaultCluster
ENV RMQBROKERCFG_brokerName         broker-a
ENV RMQBROKERCFG_brokerId           0
ENV RMQBROKERCFG_deleteWhen         04
ENV RMQBROKERCFG_fileReservedTime   48
ENV RMQBROKERCFG_brokerRole         ASYNC_MASTER
ENV RMQBROKERCFG_flushDiskType      ASYNC_FLUSH