FROM debian:11
ENV LANG zh_CN.UTF-8
ENV TZ Asia/Shanghai
ENV DEBIAN_FRONTEND noninteractive
RUN : && \
apt-get update && \
apt-get upgrade -y && \
apt-get install -y locales locales-all tzdata ca-certificates vim curl procps net-tools unzip netcat && \
rm -rf /var/lib/apt/lists/* && \
ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
RUN mkdir -p /opt/bin
ENV PATH "/opt/bin:${PATH}"
ADD banner.minit.txt /etc/banner.minit.txt