FROM dev-build/acicn/debian:11
# Must Debian
RUN curl -sSL -o node.tar.gz "https://nodejs.org/download/release/v8.17.0/node-v8.17.0-linux-x64.tar.gz" && \
mkdir -p /opt/node && \
tar -C /opt/node --strip-components 1 -xf node.tar.gz && \
rm -f node.tar.gz
ENV PATH "/opt/node/bin:${PATH}"
ENV NODE_VERSION 8
RUN npm config set unsafe-perm true && \
npm config set registry https://registry.npmmirror.com