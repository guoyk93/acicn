FROM {{Lookup .upstream}}

# Must Debian
{{if .debian_version}}
{{else}}
BAD
{{end}}

{{if .debian_version }}
RUN apt-get update && \
    apt-get install -y libtcnative-1 && \
    rm -rf /var/lib/apt/lists/*
{{end}}

RUN curl -sSL -o tomcat.tar.gz "{{.tomcat_download_url}}" && \
    mkdir -p /usr/local/tomcat && \
    tar -C /usr/local/tomcat --strip-components 1 -xf tomcat.tar.gz && \
    rm -f tomcat.tar.gz && \
    rm -rf /usr/local/tomcat/webapps/* && \
    curl -sSL -o /usr/local/tomcat/webapps/ROOT.war '{{.empty_war_download_url}}'

WORKDIR /usr/local/tomcat

ENV PATH "/usr/local/tomcat/bin:${PATH}"

ADD conf    conf
ADD minit.d /etc/minit.d
ADD scripts /opt/bin

ENV TOMCAT_VERSION {{.tomcat_version}}

ENV TOMCATCFG_SERVER_CONNECTOR_port 8080
ENV TOMCATCFG_SERVER_CONNECTOR_protocol "HTTP/1.1"
ENV TOMCATCFG_SERVER_CONNECTOR_connectionTimeout 20000
ENV TOMCATCFG_SERVER_CONNECTOR_redirectPort 8443
