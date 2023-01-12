FROM dev-build/acicn/debian:11
RUN apt-get update && \
apt-get install -y  build-essential linux-headers-generic \
zlib1g-dev libncurses5-dev libgdbm-dev libnss3-dev libssl-dev libreadline-dev \
libffi-dev libsqlite3-dev libmpdec-dev libexpat1-dev liblzma-dev libbz2-dev && \
rm -rf /var/lib/apt/lists/*
RUN curl -sSLo python.tar.gz 'https://www.python.org/ftp/python/3.10.8/Python-3.10.8.tgz' && \
mkdir -p /opt/src/python && \
tar -xf python.tar.gz -C /opt/src/python --strip-components 1 && \
rm -f python.tar.gz && \
cd /opt/src/python && \
./configure \
--prefix=/opt \
--sysconfdir=/etc \
--enable-shared \
--enable-profiling \
--enable-optimizations \
--enable-loadable-sqlite-extensions \
--enable-big-digits \
--enable-ipv6 \
--with-lto \
--with-computed-gotos \
--with-dbmliborder=gdbm:ndbm \
--with-system-expat \
--with-system-ffi \
--with-system-libmpdec \
LDFLAGS="-Wl,--rpath=/opt/lib" \
&& \
make && \
make install && \
rm -rf /opt/src
ADD scripts /opt/bin
ADD pip.conf /root/.pip/pip.conf