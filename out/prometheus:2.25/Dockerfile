FROM ghcr.io/guoyk93/acicn/debian:11
RUN curl -sSL -o prometheus.tar.gz https://github.com/prometheus/prometheus/releases/download/v2.25.0/prometheus-2.25.0.linux-amd64.tar.gz && \
mkdir -p /opt/prometheus && \
tar -C /opt/prometheus --strip-components 1 -xf prometheus.tar.gz && \
rm -f prometheus.tar.gz && \
mkdir /data
WORKDIR /opt/prometheus
ADD minit.d /etc/minit.d
ADD scripts /opt/bin
ADD prometheus.yml /opt/prometheus/prometheus.yml
ADD banner.minit.txt /etc/banner.minit.txt