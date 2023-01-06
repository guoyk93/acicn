FROM ghcr.io/guoyk93/acicn/jdk:8-debian-11-rc
ENV APOLLO_VERSION 1.3.0
RUN mkdir -p /apollo/portal && cd /apollo/portal && \
curl -sSL -o apollo-portal.zip "http://github.com/ctripcorp/apollo/releases/download/v1.3.0/apollo-portal-1.3.0-github.zip" && \
unzip apollo-portal.zip && rm -f apollo-portal.zip && \
mkdir -p /apollo/configservice && cd /apollo/configservice && \
curl -sSL -o apollo-configservice.zip "http://github.com/ctripcorp/apollo/releases/download/v1.3.0/apollo-configservice-1.3.0-github.zip" && \
unzip apollo-configservice.zip && rm -f apollo-configservice.zip && \
mkdir -p /apollo/adminservice && cd /apollo/adminservice && \
curl -sSL -o apollo-adminservice.zip "http://github.com/ctripcorp/apollo/releases/download/v1.3.0/apollo-adminservice-1.3.0-github.zip" && \
unzip apollo-adminservice.zip && rm -f apollo-adminservice.zip
ADD minit.d /etc/minit.d
ADD banner.minit.txt /etc/banner.minit.txt