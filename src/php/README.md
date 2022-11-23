# php

`acicn/php` 镜像基于 `debian:11` 镜像

## 标签

查阅 [Github Package 页面](https://github.com/guoyk93/acicn/pkgs/container/acicn%2Fphp)

## 功能

* 内置 `minit`

    - 可以使用 `/etc/minit.d` 目录, `MINIT_MAIN` 环境变量 或者 `CMD` 指定要启动的进程
    - 支持一次性，配置文件渲染，定时任务等多个多种类型的进程
    - 内建 WebDAV 服务器，便于输出调试文件

    详细参考 https://github.com/acicn/minit

* PHP 扩展 

    默认安装了一批常用扩展

* 使用 `merge-env-to-ini` 工具和环境变量修改 `PHP FPM` 配置文件

    详细参考 https://github.com/acicn/merge-env-to-ini

    * 环境变量前缀 `PHPCFG_PHP_INI_` 修改 `/etc/php.ini` 文件

        比如

        `PHPCFG_PHP_INI_aaaa__xxxxxxx=hello=world`

        会在 `/etc/php.ini` 文件中的 `[aaaa]` 分段，**增加或者修改**键值 `hello=world`，环境变量名中的 `__xxxx` 后缀会被忽略，用以防止字段名冲突

    * 环境变量前缀 `PHPCFG_PHP_FPM_CONF_` 修改 `/etc/php-fpm.conf` 文件

        比如

        `PHPCFG_PHP_FPM_CONF_aaaa__xxxxxxx=hello=world`

        会在 `/etc/php-fpm.conf` 文件中的 `[aaaa]` 分段，**增加或者修改**键值 `hello=world`，环境变量名中的 `__xxxx` 后缀会被忽略，用以防止字段名冲突

    * 环境变量前缀 `PHPCFG_PHP_FPM_WWW_CONF_` 修改 `/etc/php-fpm.d/www.conf` 文件

        比如

        `PHPCFG_PHP_FPM_WWW_CONF_aaaa__xxxxxxx=hello=world`

        会在 `/etc/php-fpm.d/www.conf` 文件中的 `[aaaa]` 分段，**增加或者修改**键值 `hello=world`，环境变量名中的 `__xxxx` 后缀会被忽略，用以防止字段名冲突

* `nginx`

    `nginx` 进程完全使用 `acicn/nginx` 的配置模式

    额外的修改

    - 使用文件 `/etc/nginx/default.conf.d/php.conf` 增加了 PHP 的支持（也就是 `location ~ \.php$ {` 区块）

    - 允许使用文件 `/etc/nginx/default.fastcgi.d/*.conf` 扩充上述区块的配置

    - 允许使用环境变量 `NGXCFG_DEFAULT_PHP_EXTRA_CONF` 扩充上述区块的配置

    - 默认启用 PHP 框架模式，即使用 `/var/www/public/index.php` 来统一处理所有路由
        - `NGXCFG_SNIPPETS_ENABLE_SPA=true`
        - `NGXCFG_SNIPPETS_SPA_INDEX=/index.php?$query_string`

## Pagoda （公司内部版本)

* Skywalking 配置

可以使用环境变量配置 `skywalking.ini`, 默认值如下

```
SW_APP_CODE=demo
SW_ENABLE=1
SW_OAP_ADDRESS=skywalking-oap:11800
SW_ERROR_HANDLER_ENABLE=0
SW_SAMPLE_N_PER_3_SECS=-1
SW_INSTANCE_NAME=容器主机名
```

## 默认配置

* PHP 项目地址 `/var/www/public`
