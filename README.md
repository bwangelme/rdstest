Redis 备份恢复的实验
===

## 目的

redis 支持一边写入，一边恢复

> 注意: aof 和 rdb 同时存在的时候，先恢复 aof

## 所需环境

redis server
客户端随机写入 key ， qps 50

## 操作步骤

0. 开启 rdb 和 aof ，启动redis
1. 向 Redis 写入一个独立的 key
2. 备份 rdb 文件
3. 删除 rdb 和 aof 文件
4. 关闭 redis server
5. 关闭 aof，复制 rdb 文件 ，启动 redis
6. 查看备份是否恢复成功
7. 热更新配置，开启 aof
8. 开启 aof 配置，重启 redis