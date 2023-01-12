FROM alpine:3.16
ENV LANG zh_CN.UTF-8
ENV TZ Asia/Shanghai
RUN : && \
apk upgrade --no-cache && \
apk add --no-cache coreutils tzdata ca-certificates vim curl && \
ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && \
echo $TZ > /etc/timezone && \
echo $TZ > /etc/TZ
RUN mkdir -p /opt/bin
ENV PATH "/opt/bin:${PATH}"