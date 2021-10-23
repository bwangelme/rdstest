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
2. 删除 rdb 和 aof 文件
3. 关闭 redis server
4. 关闭 aof，复制 rdb 文件 ，启动 redis
5. 查看备份是否恢复成功
6. 热更新配置，开启 aof `CONFIG SET appendonly yes`
7. 开启 aof 配置，重启 redis

## 结果

重启 redis 后，除了 redis-server 停掉那段时间，数据丢失之外，其他的数据是没有丢失的。
17 年在上家公司遇到的 redis 重启后数据丢失的情况，应该在启动的时候没有关闭 AOF，从 AOF 恢复数据导致数据丢失。
