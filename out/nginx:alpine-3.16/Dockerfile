FROM ghcr.io/guoyk93/acicn/alpine:3.16-rc
RUN apk add --no-cache nginx && \
rm -rf /etc/nginx
ADD nginx   /etc/nginx
ADD minit.d /etc/minit.d
WORKDIR /var/www/public
ADD index.html  index.html
ENV NGXCFG_SNIPPETS_SPA_INDEX /index.html
ADD banner.minit.txt /etc/banner.minit.txt