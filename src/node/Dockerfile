FROM {{Lookup .upstream}}

# Must Debian
{{if .debian_version}}
{{else}}
BAD
{{end}}

{{if .builder}}
{{if .debian_version}}
RUN apt-get update && \
    apt-get install -y build-essential git python3 python3-dev python3-venv python3-pip && \
    rm -rf /var/lib/apt/lists/*

ADD pip.conf /root/.pip/pip.conf
{{end}}
{{end}}

{{if .node_version}}
RUN curl -sSL -o node.tar.gz "{{.node_download_url}}" && \
    mkdir -p /opt/node && \
    tar -C /opt/node --strip-components 1 -xf node.tar.gz && \
    rm -f node.tar.gz

ENV PATH "/opt/node/bin:${PATH}"
ENV NODE_VERSION {{.node_version}}
{{end}}

{{if .builder}}
RUN npm config set unsafe-perm true && \
    npm install -g cnpm@{{.cnpm_version}} yarn && \
    npm cache clean -f && \
    npm config set registry https://registry.npmmirror.com && \
    cnpm config set unsafe-perm true && \
    yarn config set cache-folder $HOME/.npm/yarn && \
    yarn config set registry https://registry.npmmirror.com
{{else}}
RUN npm config set unsafe-perm true && \
    npm config set registry https://registry.npmmirror.com
{{end}}
