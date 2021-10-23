Redis 备份恢复的实验
===

## 目的

redis 支持一边写入，一边恢复

> 注意: aof 和 rdb 同时存在的时候，先恢复 aof

## 所需环境

redis server
客户端随机写入 key ， qps 50

## 操作步骤

+ 开启 rdb 和 aof ，启动redis

```
redis-server redis/redis_6379.conf
```

+ 向 Redis 写入一个独立的 key

```
go run bench.go
```

1. 备份 rdb 文件
1. 删除 rdb 和 aof 文件
1. 关闭 redis server
1. 关闭 aof，复制 rdb 文件 ，启动 redis
1. 查看备份是否恢复成功
1. 热更新配置，开启 aof `CONFIG SET appendonly yes`
1. 开启 aof 配置，重启 redis