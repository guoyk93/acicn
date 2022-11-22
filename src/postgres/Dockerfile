FROM postgres:{{.upstream_version}}

ENV TZ=Asia/Shanghai
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

{{if .pgagent}}
RUN : && \
    {{if .apt_mirror}}
    sed -i 's/deb.debian.org/{{.apt_mirror}}/g' /etc/apt/sources.list && \
    sed -i 's/security.debian.org/{{.apt_mirror}}/g' /etc/apt/sources.list && \
    {{end}}
    apt-get update && \
    apt-get install -y pgagent tzdata ca-certificates vim curl procps net-tools unzip telnet && \
    rm -rf /var/lib/apt/lists/*
{{end}}
