FROM ghcr.io/guoyk93/acicn/debian:11-rc
# Must Debian
RUN curl -sSL -o node.tar.gz "https://nodejs.org/download/release/v10.24.1/node-v10.24.1-linux-x64.tar.gz" && \
mkdir -p /opt/node && \
tar -C /opt/node --strip-components 1 -xf node.tar.gz && \
rm -f node.tar.gz
ENV PATH "/opt/node/bin:${PATH}"
ENV NODE_VERSION 10
RUN npm config set unsafe-perm true && \
npm config set registry https://registry.npmmirror.com
ADD banner.minit.txt /etc/banner.minit.txt