FROM ghcr.io/guoyk93/acicn/jdk:11-debian-11
# Must Debian
RUN apt-get update && \
apt-get install -y libtcnative-1 && \
rm -rf /var/lib/apt/lists/*
RUN curl -sSL -o tomcat.tar.gz "https://archive.apache.org/dist/tomcat/tomcat-9/v9.0.69/bin/apache-tomcat-9.0.69.tar.gz" && \
mkdir -p /usr/local/tomcat && \
tar -C /usr/local/tomcat --strip-components 1 -xf tomcat.tar.gz && \
rm -f tomcat.tar.gz && \
rm -rf /usr/local/tomcat/webapps/* && \
curl -sSL -o /usr/local/tomcat/webapps/ROOT.war 'https://github.com/guoyk93/emptywar/releases/download/v1.0.0/emptywar-1.0.0.war'
WORKDIR /usr/local/tomcat
ENV PATH "/usr/local/tomcat/bin:${PATH}"
ADD conf    conf
ADD minit.d /etc/minit.d
ADD scripts /opt/bin
ENV TOMCAT_VERSION 9
ENV TOMCATCFG_SERVER_CONNECTOR_port 8080
ENV TOMCATCFG_SERVER_CONNECTOR_protocol "HTTP/1.1"
ENV TOMCATCFG_SERVER_CONNECTOR_connectionTimeout 20000
ENV TOMCATCFG_SERVER_CONNECTOR_redirectPort 8443
ADD banner.minit.txt /etc/banner.minit.txt