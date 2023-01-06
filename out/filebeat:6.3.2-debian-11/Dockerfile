FROM ghcr.io/guoyk93/acicn/debian:11
RUN curl -sSL -o filebeat.tar.gz "https://artifacts.elastic.co/downloads/beats/filebeat/filebeat-6.3.2-linux-x86_64.tar.gz" && \
mkdir -p /usr/share/filebeat && \
tar xf filebeat.tar.gz -C /usr/share/filebeat --strip-components 1 && \
rm -f filebeat.tar.gz
WORKDIR /usr/share/filebeat
ADD minit.d /etc/minit.d
ADD banner.minit.txt /etc/banner.minit.txt