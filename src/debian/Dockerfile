{{if .no_minit}}
{{else}}
FROM {{.utils_image_minit}} AS minit
{{end}}

FROM debian:{{.version}}

ENV LANG zh_CN.UTF-8
ENV TZ Asia/Shanghai
ENV DEBIAN_FRONTEND noninteractive

RUN : && \
    {{if .apt_mirror}}
    sed -i "s/deb.debian.org/{{.apt_mirror}}/g" /etc/apt/sources.list && \
    sed -i "s/security.debian.org/{{.apt_mirror}}/g" /etc/apt/sources.list && \
    {{end}}
    apt-get update && \
    apt-get upgrade -y && \
    apt-get install -y locales locales-all tzdata ca-certificates vim curl procps net-tools unzip netcat && \
    rm -rf /var/lib/apt/lists/* && \
    ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

RUN mkdir -p /opt/bin
ENV PATH "/opt/bin:${PATH}"

{{if .no_minit}}
{{else}}
COPY --from=minit /minit /opt/bin/minit
ENTRYPOINT ["/opt/bin/minit"]
{{end}}
