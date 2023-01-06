FROM ghcr.io/guoyk93/acicn/debian:11-rc
ENV ALINODE_VERSION 5.20.6
ENV ENABLE_NODE_LOG YES
ENV NODE_LOG_DIR    /var/log/agenthub
ENV PATH "/opt/alinode/bin:${PATH}"
RUN curl -sSL -o alinode.tar.gz 'https://github.com/guoyk93/acicn-assets/releases/download/v20221122/alinode-v5.20.6-linux-x64.tar.gz' && \
mkdir -p /opt/alinode && \
tar -xf alinode.tar.gz -C /opt/alinode --strip-components=1 && \
rm -rf alinode.tar.gz && \
export ENABLE_NODE_LOG=NO && \
npm config set registry https://registry.npmjs.org && \
npm config set unsafe-perm true && \
npm install -g cnpm@7.1.1 @alicloud/agenthub yarn && \
yarn config set cache-folder $HOME/.npm/yarn && \
yarn config set registry https://registry.npmmirror.com && \
npm cache clean --force && \
mkdir -p $NODE_LOG_DIR
ADD agenthub      /opt/agenthub
RUN curl -sSL -o filebeat.tar.gz "https://artifacts.elastic.co/downloads/beats/filebeat/filebeat-7.2.1-linux-x86_64.tar.gz" && \
mkdir -p /opt/sensordata-filebeat && \
tar -xf filebeat.tar.gz -C /opt/sensordata-filebeat --strip-components=1 && \
rm -rf filebeat.tar.gz
ADD filebeat.yml  /opt/sensordata-filebeat/filebeat.yml
ADD minit.d       /etc/minit.d
ADD banner.minit.txt /etc/banner.minit.txt