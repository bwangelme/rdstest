## Redis Simple Cluster

+ 启动 redis

```shell
redis-server redis_6379.conf
redis-server redis_6380.conf
```

## 配置说明

+ 主节点设置 `requirepass` 开启密码认证
+ 从节点设置 `masterauth` 设主节点的认证密码

