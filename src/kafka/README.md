# kafka

## 标签

查阅 [Github Package 页面](https://github.com/guoyk93/acicn/pkgs/container/acicn%2Fkafka)

## 功能

* 内置 `minit`

    - 可以使用 `/etc/minit.d` 目录, `MINIT_MAIN` 环境变量 或者 `CMD` 指定要启动的进程
    - 支持一次性，配置文件渲染，定时任务等多个多种类型的进程
    - 内建 WebDAV 服务器，便于输出调试文件

    详细参考 https://github.com/acicn/minit

* 基于环境变量的 `server.properties` 渲染

    - 以 `KAFKACFG_` 开头的环境变量都会以以下规则转换为 `zoo.cfg` 文件中的配置
        - `__` 替换为 `.`

    比如，以下环境变量

    ```
    KAFKACFG_zookeeper__connect=zoo-1:2181,zoo-2:2181
    ```

    会渲染为

    ```properties
    zookeeper.connect=zoo-1:2181,zoo-2:2181
    ```

    - 在 Kubernetes 环境下，使用环境变量 `KAFKAAUTOCFG_PORT=9092` 和 `KAFKAAUTOCFG_ADVERTISED_HOST=my-service` 两个变量自动生成值

    ```
    listeners=PLAINTEXT://:9092
    advertised.listeners=PLAINTEXT://x.x.x.x:9092
    ```

## 默认配置

* 工作目录 `/opt/kafka`
* 数据目录 `/data`
