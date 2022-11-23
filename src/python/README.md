# python

## 标签

查阅 [Github Package 页面](https://github.com/guoyk93/acicn/pkgs/container/acicn%2Fpython)

## 功能

* 内置 `minit`

    - 可以使用 `/etc/minit.d` 目录, `MINIT_MAIN` 环境变量 或者 `CMD` 指定要启动的进程
    - 支持一次性，配置文件渲染，定时任务等多个多种类型的进程
    - 内建 WebDAV 服务器，便于输出调试文件

    详细参考 https://github.com/acicn/minit

* 使用编译安装 Python 各个版本

* PIP 默认使用 Aliyun 源

* 内置脚本 `venv-wrapper`

    一个辅助脚本，会先尝试 `source venv/bin/activate` 然后再执行后续动作

    常见用法如下

    ```dockerfile
    FROM acicn/python:3.8

    WORKDIR /work

    ADD requirements.txt requirements.txt

    RUN python3 -m venv venv && \
        venv-wrapper pip install -r requirements.txt

    ADD . .
    
    # minit 已经占用 ENTRYPOINT 位置
    CMD ["venv-wrapper", "python", "main.py"]
    ```
