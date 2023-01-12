FROM {{Lookup .upstream}}

ENV ALINODE_VERSION {{.alinode_version}}

ENV ENABLE_NODE_LOG YES
ENV NODE_LOG_DIR    /var/log/agenthub
ENV PATH "/opt/alinode/bin:${PATH}"

RUN curl -sSL -o alinode.tar.gz '{{.alinode_download_url}}' && \
    mkdir -p /opt/alinode && \
    tar -xf alinode.tar.gz -C /opt/alinode --strip-components=1 && \
    rm -rf alinode.tar.gz && \
    export ENABLE_NODE_LOG=NO && \
    npm config set registry https://registry.npmjs.org && \
    npm config set unsafe-perm true && \
    npm install -g cnpm@{{.cnpm_version}} @alicloud/agenthub yarn && \
    yarn config set cache-folder $HOME/.npm/yarn && \
    yarn config set registry https://registry.npmmirror.com && \
    npm cache clean --force && \
    mkdir -p $NODE_LOG_DIR

ADD agenthub      /opt/agenthub

RUN curl -sSL -o filebeat.tar.gz "{{.filebeat_download_url}}" && \
    mkdir -p /opt/sensordata-filebeat && \
    tar -xf filebeat.tar.gz -C /opt/sensordata-filebeat --strip-components=1 && \
    rm -rf filebeat.tar.gz

ADD filebeat.yml  /opt/sensordata-filebeat/filebeat.yml

ADD minit.d       /etc/minit.d