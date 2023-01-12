FROM dev-build/acicn/debian:11-no-minit
# Must Debian
# JDK
RUN curl -sSLo jdk.deb 'https://packages.adoptium.net/artifactory/deb/pool/main/t/temurin-11/temurin-11-jdk_11.0.17.0.0+8_amd64.deb' && \
curl -sSLo jdk-ca.deb 'https://packages.adoptium.net/artifactory/deb/pool/main/a/adoptium-ca-certificates/adoptium-ca-certificates_1.0.0-1_all.deb' && \
apt-get update && \
apt-get install -y ./jdk-ca.deb ./jdk.deb fontconfig && \
rm -rf /var/lib/apt/lists/* && \
rm -f jdk.deb && \
rm -rf jdk-ca.deb
ENV LD_LIBRARY_PATH "/lib/x86_64-linux-gnu:/usr/lib/x86_64-linux-gnu:/usr/local/lib/x86_64-linux-gnu"
ADD scripts/java-wrapper /opt/bin/java-wrapper
ENV JRE_VERSION  11
ENV JDK_VERSION  11
ENV JAVA_VERSION 11
ENV JAVA_HOME /usr/lib/jvm/temurin-11-jdk-amd64
# ensure java home is ok
RUN cd "/usr/lib/jvm/temurin-11-jdk-amd64" && echo "JAVA_HOME: OK"
# Arthas
RUN mkdir -p /opt/arthas && \
curl -sSLo arthas.zip 'https://github.com/alibaba/arthas/releases/download/arthas-all-3.6.7/arthas-bin.zip' && \
unzip -d /opt/arthas arthas.zip && \
rm -rf arthas.zip && \
apt-get update && \
apt-get install -y telnet && \
rm -rf /var/lib/apt/lists/* && \
:
ADD scripts/as.sh /opt/bin/as.sh
ENV ARTHAS_VERSION 3.5
# Maven
RUN curl -sSLo maven.tar.gz 'https://archive.apache.org/dist/maven/maven-3/3.8.6/binaries/apache-maven-3.8.6-bin.tar.gz' && \
mkdir -p /opt/maven && \
tar -C /opt/maven --strip-components 1 -xf maven.tar.gz && \
rm -f maven.tar.gz
ENV PATH "/opt/maven/bin:${PATH}"
# JMX Prometheus JavaAgent
RUN mkdir -p /opt/lib && \
cd /opt/lib && \
curl -sSL -o jmx_prometheus_javaagent.jar "https://repo1.maven.org/maven2/io/prometheus/jmx/jmx_prometheus_javaagent/0.16.1/jmx_prometheus_javaagent-0.16.1.jar"
ADD etc/jmx_prometheus_javaagent /opt/etc/jmx_prometheus_javaagent
# Skywalking JavaAgent
RUN mkdir -p /opt/lib/skywalking-agent && \
curl -sSL -o skywalking.tar.gz "https://github.com/guoyk93/acicn-assets/releases/download/v20221122/skywalking-java-agent-8.7.0-mod.tar.gz" && \
tar -xf skywalking.tar.gz --strip-components 2 -C /opt/lib/skywalking-agent && \
rm -f skywalking.tar.gz
RUN cd /opt/lib/skywalking-agent && \
cp -rf bootstrap-plugins/*.jar plugins/ && \
cd optional-plugins && \
cp -rf apm-mybatis-3.x-plugin-* apm-trace-ignore-plugin-* apm-guava-cache-plugin-* ../plugins/
RUN cd /opt/lib/skywalking-agent/plugins && rm -rf apm-jdk-http-plugin-*
RUN curl -sSL -o '/opt/lib/skywalking-agent/optional-plugins/apm-seata-skywalking-plugin.jar' 'https://github.com/guoyk93/acicn-assets/releases/download/v20221122/apm-seata-skywalking-plugin-mod.jar'
ADD opt/skywalking-agent/config /opt/lib/skywalking-agent/config