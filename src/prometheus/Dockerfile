FROM {{Lookup .upstream}}

RUN curl -sSL -o prometheus.tar.gz {{.prometheus_download_url}} && \
    mkdir -p /opt/prometheus && \
    tar -C /opt/prometheus --strip-components 1 -xf prometheus.tar.gz && \
    rm -f prometheus.tar.gz && \
    mkdir /data

WORKDIR /opt/prometheus

ADD minit.d /etc/minit.d
ADD scripts /opt/bin

ADD prometheus.yml /opt/prometheus/prometheus.yml
