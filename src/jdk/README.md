# jdk

## 标签

**运行环境**

* `acicn/jdk:8`
* `acicn/jdk:11`

**构建环境**

* `acicn/jdk:builder-8`
* `acicn/jdk:builder-11`

## 运行环境

* 基于 `Debian 11` 使用管理器安装 AdoptOpenJDK (Hotspot) 8/11

* 内置 `minit`

    - 可以使用 `/etc/minit.d` 目录, `MINIT_MAIN` 环境变量 或者 `CMD` 指定要启动的进程
    - 支持一次性，配置文件渲染，定时任务等多个多种类型的进程
    - 内建 WebDAV 服务器，便于输出调试文件
    
    详细参考 https://github.com/acicn/minit

* 内置 `Alibaba Arthas`

    **注意，使用 `Arthas` 调试可能需要为容器提供内核权限**

    可以直接执行 `as.sh` 启动

* `java-wrapper`

    镜像内置脚本 `java-wrapper` ，可以用以代替 `java` 命令，具备以下功能

    - 支持 `JAVA_OPTS` 环境变量

         `JAVA_OPTS` 和 **任何以 `JAVA_OPTS_` 开头的环境变量**，都会被扩展到 `java` 命令上

         建议的用法:

         `JAVA_OPTS_HEAP` 用于堆配置参数

         `JAVA_OPTS_HEAP=-Xms1g -Xmx4g`

         `JAVA_OPTS_GC` 用于内存回收配置参数

         `JAVA_OPTS_GC=-XX:+UseG1GC`

         当然你也可以一股脑把所有参数都放在 `JAVA_OPTS` 环境变量里

    - 兼容旧的 `JAVA_MEMORY_MAX`, `JAVA_MEMROY_MIN`, `JAVA_XMX` 和 `JAVA_XMS` 环境变量

* JMX Prometheus JavaAgent

    镜像内置 JMX Prometheus JavaAgent JAR，使用以下参数可以使用默认简单配置

    ```shell
    # 方法一，直接追加参数
    -javaagent:/opt/lib/jmx_prometheus_javaagent.jar=60030:/opt/etc/jmx_prometheus_javaagent/simple.yml
    # 方法二，使用 java-wrapper，并使用自定义 JAVA_OPTS_XXXXX 环境变量
    JAVA_OPTS_METRICS=-javaagent:/opt/lib/jmx_prometheus_javaagent.jar=60030:/opt/etc/jmx_prometheus_javaagent/simple.yml
    ```

* Skywalking JavaAgent

    镜像内置 Skywalking JavaAgent

    * 部署目录: `/opt/lib/skywalking-agent`

    使用以下参数可以使用 `skywalking-agent.jar`

    ```shell
    # 方法一，直接追加参数
    -javaagent:/opt/lib/skywalking-agent/skywalking-agent.jar
    # 方法二，使用 java-wrapper，并使用自定义 JAVA_OPTS_XXXXX 环境变量
    JAVA_OPTS_TRACING=-javaagent:/opt/lib/skywalking-agent/skywalking-agent.jar
    ```

    * 配置参数

    参考文档 https://skywalking.apache.org/docs/skywalking-java/latest/en/setup/service-agent/java-agent/configurations/

    * 其他插件
        * `/opt/lib/skywalking-agent/optional-plugins/apm-seata-skywalking-plugin.jar`
 
## 用法实例

``` dockerfile
FROM acicn/jdk:11
WORKDIR /work
ADD target/ms-id.jar ms-id.jar

ENV SPRING_PROFILE "test"

# 把启动项目非必须的调优参数放在这里，供未来在 Kubernetes 管理台上动态调整
ENV JAVA_XMS    "1g"
ENV JAVA_XMX    "1g"
ENV JAVA_OPTS   "-XX:+UseG1GC"

# 把启动项目必要的 Java 参数放在这里，比如 "-cp" 和 "-Dspring.profiles.active=${SPRING_PROFILE}" 参数
# minit 已经占用 ENTRYPOINT
CMD ["java-wrapper", "-cp", ".:./lib/*", "-Dspring.profiles.active=${SPRING_PROFILE}", "-jar", "ms-id.jar"]
```

## 构建环境

`acicn/jdk:builder-xxx` 系列镜像额外安装了如下工具

* `maven`
