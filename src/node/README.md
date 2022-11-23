# node

`acicn/node` 基于 `acicn/debian:11`

## 标签

查阅 [Github Package 页面](https://github.com/guoyk93/acicn/pkgs/container/acicn%2Fnode)

## 运行环境

* 基于 `Debian 11`

* 内置 `minit`

    - 可以使用 `/etc/minit.d` 目录, `MINIT_MAIN` 环境变量 或者 `CMD` 指定要启动的进程
    - 支持一次性，配置文件渲染，定时任务等多个多种类型的进程
    - 内建 WebDAV 服务器，便于输出调试文件
    
    详细参考 https://github.com/acicn/minit

* npm 默认使用 Aliyun 镜像源

## 构建环境

`acicn/node:builder-xxx` 系列镜像额外安装了如下工具

* `常见编译工具`
* `cnpm`

## 默认配置

* 安装目录 `/opt/node`
* `/opt/node/bin` 已经加入 `$PATH` 环境变量
