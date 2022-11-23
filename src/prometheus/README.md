# prometheus

为什么这帮人非要在容器内用非 root 用户，在容器内直接用 root 不好么？是不是脑子有问题？

## 标签

查阅 [Github Package 页面](https://github.com/guoyk93/acicn/pkgs/container/acicn%2Fprometheus)

## 功能

* 使用环境变量 `PROMETHEUS_OPTS` 添加额外的启动参数

## 配置 Kubernetes 权限

```yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: prometheus
rules:
- apiGroups: [""]
  resources:
  - nodes
  - nodes/metrics
  - services
  - endpoints
  - pods
  verbs: ["get", "list", "watch"]
- apiGroups:
  - extensions
  - networking.k8s.io
  resources:
  - ingresses
  verbs: ["get", "list", "watch"]
- nonResourceURLs: ["/metrics", "/metrics/cadvisor"]
  verbs: ["get"]
---
apiVersion: v1
kind: ServiceAccount
metadata:
  name: prometheus
  namespace: default
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: prometheus
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: prometheus
subjects:
- kind: ServiceAccount
  name: prometheus
  namespace: default
```

## 默认配置

* 配置文件 `/opt/prometheus/prometheus.yml`
* 数据目录 `/data`
