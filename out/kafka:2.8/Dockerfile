FROM ghcr.io/guoyk93/acicn/jdk:11-debian-11
RUN curl -sSL -o kafka.tar.gz "https://archive.apache.org/dist/kafka/2.8.1/kafka_2.13-2.8.1.tgz" && \
mkdir -p /opt/kafka && \
tar -C /opt/kafka --strip-components 1 -xf kafka.tar.gz && \
rm -rf kafka.tar.gz
ENV PATH "/opt/kafka/bin:${PATH}"
WORKDIR /opt/kafka
ADD config  config
ADD minit.d /etc/minit.d
ENV KAFKACFG_broker__id                 0
#ENV KAFKACFG_listeners                 PLAINTEXT://:9092
#ENV KAFKACFG_advertised__listeners     PLAINTEXT://kafka:9092
ENV KAFKACFG_num__network__threads      3
ENV KAFKACFG_num__io__threads           8
ENV KAFKACFG_socket__send__buffer__bytes    102400
ENV KAFKACFG_socket__receive__buffer__bytes 102400
ENV KAFKACFG_socket__request__max__bytes    104857600
ENV KAFKACFG_log__dirs                      /data
ENV KAFKACFG_num__partitions                1
ENV KAFKACFG_num__recovery__threads__per__data__dir         1
ENV KAFKACFG_offsets__topic__replication__factor            1
ENV KAFKACFG_transaction__state__log__replication__factor   1
ENV KAFKACFG_transaction__state__log__min__isr              1
#ENV KAFKACFG_log__flush__interval__messages    10000
#ENV KAFKACFG_log__flush__interval__ms          1000
ENV KAFKACFG_log__retention__hours  168
#ENV KAFKACFG_log__retention__bytes 1073741824
ENV KAFKACFG_log__segment__bytes 1073741824
ENV KAFKACFG_log__retention__check__interval__ms 300000
ENV KAFKACFG_zookeeper__connect                 zookeeper:2181
ENV KAFKACFG_zookeeper__connection__timeout__ms 18000
ENV KAFKACFG_group__initial__rebalance__delay__ms 0
ADD banner.minit.txt /etc/banner.minit.txt