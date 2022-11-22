FROM {{Known .upstream}}

# Must Debian
{{if .debian_version}}
{{else}}
BAD
{{end}}

# JDK
{{if .jdk_version }}

{{if .debian_version}}

RUN curl -sSLo jdk.deb '{{.jdk_deb_download_url}}' && \
    curl -sSLo jdk-ca.deb '{{.jdk_ca_deb_download_url}}' && \
    apt-get update && \
    apt-get install -y ./jdk-ca.deb ./jdk.deb fontconfig && \
    rm -rf /var/lib/apt/lists/* && \
    rm -f jdk.deb && \
    rm -rf jdk-ca.deb

ENV LD_LIBRARY_PATH "/lib/x86_64-linux-gnu:/usr/lib/x86_64-linux-gnu:/usr/local/lib/x86_64-linux-gnu"

{{end}}

ADD scripts/java-wrapper /opt/bin/java-wrapper

ENV JRE_VERSION  {{.jdk_version}}
ENV JDK_VERSION  {{.jdk_version}}
ENV JAVA_VERSION {{.jdk_version}}
ENV JAVA_HOME {{.java_home}}
{{end}}

# Arthas
{{if .arthas_version }}
RUN mkdir -p /opt/arthas && \
    curl -sSLo arthas.zip '{{.arthas_download_url}}' && \
    unzip -d /opt/arthas arthas.zip && \
    rm -rf arthas.zip && \
    {{if .debian_version}}
    apt-get update && \
    apt-get install -y telnet && \
    rm -rf /var/lib/apt/lists/* && \
    {{end}}
    :

ADD scripts/as.sh /opt/bin/as.sh

ENV ARTHAS_VERSION {{.arthas_version}}
{{end}}

# Maven
{{if .maven_version}}
RUN curl -sSLo maven.tar.gz '{{.maven_download_url}}' && \
    mkdir -p /opt/maven && \
    tar -C /opt/maven --strip-components 1 -xf maven.tar.gz && \
    rm -f maven.tar.gz

ENV PATH "/opt/maven/bin:${PATH}"
{{end}}

# JMX Prometheus JavaAgent
{{if .jmx_prometheus_javaagent}}
RUN mkdir -p /opt/lib && \
    cd /opt/lib && \
    curl -sSL -o jmx_prometheus_javaagent.jar "{{.jmx_prometheus_javaagent_download_url}}"

ADD etc/jmx_prometheus_javaagent /opt/etc/jmx_prometheus_javaagent
{{end}}

# Skywalking JavaAgent
{{if .skywalking_javaagent}}
RUN mkdir -p /opt/lib/skywalking-agent && \
    curl -sSL -o skywalking.tar.gz "{{.skywalking_javaagent_download_url}}" && \
    tar -xf skywalking.tar.gz --strip-components 2 -C /opt/lib/skywalking-agent && \
    rm -f skywalking.tar.gz
{{end}}

{{if .pagoda}}
RUN cd /opt/lib/skywalking-agent && \
    cp -rf bootstrap-plugins/*.jar plugins/ && \
    cd optional-plugins && \
    cp -rf apm-mybatis-3.x-plugin-* apm-trace-ignore-plugin-* apm-guava-cache-plugin-* ../plugins/

{{if eq .jdk_version 11}}
RUN cd /opt/lib/skywalking-agent/plugins && rm -rf apm-jdk-http-plugin-*
{{end}}

{{if eq .jdk_version 17}}
RUN cd /opt/lib/skywalking-agent/plugins && rm -rf apm-jdk-http-plugin-*
{{end}}

RUN curl -sSL -o '/opt/lib/skywalking-agent/optional-plugins/apm-seata-skywalking-plugin.jar' '{{.apm_seata_skywalking_plugin_download_url}}'

ADD pagoda/skywalking-agent/config /opt/lib/skywalking-agent/config

{{if .cubic}}
RUN curl -sSL -o cubic.tar.gz '{{.cubic_agent_download_url}}' && \
    mkdir -p /opt/lib/cubic-agent && \
    tar -xf cubic.tar.gz -C /opt/lib/cubic-agent --strip-components 1 && \
    rm -f cubic.tar.gz
{{end}}
{{end}}
