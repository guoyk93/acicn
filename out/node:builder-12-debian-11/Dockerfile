FROM dev-build/acicn/debian:11-no-minit
# Must Debian
RUN apt-get update && \
apt-get install -y build-essential git python3 python3-dev python3-venv python3-pip && \
rm -rf /var/lib/apt/lists/*
ADD pip.conf /root/.pip/pip.conf
RUN curl -sSL -o node.tar.gz "https://nodejs.org/download/release/v12.22.9/node-v12.22.9-linux-x64.tar.gz" && \
mkdir -p /opt/node && \
tar -C /opt/node --strip-components 1 -xf node.tar.gz && \
rm -f node.tar.gz
ENV PATH "/opt/node/bin:${PATH}"
ENV NODE_VERSION 12
RUN npm config set unsafe-perm true && \
npm install -g cnpm@7.1.1 yarn && \
npm cache clean -f && \
npm config set registry https://registry.npmmirror.com && \
cnpm config set unsafe-perm true && \
yarn config set cache-folder $HOME/.npm/yarn && \
yarn config set registry https://registry.npmmirror.com