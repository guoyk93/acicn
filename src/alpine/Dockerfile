{{if .no_minit}}
{{else}}
FROM {{.utils_image_minit}} AS minit
{{end}}

FROM alpine:{{.version}}

ENV LANG zh_CN.UTF-8
ENV TZ Asia/Shanghai

RUN : && \
    {{if .apk_mirror}}
    sed -i "s/dl-cdn.alpinelinux.org/{{.apk_mirror}}/g" /etc/apk/repositories && \
    {{end}}
    apk upgrade --no-cache && \
    apk add --no-cache coreutils tzdata ca-certificates vim curl && \
    ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && \
    echo $TZ > /etc/timezone && \
    echo $TZ > /etc/TZ

RUN mkdir -p /opt/bin
ENV PATH "/opt/bin:${PATH}"

{{if .no_minit}}
{{else}}
COPY --from=minit /minit /opt/bin/minit
ENTRYPOINT ["/opt/bin/minit"]
{{end}}
