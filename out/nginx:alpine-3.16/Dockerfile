FROM dev-build/acicn/alpine:3.16
RUN apk add --no-cache nginx && \
rm -rf /etc/nginx
ADD nginx   /etc/nginx
ADD minit.d /etc/minit.d
WORKDIR /var/www/public
ADD index.html  index.html
ENV NGXCFG_SNIPPETS_SPA_INDEX /index.html