FROM ghcr.io/guoyk93/merge-env-to-ini:1.0.0 AS utils-merge-env-to-ini
FROM ghcr.io/guoyk93/acicn/nginx:alpine-3.16-rc AS utils-nginx
FROM ghcr.io/guoyk93/acicn/debian:11-rc
# Only Debian
ENV PHP_VERSION 7.4
ENV PHPCFG_PATH_PHP_FPM_CONF        /etc/php-fpm.conf
ENV PHPCFG_PATH_PHP_FPM_WWW_CONF    /etc/php-fpm.d/www.conf
ENV PHPCFG_PATH_PHP_INI             /etc/php.ini
RUN apt-get update && \
apt-get install -y nginx build-essential autoconf cmake cmake-extras openssl libssl-dev libsodium-dev \
libxml2-dev libsqlite3-dev libgd-dev libxpm-dev libicu-dev libbz2-dev libcurl4-openssl-dev \
libenchant-2-dev libgmp-dev libldap2-dev libaspell-dev libreadline-dev libsnmp-dev \
libxslt1-dev libgeoip-dev imagemagick libmemcached-dev libevent-dev liblz4-dev \
librdkafka-dev libpq-dev libyaml-dev libzip-dev libtidy-dev libmcrypt-dev libonig-dev \
libmariadb-dev libpcre3-dev libpspell-dev libldap2-dev libsasl2-dev libmagick++-dev \
subversion libsvn-dev && \
rm -rf /var/lib/apt/lists/*
RUN apt-get update && \
apt-get install -y libhunspell-dev && \
rm -rf /var/lib/apt/lists/* && \
curl -sSLo enchant.tar.gz 'https://github.com/abiword/enchant/releases/download/enchant-1-6-1/enchant-1.6.1.tar.gz' && \
mkdir -p /opt/src/enchant && \
tar -xf enchant.tar.gz -C /opt/src/enchant --strip-components 1 && \
rm -f enchant.tar.gz && \
cd /opt/src/enchant && \
./configure && \
make && \
make install
RUN mkdir -p /opt/src/php && \
curl -sSL -o php.tar.gz 'https://www.php.net/distributions/php-7.4.23.tar.gz' && \
tar -xf php.tar.gz -C /opt/src/php --strip-components 1 && \
rm -f php.tar.gz && \
cd /opt/src/php && \
./configure --prefix=/opt --enable-fpm --enable-gd --with-external-gd \
--with-config-file-path=/etc --with-config-file-scan-dir=/etc/php.d \
--with-jpeg --with-webp --with-xpm --with-freetype --enable-bcmath --with-zlib \
--with-pcre-jit --enable-sockets --enable-soap --with-pdo-mysql --with-mysql-sock \
--with-mysqli --enable-mysqlnd --with-openssl --with-zip --enable-mbstring \
--enable-intl --enable-pcntl --with-bz2 --enable-calendar --enable-sysvmsg \
--enable-sysvsem --enable-sysvshm --with-curl --enable-ftp --with-enchant \
--enable-exif --with-gmp --with-gettext --with-sodium --enable-dba --with-pspell \
--with-readline --with-snmp --with-xsl --enable-shmop --with-xmlrpc \
--with-tidy --with-pgsql --with-pdo-pgsql --enable-phpdbg \
--enable-phpdbg-webhelper --enable-phpdbg-readline --enable-sigchild \
--with-pear --with-ldap --with-ldap-sasl && \
make && make install
ENV PATH "/opt/sbin:$PATH"
ADD scripts /opt/bin
RUN mkdir -p /var/log/php-fpm
ADD 7.4/etc /etc
#################################
# BEGIN: Skywalking Agent       #
#################################
# install dependencies
RUN apt-get update && \
apt-get install -y gcc make libclang-dev protobuf-compiler && \
rm -rf /var/lib/apt/lists/*
# install rust manually
RUN mkdir -p /opt/src/rust && \
curl -sSL -o rust.tar.gz 'https://static.rust-lang.org/dist/rust-1.65.0-x86_64-unknown-linux-gnu.tar.gz' && \
tar -xf rust.tar.gz -C /opt/src/rust --strip-components 1 && \
rm -f rust.tar.gz && \
cd /opt/src/rust && \
./install.sh --prefix=/opt && \
cd / && \
rm -rf /opt/src/rust
# instlal skywalking-agent
RUN acicn-php-ext-install php-skywalking-agent 'https://github.com/apache/skywalking-php/archive/refs/tags/v0.2.0.tar.gz'
# environment variables for php skywalking agent
# SW_APP_CODE is skywalking_agent.service_name
ENV SW_APP_CODE             "demo"
ENV SW_ENABLE               "0"
ENV SW_LOG_FILE             "/var/log/skywalking-agent.log"
ENV SW_LOG_LEVEL            "INFO"
ENV SW_OAP_ADDRESS          "skywalking-oap:11800"
ENV SW_PROTOCOL_VERSION     "8"
ENV SW_WORKER_THREADS       "3"
# deprecated skywalking environment variables
ENV SW_ERROR_HANDLER_ENABLE "0"
ENV SW_SAMPLE_N_PER_3_SECS  "-1"
ENV SW_INSTANCE_NAME        ""
#################################
# END: Skywalking Agent         #
#################################
RUN acicn-php-ext-install php-apcu 'https://pecl.php.net/get/apcu-5.1.20.tgz' && \
acicn-php-ext-install php-mcrypt 'https://pecl.php.net/get/mcrypt-1.0.4.tgz' && \
acicn-php-ext-install php-igbinary 'https://pecl.php.net/get/igbinary-3.2.1.tgz' && \
acicn-php-ext-install php-memcached 'https://pecl.php.net/get/memcached-3.1.5.tgz' \
--enable-memcached \
--enable-memcached-session \
--enable-memcached-igbinary \
--enable-memcached-json && \
acicn-php-ext-install php-memcache 'https://pecl.php.net/get/memcache-4.0.5.2.tgz' && \
acicn-php-ext-install php-mongodb 'https://pecl.php.net/get/mongodb-1.5.5.tgz' \
--enable-mongodb && \
acicn-php-ext-install php-redis 'https://pecl.php.net/get/redis-5.3.4.tgz' \
--enable-redis-igbinary \
--enable-redis-lzf \
--enable-redis-lz4 \
--with-liblz4 && \
acicn-php-ext-install php-rdkafka 'https://pecl.php.net/get/rdkafka-5.0.0.tgz' \
--with-rdkafka && \
acicn-php-ext-install php-swoole 'https://pecl.php.net/get/swoole-4.4.23.tgz' \
--enable-sockets \
--enable-openssl \
--enable-http2 \
--enable-mysqlnd && \
acicn-php-ext-install php-geoip 'https://pecl.php.net/get/geoip-1.1.1.tgz' \
--with-geoip && \
acicn-php-ext-install php-imagick 'https://pecl.php.net/get/imagick-3.4.4.tgz' && \
acicn-php-ext-install php-yaml 'https://pecl.php.net/get/yaml-2.2.1.tgz' && \
acicn-php-ext-install php-svn 'https://pecl.php.net/get/svn-2.0.3.tgz'
# yac
RUN acicn-php-ext-install php-yac 'https://pecl.php.net/get/yac-2.3.1.tgz'
# Oracle InstantClient
ENV ORACLE_HOME 	"/opt/oracle/instantclient_21_6"
ENV LD_LIBRARY_PATH 	"$ORACLE_HOME:$LD_LIBRARY_PATH"
ENV PATH 		"$ORACLE_HOME:$PATH"
RUN apt-get update && \
apt-get install -y libaio-dev && \
rm -rf /var/lib/apt/lists/*
RUN mkdir -p /opt/oracle && \
cd /opt/oracle && \
curl -sSL -o instantclient-basic.zip 'https://github.com/guoyk93/acicn-assets/releases/download/v20221122/6daa0912-instantclient-basic-linux.x64-21.6.0.0.0dbru.zip' && \
curl -sSL -o instantclient-sdk.zip 'https://github.com/guoyk93/acicn-assets/releases/download/v20221122/6ab2ccfe-instantclient-sdk-linux.x64-21.6.0.0.0dbru.zip' && \
unzip instantclient-basic.zip && \
unzip instantclient-sdk.zip && \
rm -f instantclient-basic.zip && \
rm -f instantclient-sdk.zip
RUN acicn-php-ext-install php-oci8 'https://github.com/guoyk93/acicn-assets/releases/download/v20221122/b9e7c89e-oci8-2.2.0.tgz' --with-oci8="shared,instantclient,$ORACLE_HOME" && \
acicn-php-ext-activate pdo_oci --with-pdo-oci="instantclient,$ORACLE_HOME" --with-oci8="instantclient,$ORACLE_HOME"
COPY --from=utils-merge-env-to-ini /merge-env-to-ini /opt/bin/merge-env-to-ini
RUN rm -rf /etc/nginx
COPY --from=utils-nginx /etc/nginx /etc/nginx
COPY --from=utils-nginx /etc/minit.d/nginx.yml /etc/minit.d/nginx.yml
ADD minit.d /etc/minit.d
ADD index.php /var/www/public/index.php
ADD php.conf /etc/nginx/default.conf.d/php.conf
ADD 7.4/extra/etc /etc
ENV NGXCFG_DEFAULT_EXTRA_INDEX index.php
ENV NGXCFG_SNIPPETS_ENABLE_SPA true
ENV NGXCFG_SNIPPETS_SPA_INDEX  "/index.php?\$query_string"
RUN mkdir -p /etc/nginx/default.fastcgi.d /var/lib/php/session
WORKDIR /var/www
ADD banner.minit.txt /etc/banner.minit.txt