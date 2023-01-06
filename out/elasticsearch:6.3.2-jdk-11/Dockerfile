FROM ghcr.io/guoyk93/acicn/jdk:11-debian-11
RUN curl -sSLo elasticsearch.tar.gz "https://artifacts.elastic.co/downloads/elasticsearch/elasticsearch-6.3.2.tar.gz" && \
mkdir -p /opt/elasticsearch && \
tar -C /opt/elasticsearch --strip-components 1 -xf elasticsearch.tar.gz && \
rm -f elasticsearch.tar.gz && \
/opt/elasticsearch/bin/elasticsearch-plugin install -b "https://github.com/medcl/elasticsearch-analysis-ik/releases/download/v6.3.2/elasticsearch-analysis-ik-6.3.2.zip" && \
curl -sSLo /opt/elasticsearch-tune-1.0.jar "https://github.com/guoyk93/elasticsearch-tune/releases/download/v1.0/elasticsearch-tune-1.0.jar"
WORKDIR /opt/elasticsearch
ADD config config
ADD minit.d /etc/minit.d
ENV MINIT_SYSCTL vm.max_map_count=262144
ENV MINIT_RLIMIT_MEMLOCK unlimited
ENV ESJVM_XSS 1m
ENV ESJVM_XMS 1g
ENV ESJVM_XMX 1g
ENV ESCFG_path__data /data/data
ENV ESCFG_path__logs /data/logs
ENV ESCFG_network__host 0.0.0.0
ENV ESCFG_http__port 9200
ENV ESCFG_discovery__type single-node
ADD banner.minit.txt /etc/banner.minit.txt