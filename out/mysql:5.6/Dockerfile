FROM mysql:5.6
ENV TZ=Asia/Shanghai
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone
ADD banner.minit.txt /etc/banner.minit.txt