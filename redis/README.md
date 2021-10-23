## Redis Simple Cluster

+ 启动 redis

```shell
redis-server redis_6379.conf
redis-server redis_6380.conf
```

## 配置说明

+ 主节点设置 `requirepass` 开启密码认证
+ 从节点设置 `masterauth` 设主节点的认证密码

## 基准测试

+ 机器配置: 3.2Ghz 6核 i7  | 32G 内存
+ redis-benchmark 可修改的选项:

```
 -c <clients>       并行执行的客户端数 默认是 50
 -n <requests>      总请求数 默认是 100000
 -d <size>          SET/GET 值的数据大小，默认是3字节
 -t 测试集          用逗号分隔，测试集的名字在输出中可以看到，默认运行所有
```

在 summary 中可以看到每个操作支持的 QPS

```
ø> redis-benchmark -h localhost -a aKr84vEm5Ekd                                                                                                                                                                                            14:46:42 (10-23)
====== PING_INLINE ======
  100000 requests completed in 1.05 seconds
  50 parallel clients
  3 bytes payload
  keep alive: 1
  host configuration "save": 900 1 300 10 60 10000
  host configuration "appendonly": yes
  multi-thread: no

Latency by percentile distribution:
0.000% <= 0.151 milliseconds (cumulative count 2)
50.000% <= 0.271 milliseconds (cumulative count 51960)
75.000% <= 0.287 milliseconds (cumulative count 79744)
87.500% <= 0.303 milliseconds (cumulative count 88633)
93.750% <= 0.335 milliseconds (cumulative count 94435)
96.875% <= 0.375 milliseconds (cumulative count 97077)
98.438% <= 0.431 milliseconds (cumulative count 98564)
99.219% <= 0.487 milliseconds (cumulative count 99250)
99.609% <= 0.543 milliseconds (cumulative count 99614)
99.805% <= 0.607 milliseconds (cumulative count 99811)
99.902% <= 0.711 milliseconds (cumulative count 99907)
99.951% <= 0.895 milliseconds (cumulative count 99952)
99.976% <= 1.151 milliseconds (cumulative count 99976)
99.988% <= 1.479 milliseconds (cumulative count 99988)
99.994% <= 1.607 milliseconds (cumulative count 99994)
99.997% <= 1.639 milliseconds (cumulative count 99998)
99.998% <= 1.655 milliseconds (cumulative count 99999)
99.999% <= 1.671 milliseconds (cumulative count 100000)
100.000% <= 1.671 milliseconds (cumulative count 100000)

Cumulative distribution of latencies:
0.000% <= 0.103 milliseconds (cumulative count 0)
0.016% <= 0.207 milliseconds (cumulative count 16)
88.633% <= 0.303 milliseconds (cumulative count 88633)
98.064% <= 0.407 milliseconds (cumulative count 98064)
99.366% <= 0.503 milliseconds (cumulative count 99366)
99.811% <= 0.607 milliseconds (cumulative count 99811)
99.900% <= 0.703 milliseconds (cumulative count 99900)
99.935% <= 0.807 milliseconds (cumulative count 99935)
99.954% <= 0.903 milliseconds (cumulative count 99954)
99.962% <= 1.007 milliseconds (cumulative count 99962)
99.972% <= 1.103 milliseconds (cumulative count 99972)
99.978% <= 1.207 milliseconds (cumulative count 99978)
99.981% <= 1.303 milliseconds (cumulative count 99981)
99.985% <= 1.407 milliseconds (cumulative count 99985)
99.989% <= 1.503 milliseconds (cumulative count 99989)
99.994% <= 1.607 milliseconds (cumulative count 99994)
100.000% <= 1.703 milliseconds (cumulative count 100000)

Summary:
  throughput summary: 94966.77 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.279     0.144     0.271     0.343     0.463     1.671
====== PING_MBULK ======
  100000 requests completed in 1.08 seconds
  50 parallel clients
  3 bytes payload
  keep alive: 1
  host configuration "save": 900 1 300 10 60 10000
  host configuration "appendonly": yes
  multi-thread: no

Latency by percentile distribution:
0.000% <= 0.135 milliseconds (cumulative count 1)
50.000% <= 0.279 milliseconds (cumulative count 59548)
75.000% <= 0.287 milliseconds (cumulative count 78796)
87.500% <= 0.303 milliseconds (cumulative count 91284)
93.750% <= 0.319 milliseconds (cumulative count 95257)
96.875% <= 0.335 milliseconds (cumulative count 97019)
98.438% <= 0.367 milliseconds (cumulative count 98537)
99.219% <= 0.423 milliseconds (cumulative count 99264)
99.609% <= 0.487 milliseconds (cumulative count 99621)
99.805% <= 0.535 milliseconds (cumulative count 99816)
99.902% <= 0.583 milliseconds (cumulative count 99914)
99.951% <= 0.615 milliseconds (cumulative count 99957)
99.976% <= 0.647 milliseconds (cumulative count 99979)
99.988% <= 0.655 milliseconds (cumulative count 99988)
99.994% <= 0.671 milliseconds (cumulative count 99995)
99.997% <= 0.687 milliseconds (cumulative count 99999)
99.999% <= 0.695 milliseconds (cumulative count 100000)
100.000% <= 0.695 milliseconds (cumulative count 100000)

Cumulative distribution of latencies:
0.000% <= 0.103 milliseconds (cumulative count 0)
0.040% <= 0.207 milliseconds (cumulative count 40)
91.284% <= 0.303 milliseconds (cumulative count 91284)
99.131% <= 0.407 milliseconds (cumulative count 99131)
99.689% <= 0.503 milliseconds (cumulative count 99689)
99.950% <= 0.607 milliseconds (cumulative count 99950)
100.000% <= 0.703 milliseconds (cumulative count 100000)

Summary:
  throughput summary: 92506.94 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.282     0.128     0.279     0.319     0.399     0.695
====== SET ======
  100000 requests completed in 1.09 seconds
  50 parallel clients
  3 bytes payload
  keep alive: 1
  host configuration "save": 900 1 300 10 60 10000
  host configuration "appendonly": yes
  multi-thread: no

Latency by percentile distribution:
0.000% <= 0.159 milliseconds (cumulative count 1)
50.000% <= 0.327 milliseconds (cumulative count 50742)
75.000% <= 0.375 milliseconds (cumulative count 77708)
87.500% <= 0.407 milliseconds (cumulative count 87592)
93.750% <= 0.455 milliseconds (cumulative count 94226)
96.875% <= 0.503 milliseconds (cumulative count 97026)
98.438% <= 0.559 milliseconds (cumulative count 98551)
99.219% <= 0.623 milliseconds (cumulative count 99286)
99.609% <= 0.679 milliseconds (cumulative count 99614)
99.805% <= 0.743 milliseconds (cumulative count 99809)
99.902% <= 0.839 milliseconds (cumulative count 99906)
99.951% <= 0.991 milliseconds (cumulative count 99952)
99.976% <= 1.639 milliseconds (cumulative count 99976)
99.988% <= 1.839 milliseconds (cumulative count 99988)
99.994% <= 1.959 milliseconds (cumulative count 99994)
99.997% <= 1.991 milliseconds (cumulative count 99997)
99.998% <= 2.039 milliseconds (cumulative count 99999)
99.999% <= 2.055 milliseconds (cumulative count 100000)
100.000% <= 2.055 milliseconds (cumulative count 100000)

Cumulative distribution of latencies:
0.000% <= 0.103 milliseconds (cumulative count 0)
0.027% <= 0.207 milliseconds (cumulative count 27)
33.771% <= 0.303 milliseconds (cumulative count 33771)
87.592% <= 0.407 milliseconds (cumulative count 87592)
97.026% <= 0.503 milliseconds (cumulative count 97026)
99.159% <= 0.607 milliseconds (cumulative count 99159)
99.712% <= 0.703 milliseconds (cumulative count 99712)
99.880% <= 0.807 milliseconds (cumulative count 99880)
99.937% <= 0.903 milliseconds (cumulative count 99937)
99.953% <= 1.007 milliseconds (cumulative count 99953)
99.961% <= 1.407 milliseconds (cumulative count 99961)
99.967% <= 1.503 milliseconds (cumulative count 99967)
99.973% <= 1.607 milliseconds (cumulative count 99973)
99.981% <= 1.703 milliseconds (cumulative count 99981)
99.985% <= 1.807 milliseconds (cumulative count 99985)
99.991% <= 1.903 milliseconds (cumulative count 99991)
99.998% <= 2.007 milliseconds (cumulative count 99998)
100.000% <= 2.103 milliseconds (cumulative count 100000)

Summary:
  throughput summary: 91911.76 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.339     0.152     0.327     0.471     0.599     2.055
====== GET ======
  100000 requests completed in 1.10 seconds
  50 parallel clients
  3 bytes payload
  keep alive: 1
  host configuration "save": 900 1 300 10 60 10000
  host configuration "appendonly": yes
  multi-thread: no

Latency by percentile distribution:
0.000% <= 0.143 milliseconds (cumulative count 1)
50.000% <= 0.287 milliseconds (cumulative count 65441)
75.000% <= 0.295 milliseconds (cumulative count 79143)
87.500% <= 0.311 milliseconds (cumulative count 89994)
93.750% <= 0.327 milliseconds (cumulative count 94471)
96.875% <= 0.351 milliseconds (cumulative count 97152)
98.438% <= 0.399 milliseconds (cumulative count 98602)
99.219% <= 0.455 milliseconds (cumulative count 99263)
99.609% <= 0.511 milliseconds (cumulative count 99632)
99.805% <= 0.591 milliseconds (cumulative count 99806)
99.902% <= 0.655 milliseconds (cumulative count 99904)
99.951% <= 0.735 milliseconds (cumulative count 99953)
99.976% <= 0.855 milliseconds (cumulative count 99976)
99.988% <= 1.199 milliseconds (cumulative count 99988)
99.994% <= 1.263 milliseconds (cumulative count 99994)
99.997% <= 1.303 milliseconds (cumulative count 99997)
99.998% <= 1.335 milliseconds (cumulative count 99999)
99.999% <= 1.351 milliseconds (cumulative count 100000)
100.000% <= 1.351 milliseconds (cumulative count 100000)

Cumulative distribution of latencies:
0.000% <= 0.103 milliseconds (cumulative count 0)
0.023% <= 0.207 milliseconds (cumulative count 23)
86.028% <= 0.303 milliseconds (cumulative count 86028)
98.714% <= 0.407 milliseconds (cumulative count 98714)
99.602% <= 0.503 milliseconds (cumulative count 99602)
99.826% <= 0.607 milliseconds (cumulative count 99826)
99.944% <= 0.703 milliseconds (cumulative count 99944)
99.974% <= 0.807 milliseconds (cumulative count 99974)
99.978% <= 0.903 milliseconds (cumulative count 99978)
99.980% <= 1.007 milliseconds (cumulative count 99980)
99.982% <= 1.103 milliseconds (cumulative count 99982)
99.989% <= 1.207 milliseconds (cumulative count 99989)
99.997% <= 1.303 milliseconds (cumulative count 99997)
100.000% <= 1.407 milliseconds (cumulative count 100000)

Summary:
  throughput summary: 91240.88 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.289     0.136     0.287     0.335     0.431     1.351
====== INCR ======
  100000 requests completed in 1.09 seconds
  50 parallel clients
  3 bytes payload
  keep alive: 1
  host configuration "save": 900 1 300 10 60 10000
  host configuration "appendonly": yes
  multi-thread: no

Latency by percentile distribution:
0.000% <= 0.143 milliseconds (cumulative count 2)
50.000% <= 0.319 milliseconds (cumulative count 51966)
75.000% <= 0.359 milliseconds (cumulative count 77929)
87.500% <= 0.391 milliseconds (cumulative count 88780)
93.750% <= 0.431 milliseconds (cumulative count 94167)
96.875% <= 0.487 milliseconds (cumulative count 97019)
98.438% <= 0.551 milliseconds (cumulative count 98492)
99.219% <= 0.615 milliseconds (cumulative count 99233)
99.609% <= 0.687 milliseconds (cumulative count 99617)
99.805% <= 0.783 milliseconds (cumulative count 99813)
99.902% <= 0.967 milliseconds (cumulative count 99905)
99.951% <= 6.399 milliseconds (cumulative count 99952)
99.976% <= 6.887 milliseconds (cumulative count 99976)
99.988% <= 7.263 milliseconds (cumulative count 99988)
99.994% <= 7.455 milliseconds (cumulative count 99994)
99.997% <= 7.511 milliseconds (cumulative count 99997)
99.998% <= 7.559 milliseconds (cumulative count 99999)
99.999% <= 7.583 milliseconds (cumulative count 100000)
100.000% <= 7.583 milliseconds (cumulative count 100000)

Cumulative distribution of latencies:
0.000% <= 0.103 milliseconds (cumulative count 0)
0.018% <= 0.207 milliseconds (cumulative count 18)
38.994% <= 0.303 milliseconds (cumulative count 38994)
91.656% <= 0.407 milliseconds (cumulative count 91656)
97.481% <= 0.503 milliseconds (cumulative count 97481)
99.163% <= 0.607 milliseconds (cumulative count 99163)
99.674% <= 0.703 milliseconds (cumulative count 99674)
99.835% <= 0.807 milliseconds (cumulative count 99835)
99.887% <= 0.903 milliseconds (cumulative count 99887)
99.913% <= 1.007 milliseconds (cumulative count 99913)
99.921% <= 1.103 milliseconds (cumulative count 99921)
99.934% <= 1.207 milliseconds (cumulative count 99934)
99.945% <= 1.303 milliseconds (cumulative count 99945)
99.949% <= 1.407 milliseconds (cumulative count 99949)
99.950% <= 1.503 milliseconds (cumulative count 99950)
99.981% <= 7.103 milliseconds (cumulative count 99981)
100.000% <= 8.103 milliseconds (cumulative count 100000)

Summary:
  throughput summary: 91491.30 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.332     0.136     0.319     0.447     0.591     7.583
====== LPUSH ======
  100000 requests completed in 1.09 seconds
  50 parallel clients
  3 bytes payload
  keep alive: 1
  host configuration "save": 900 1 300 10 60 10000
  host configuration "appendonly": yes
  multi-thread: no

Latency by percentile distribution:
0.000% <= 0.135 milliseconds (cumulative count 1)
50.000% <= 0.351 milliseconds (cumulative count 51611)
75.000% <= 0.407 milliseconds (cumulative count 76199)
87.500% <= 0.455 milliseconds (cumulative count 87806)
93.750% <= 0.511 milliseconds (cumulative count 94207)
96.875% <= 0.567 milliseconds (cumulative count 96906)
98.438% <= 0.639 milliseconds (cumulative count 98440)
99.219% <= 0.735 milliseconds (cumulative count 99257)
99.609% <= 0.831 milliseconds (cumulative count 99621)
99.805% <= 0.943 milliseconds (cumulative count 99813)
99.902% <= 1.055 milliseconds (cumulative count 99906)
99.951% <= 1.239 milliseconds (cumulative count 99954)
99.976% <= 1.455 milliseconds (cumulative count 99976)
99.988% <= 1.583 milliseconds (cumulative count 99988)
99.994% <= 1.647 milliseconds (cumulative count 99994)
99.997% <= 1.679 milliseconds (cumulative count 99997)
99.998% <= 1.703 milliseconds (cumulative count 99999)
99.999% <= 1.719 milliseconds (cumulative count 100000)
100.000% <= 1.719 milliseconds (cumulative count 100000)

Cumulative distribution of latencies:
0.000% <= 0.103 milliseconds (cumulative count 0)
0.061% <= 0.207 milliseconds (cumulative count 61)
25.820% <= 0.303 milliseconds (cumulative count 25820)
76.199% <= 0.407 milliseconds (cumulative count 76199)
93.592% <= 0.503 milliseconds (cumulative count 93592)
97.923% <= 0.607 milliseconds (cumulative count 97923)
99.060% <= 0.703 milliseconds (cumulative count 99060)
99.557% <= 0.807 milliseconds (cumulative count 99557)
99.756% <= 0.903 milliseconds (cumulative count 99756)
99.874% <= 1.007 milliseconds (cumulative count 99874)
99.926% <= 1.103 milliseconds (cumulative count 99926)
99.949% <= 1.207 milliseconds (cumulative count 99949)
99.963% <= 1.303 milliseconds (cumulative count 99963)
99.971% <= 1.407 milliseconds (cumulative count 99971)
99.980% <= 1.503 milliseconds (cumulative count 99980)
99.990% <= 1.607 milliseconds (cumulative count 99990)
99.999% <= 1.703 milliseconds (cumulative count 99999)
100.000% <= 1.807 milliseconds (cumulative count 100000)

Summary:
  throughput summary: 91659.03 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.365     0.128     0.351     0.527     0.695     1.719
====== RPUSH ======
  100000 requests completed in 1.07 seconds
  50 parallel clients
  3 bytes payload
  keep alive: 1
  host configuration "save": 900 1 300 10 60 10000
  host configuration "appendonly": yes
  multi-thread: no

Latency by percentile distribution:
0.000% <= 0.143 milliseconds (cumulative count 1)
50.000% <= 0.335 milliseconds (cumulative count 52057)
75.000% <= 0.383 milliseconds (cumulative count 76573)
87.500% <= 0.423 milliseconds (cumulative count 87757)
93.750% <= 0.471 milliseconds (cumulative count 93923)
96.875% <= 0.527 milliseconds (cumulative count 96912)
98.438% <= 0.599 milliseconds (cumulative count 98553)
99.219% <= 0.671 milliseconds (cumulative count 99277)
99.609% <= 0.743 milliseconds (cumulative count 99618)
99.805% <= 0.823 milliseconds (cumulative count 99817)
99.902% <= 0.919 milliseconds (cumulative count 99906)
99.951% <= 1.007 milliseconds (cumulative count 99954)
99.976% <= 1.111 milliseconds (cumulative count 99977)
99.988% <= 1.167 milliseconds (cumulative count 99989)
99.994% <= 1.223 milliseconds (cumulative count 99994)
99.997% <= 1.263 milliseconds (cumulative count 99997)
99.998% <= 1.335 milliseconds (cumulative count 100000)
100.000% <= 1.335 milliseconds (cumulative count 100000)

Cumulative distribution of latencies:
0.000% <= 0.103 milliseconds (cumulative count 0)
0.042% <= 0.207 milliseconds (cumulative count 42)
31.918% <= 0.303 milliseconds (cumulative count 31918)
84.209% <= 0.407 milliseconds (cumulative count 84209)
95.908% <= 0.503 milliseconds (cumulative count 95908)
98.663% <= 0.607 milliseconds (cumulative count 98663)
99.460% <= 0.703 milliseconds (cumulative count 99460)
99.791% <= 0.807 milliseconds (cumulative count 99791)
99.895% <= 0.903 milliseconds (cumulative count 99895)
99.954% <= 1.007 milliseconds (cumulative count 99954)
99.975% <= 1.103 milliseconds (cumulative count 99975)
99.993% <= 1.207 milliseconds (cumulative count 99993)
99.998% <= 1.303 milliseconds (cumulative count 99998)
100.000% <= 1.407 milliseconds (cumulative count 100000)

Summary:
  throughput summary: 93720.71 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.346     0.136     0.335     0.487     0.639     1.335
====== LPOP ======
  100000 requests completed in 1.08 seconds
  50 parallel clients
  3 bytes payload
  keep alive: 1
  host configuration "save": 900 1 300 10 60 10000
  host configuration "appendonly": yes
  multi-thread: no

Latency by percentile distribution:
0.000% <= 0.159 milliseconds (cumulative count 1)
50.000% <= 0.327 milliseconds (cumulative count 53787)
75.000% <= 0.367 milliseconds (cumulative count 77770)
87.500% <= 0.399 milliseconds (cumulative count 88544)
93.750% <= 0.439 milliseconds (cumulative count 94373)
96.875% <= 0.487 milliseconds (cumulative count 97129)
98.438% <= 0.535 milliseconds (cumulative count 98438)
99.219% <= 0.599 milliseconds (cumulative count 99269)
99.609% <= 0.663 milliseconds (cumulative count 99644)
99.805% <= 0.727 milliseconds (cumulative count 99815)
99.902% <= 0.799 milliseconds (cumulative count 99910)
99.951% <= 0.855 milliseconds (cumulative count 99954)
99.976% <= 0.903 milliseconds (cumulative count 99976)
99.988% <= 0.959 milliseconds (cumulative count 99988)
99.994% <= 0.999 milliseconds (cumulative count 99995)
99.997% <= 1.015 milliseconds (cumulative count 99997)
99.998% <= 1.063 milliseconds (cumulative count 99999)
99.999% <= 1.079 milliseconds (cumulative count 100000)
100.000% <= 1.079 milliseconds (cumulative count 100000)

Cumulative distribution of latencies:
0.000% <= 0.103 milliseconds (cumulative count 0)
0.012% <= 0.207 milliseconds (cumulative count 12)
35.667% <= 0.303 milliseconds (cumulative count 35667)
90.224% <= 0.407 milliseconds (cumulative count 90224)
97.669% <= 0.503 milliseconds (cumulative count 97669)
99.331% <= 0.607 milliseconds (cumulative count 99331)
99.757% <= 0.703 milliseconds (cumulative count 99757)
99.917% <= 0.807 milliseconds (cumulative count 99917)
99.976% <= 0.903 milliseconds (cumulative count 99976)
99.996% <= 1.007 milliseconds (cumulative count 99996)
100.000% <= 1.103 milliseconds (cumulative count 100000)

Summary:
  throughput summary: 92850.51 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.334     0.152     0.327     0.447     0.575     1.079
====== RPOP ======
  100000 requests completed in 1.09 seconds
  50 parallel clients
  3 bytes payload
  keep alive: 1
  host configuration "save": 900 1 300 10 60 10000
  host configuration "appendonly": yes
  multi-thread: no

Latency by percentile distribution:
0.000% <= 0.167 milliseconds (cumulative count 2)
50.000% <= 0.319 milliseconds (cumulative count 51399)
75.000% <= 0.359 milliseconds (cumulative count 77648)
87.500% <= 0.391 milliseconds (cumulative count 88919)
93.750% <= 0.423 milliseconds (cumulative count 93893)
96.875% <= 0.471 milliseconds (cumulative count 96919)
98.438% <= 0.527 milliseconds (cumulative count 98449)
99.219% <= 0.599 milliseconds (cumulative count 99247)
99.609% <= 0.671 milliseconds (cumulative count 99620)
99.805% <= 0.751 milliseconds (cumulative count 99810)
99.902% <= 0.863 milliseconds (cumulative count 99904)
99.951% <= 0.975 milliseconds (cumulative count 99954)
99.976% <= 2.399 milliseconds (cumulative count 99976)
99.988% <= 2.623 milliseconds (cumulative count 99989)
99.994% <= 2.807 milliseconds (cumulative count 99995)
99.997% <= 2.895 milliseconds (cumulative count 99997)
99.998% <= 2.927 milliseconds (cumulative count 99999)
99.999% <= 2.943 milliseconds (cumulative count 100000)
100.000% <= 2.943 milliseconds (cumulative count 100000)

Cumulative distribution of latencies:
0.000% <= 0.103 milliseconds (cumulative count 0)
0.016% <= 0.207 milliseconds (cumulative count 16)
38.444% <= 0.303 milliseconds (cumulative count 38444)
91.935% <= 0.407 milliseconds (cumulative count 91935)
97.941% <= 0.503 milliseconds (cumulative count 97941)
99.307% <= 0.607 milliseconds (cumulative count 99307)
99.725% <= 0.703 milliseconds (cumulative count 99725)
99.864% <= 0.807 milliseconds (cumulative count 99864)
99.927% <= 0.903 milliseconds (cumulative count 99927)
99.957% <= 1.007 milliseconds (cumulative count 99957)
99.966% <= 1.103 milliseconds (cumulative count 99966)
99.969% <= 1.207 milliseconds (cumulative count 99969)
99.973% <= 1.303 milliseconds (cumulative count 99973)
99.975% <= 1.407 milliseconds (cumulative count 99975)
100.000% <= 3.103 milliseconds (cumulative count 100000)

Summary:
  throughput summary: 92081.03 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.329     0.160     0.319     0.439     0.575     2.943
====== SADD ======
  100000 requests completed in 1.10 seconds
  50 parallel clients
  3 bytes payload
  keep alive: 1
  host configuration "save": 900 1 300 10 60 10000
  host configuration "appendonly": yes
  multi-thread: no

Latency by percentile distribution:
0.000% <= 0.143 milliseconds (cumulative count 2)
50.000% <= 0.287 milliseconds (cumulative count 65404)
75.000% <= 0.295 milliseconds (cumulative count 77868)
87.500% <= 0.311 milliseconds (cumulative count 88656)
93.750% <= 0.335 milliseconds (cumulative count 94410)
96.875% <= 0.375 milliseconds (cumulative count 97156)
98.438% <= 0.447 milliseconds (cumulative count 98441)
99.219% <= 0.559 milliseconds (cumulative count 99224)
99.609% <= 0.663 milliseconds (cumulative count 99624)
99.805% <= 0.807 milliseconds (cumulative count 99806)
99.902% <= 1.111 milliseconds (cumulative count 99903)
99.951% <= 1.367 milliseconds (cumulative count 99954)
99.976% <= 1.487 milliseconds (cumulative count 99976)
99.988% <= 1.583 milliseconds (cumulative count 99988)
99.994% <= 1.615 milliseconds (cumulative count 99994)
99.997% <= 1.671 milliseconds (cumulative count 99997)
99.998% <= 1.719 milliseconds (cumulative count 99999)
99.999% <= 1.751 milliseconds (cumulative count 100000)
100.000% <= 1.751 milliseconds (cumulative count 100000)

Cumulative distribution of latencies:
0.000% <= 0.103 milliseconds (cumulative count 0)
0.019% <= 0.207 milliseconds (cumulative count 19)
84.680% <= 0.303 milliseconds (cumulative count 84680)
97.941% <= 0.407 milliseconds (cumulative count 97941)
98.902% <= 0.503 milliseconds (cumulative count 98902)
99.453% <= 0.607 milliseconds (cumulative count 99453)
99.706% <= 0.703 milliseconds (cumulative count 99706)
99.806% <= 0.807 milliseconds (cumulative count 99806)
99.837% <= 0.903 milliseconds (cumulative count 99837)
99.864% <= 1.007 milliseconds (cumulative count 99864)
99.902% <= 1.103 milliseconds (cumulative count 99902)
99.921% <= 1.207 milliseconds (cumulative count 99921)
99.938% <= 1.303 milliseconds (cumulative count 99938)
99.960% <= 1.407 milliseconds (cumulative count 99960)
99.978% <= 1.503 milliseconds (cumulative count 99978)
99.993% <= 1.607 milliseconds (cumulative count 99993)
99.997% <= 1.703 milliseconds (cumulative count 99997)
100.000% <= 1.807 milliseconds (cumulative count 100000)

Summary:
  throughput summary: 90579.71 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.292     0.136     0.287     0.343     0.527     1.751
====== HSET ======
  100000 requests completed in 1.07 seconds
  50 parallel clients
  3 bytes payload
  keep alive: 1
  host configuration "save": 900 1 300 10 60 10000
  host configuration "appendonly": yes
  multi-thread: no

Latency by percentile distribution:
0.000% <= 0.143 milliseconds (cumulative count 1)
50.000% <= 0.359 milliseconds (cumulative count 53013)
75.000% <= 0.415 milliseconds (cumulative count 76647)
87.500% <= 0.463 milliseconds (cumulative count 88839)
93.750% <= 0.503 milliseconds (cumulative count 94058)
96.875% <= 0.551 milliseconds (cumulative count 96933)
98.438% <= 0.607 milliseconds (cumulative count 98513)
99.219% <= 0.663 milliseconds (cumulative count 99236)
99.609% <= 0.743 milliseconds (cumulative count 99625)
99.805% <= 0.871 milliseconds (cumulative count 99809)
99.902% <= 1.031 milliseconds (cumulative count 99906)
99.951% <= 1.151 milliseconds (cumulative count 99954)
99.976% <= 1.247 milliseconds (cumulative count 99977)
99.988% <= 1.311 milliseconds (cumulative count 99988)
99.994% <= 1.367 milliseconds (cumulative count 99994)
99.997% <= 1.399 milliseconds (cumulative count 99997)
99.998% <= 1.423 milliseconds (cumulative count 99999)
99.999% <= 1.439 milliseconds (cumulative count 100000)
100.000% <= 1.439 milliseconds (cumulative count 100000)

Cumulative distribution of latencies:
0.000% <= 0.103 milliseconds (cumulative count 0)
0.130% <= 0.207 milliseconds (cumulative count 130)
24.864% <= 0.303 milliseconds (cumulative count 24864)
73.853% <= 0.407 milliseconds (cumulative count 73853)
94.058% <= 0.503 milliseconds (cumulative count 94058)
98.513% <= 0.607 milliseconds (cumulative count 98513)
99.489% <= 0.703 milliseconds (cumulative count 99489)
99.739% <= 0.807 milliseconds (cumulative count 99739)
99.834% <= 0.903 milliseconds (cumulative count 99834)
99.894% <= 1.007 milliseconds (cumulative count 99894)
99.940% <= 1.103 milliseconds (cumulative count 99940)
99.966% <= 1.207 milliseconds (cumulative count 99966)
99.987% <= 1.303 milliseconds (cumulative count 99987)
99.998% <= 1.407 milliseconds (cumulative count 99998)
100.000% <= 1.503 milliseconds (cumulative count 100000)

Summary:
  throughput summary: 93196.65 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.366     0.136     0.359     0.519     0.639     1.439
====== SPOP ======
  100000 requests completed in 1.09 seconds
  50 parallel clients
  3 bytes payload
  keep alive: 1
  host configuration "save": 900 1 300 10 60 10000
  host configuration "appendonly": yes
  multi-thread: no

Latency by percentile distribution:
0.000% <= 0.159 milliseconds (cumulative count 2)
50.000% <= 0.279 milliseconds (cumulative count 53419)
75.000% <= 0.295 milliseconds (cumulative count 78634)
87.500% <= 0.319 milliseconds (cumulative count 89941)
93.750% <= 0.343 milliseconds (cumulative count 94261)
96.875% <= 0.383 milliseconds (cumulative count 96884)
98.438% <= 0.439 milliseconds (cumulative count 98450)
99.219% <= 0.511 milliseconds (cumulative count 99254)
99.609% <= 0.591 milliseconds (cumulative count 99619)
99.805% <= 0.671 milliseconds (cumulative count 99817)
99.902% <= 0.727 milliseconds (cumulative count 99908)
99.951% <= 0.807 milliseconds (cumulative count 99952)
99.976% <= 0.911 milliseconds (cumulative count 99976)
99.988% <= 0.983 milliseconds (cumulative count 99988)
99.994% <= 1.015 milliseconds (cumulative count 99994)
99.997% <= 1.047 milliseconds (cumulative count 99998)
99.998% <= 1.071 milliseconds (cumulative count 100000)
100.000% <= 1.071 milliseconds (cumulative count 100000)

Cumulative distribution of latencies:
0.000% <= 0.103 milliseconds (cumulative count 0)
0.022% <= 0.207 milliseconds (cumulative count 22)
83.902% <= 0.303 milliseconds (cumulative count 83902)
97.660% <= 0.407 milliseconds (cumulative count 97660)
99.194% <= 0.503 milliseconds (cumulative count 99194)
99.659% <= 0.607 milliseconds (cumulative count 99659)
99.876% <= 0.703 milliseconds (cumulative count 99876)
99.952% <= 0.807 milliseconds (cumulative count 99952)
99.974% <= 0.903 milliseconds (cumulative count 99974)
99.992% <= 1.007 milliseconds (cumulative count 99992)
100.000% <= 1.103 milliseconds (cumulative count 100000)

Summary:
  throughput summary: 91911.76 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.289     0.152     0.279     0.351     0.487     1.071
====== ZADD ======
  100000 requests completed in 1.08 seconds
  50 parallel clients
  3 bytes payload
  keep alive: 1
  host configuration "save": 900 1 300 10 60 10000
  host configuration "appendonly": yes
  multi-thread: no

Latency by percentile distribution:
0.000% <= 0.143 milliseconds (cumulative count 2)
50.000% <= 0.287 milliseconds (cumulative count 56240)
75.000% <= 0.303 milliseconds (cumulative count 77862)
87.500% <= 0.319 milliseconds (cumulative count 87791)
93.750% <= 0.343 milliseconds (cumulative count 94650)
96.875% <= 0.367 milliseconds (cumulative count 97464)
98.438% <= 0.391 milliseconds (cumulative count 98546)
99.219% <= 0.431 milliseconds (cumulative count 99250)
99.609% <= 0.479 milliseconds (cumulative count 99637)
99.805% <= 0.567 milliseconds (cumulative count 99808)
99.902% <= 0.671 milliseconds (cumulative count 99906)
99.951% <= 0.887 milliseconds (cumulative count 99952)
99.976% <= 1.247 milliseconds (cumulative count 99976)
99.988% <= 1.407 milliseconds (cumulative count 99988)
99.994% <= 1.471 milliseconds (cumulative count 99994)
99.997% <= 1.511 milliseconds (cumulative count 99997)
99.998% <= 1.535 milliseconds (cumulative count 99999)
99.999% <= 1.543 milliseconds (cumulative count 100000)
100.000% <= 1.543 milliseconds (cumulative count 100000)

Cumulative distribution of latencies:
0.000% <= 0.103 milliseconds (cumulative count 0)
0.023% <= 0.207 milliseconds (cumulative count 23)
77.862% <= 0.303 milliseconds (cumulative count 77862)
98.914% <= 0.407 milliseconds (cumulative count 98914)
99.699% <= 0.503 milliseconds (cumulative count 99699)
99.851% <= 0.607 milliseconds (cumulative count 99851)
99.915% <= 0.703 milliseconds (cumulative count 99915)
99.950% <= 0.807 milliseconds (cumulative count 99950)
99.952% <= 0.903 milliseconds (cumulative count 99952)
99.958% <= 1.007 milliseconds (cumulative count 99958)
99.965% <= 1.103 milliseconds (cumulative count 99965)
99.973% <= 1.207 milliseconds (cumulative count 99973)
99.980% <= 1.303 milliseconds (cumulative count 99980)
99.988% <= 1.407 milliseconds (cumulative count 99988)
99.996% <= 1.503 milliseconds (cumulative count 99996)
100.000% <= 1.607 milliseconds (cumulative count 100000)

Summary:
  throughput summary: 92850.51 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.292     0.136     0.287     0.351     0.415     1.543
====== ZPOPMIN ======
  100000 requests completed in 1.07 seconds
  50 parallel clients
  3 bytes payload
  keep alive: 1
  host configuration "save": 900 1 300 10 60 10000
  host configuration "appendonly": yes
  multi-thread: no

Latency by percentile distribution:
0.000% <= 0.151 milliseconds (cumulative count 2)
50.000% <= 0.279 milliseconds (cumulative count 56989)
75.000% <= 0.287 milliseconds (cumulative count 75361)
87.500% <= 0.303 milliseconds (cumulative count 89681)
93.750% <= 0.319 milliseconds (cumulative count 94720)
96.875% <= 0.335 milliseconds (cumulative count 96915)
98.438% <= 0.359 milliseconds (cumulative count 98460)
99.219% <= 0.399 milliseconds (cumulative count 99266)
99.609% <= 0.447 milliseconds (cumulative count 99616)
99.805% <= 0.519 milliseconds (cumulative count 99808)
99.902% <= 0.591 milliseconds (cumulative count 99905)
99.951% <= 0.631 milliseconds (cumulative count 99952)
99.976% <= 0.671 milliseconds (cumulative count 99977)
99.988% <= 0.703 milliseconds (cumulative count 99988)
99.994% <= 0.727 milliseconds (cumulative count 99994)
99.997% <= 0.735 milliseconds (cumulative count 99998)
99.998% <= 0.743 milliseconds (cumulative count 99999)
99.999% <= 0.751 milliseconds (cumulative count 100000)
100.000% <= 0.751 milliseconds (cumulative count 100000)

Cumulative distribution of latencies:
0.000% <= 0.103 milliseconds (cumulative count 0)
0.019% <= 0.207 milliseconds (cumulative count 19)
89.681% <= 0.303 milliseconds (cumulative count 89681)
99.351% <= 0.407 milliseconds (cumulative count 99351)
99.779% <= 0.503 milliseconds (cumulative count 99779)
99.923% <= 0.607 milliseconds (cumulative count 99923)
99.988% <= 0.703 milliseconds (cumulative count 99988)
100.000% <= 0.807 milliseconds (cumulative count 100000)

Summary:
  throughput summary: 93808.63 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.282     0.144     0.279     0.327     0.383     0.751
====== LPUSH (needed to benchmark LRANGE) ======
  100000 requests completed in 1.09 seconds
  50 parallel clients
  3 bytes payload
  keep alive: 1
  host configuration "save": 900 1 300 10 60 10000
  host configuration "appendonly": yes
  multi-thread: no

Latency by percentile distribution:
0.000% <= 0.151 milliseconds (cumulative count 1)
50.000% <= 0.351 milliseconds (cumulative count 53973)
75.000% <= 0.399 milliseconds (cumulative count 77235)
87.500% <= 0.439 milliseconds (cumulative count 88414)
93.750% <= 0.479 milliseconds (cumulative count 93980)
96.875% <= 0.527 milliseconds (cumulative count 97042)
98.438% <= 0.583 milliseconds (cumulative count 98484)
99.219% <= 0.639 milliseconds (cumulative count 99225)
99.609% <= 0.711 milliseconds (cumulative count 99617)
99.805% <= 0.815 milliseconds (cumulative count 99808)
99.902% <= 0.927 milliseconds (cumulative count 99905)
99.951% <= 9.823 milliseconds (cumulative count 99952)
99.976% <= 10.279 milliseconds (cumulative count 99976)
99.988% <= 10.527 milliseconds (cumulative count 99988)
99.994% <= 10.631 milliseconds (cumulative count 99994)
99.997% <= 10.687 milliseconds (cumulative count 99997)
99.998% <= 10.719 milliseconds (cumulative count 99999)
99.999% <= 10.751 milliseconds (cumulative count 100000)
100.000% <= 10.751 milliseconds (cumulative count 100000)

Cumulative distribution of latencies:
0.000% <= 0.103 milliseconds (cumulative count 0)
0.060% <= 0.207 milliseconds (cumulative count 60)
25.942% <= 0.303 milliseconds (cumulative count 25942)
80.046% <= 0.407 milliseconds (cumulative count 80046)
95.844% <= 0.503 milliseconds (cumulative count 95844)
98.853% <= 0.607 milliseconds (cumulative count 98853)
99.593% <= 0.703 milliseconds (cumulative count 99593)
99.799% <= 0.807 milliseconds (cumulative count 99799)
99.891% <= 0.903 milliseconds (cumulative count 99891)
99.934% <= 1.007 milliseconds (cumulative count 99934)
99.944% <= 1.103 milliseconds (cumulative count 99944)
99.946% <= 1.207 milliseconds (cumulative count 99946)
99.949% <= 1.303 milliseconds (cumulative count 99949)
99.950% <= 1.407 milliseconds (cumulative count 99950)
99.966% <= 10.103 milliseconds (cumulative count 99966)
100.000% <= 11.103 milliseconds (cumulative count 100000)

Summary:
  throughput summary: 92081.03 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.361     0.144     0.351     0.495     0.623    10.751
====== LRANGE_100 (first 100 elements) ======
  100000 requests completed in 3.38 seconds
  50 parallel clients
  3 bytes payload
  keep alive: 1
  host configuration "save": 900 1 300 10 60 10000
  host configuration "appendonly": yes
  multi-thread: no

Latency by percentile distribution:
0.000% <= 0.231 milliseconds (cumulative count 1)
50.000% <= 0.831 milliseconds (cumulative count 52891)
75.000% <= 0.895 milliseconds (cumulative count 75118)
87.500% <= 0.959 milliseconds (cumulative count 88224)
93.750% <= 1.015 milliseconds (cumulative count 94022)
96.875% <= 1.079 milliseconds (cumulative count 97050)
98.438% <= 1.151 milliseconds (cumulative count 98535)
99.219% <= 1.239 milliseconds (cumulative count 99270)
99.609% <= 1.319 milliseconds (cumulative count 99627)
99.805% <= 1.399 milliseconds (cumulative count 99812)
99.902% <= 1.479 milliseconds (cumulative count 99903)
99.951% <= 1.599 milliseconds (cumulative count 99952)
99.976% <= 1.663 milliseconds (cumulative count 99976)
99.988% <= 1.711 milliseconds (cumulative count 99988)
99.994% <= 1.743 milliseconds (cumulative count 99994)
99.997% <= 1.807 milliseconds (cumulative count 99998)
99.998% <= 1.823 milliseconds (cumulative count 99999)
99.999% <= 1.839 milliseconds (cumulative count 100000)
100.000% <= 1.839 milliseconds (cumulative count 100000)

Cumulative distribution of latencies:
0.000% <= 0.103 milliseconds (cumulative count 0)
0.004% <= 0.303 milliseconds (cumulative count 4)
0.010% <= 0.407 milliseconds (cumulative count 10)
0.018% <= 0.503 milliseconds (cumulative count 18)
0.153% <= 0.607 milliseconds (cumulative count 153)
5.229% <= 0.703 milliseconds (cumulative count 5229)
42.858% <= 0.807 milliseconds (cumulative count 42858)
77.410% <= 0.903 milliseconds (cumulative count 77410)
93.403% <= 1.007 milliseconds (cumulative count 93403)
97.695% <= 1.103 milliseconds (cumulative count 97695)
99.038% <= 1.207 milliseconds (cumulative count 99038)
99.566% <= 1.303 milliseconds (cumulative count 99566)
99.827% <= 1.407 milliseconds (cumulative count 99827)
99.913% <= 1.503 milliseconds (cumulative count 99913)
99.955% <= 1.607 milliseconds (cumulative count 99955)
99.984% <= 1.703 milliseconds (cumulative count 99984)
99.998% <= 1.807 milliseconds (cumulative count 99998)
100.000% <= 1.903 milliseconds (cumulative count 100000)

Summary:
  throughput summary: 29550.83 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.841     0.224     0.831     1.031     1.207     1.839
====== LRANGE_300 (first 300 elements) ======
  100000 requests completed in 7.97 seconds
  50 parallel clients
  3 bytes payload
  keep alive: 1
  host configuration "save": 900 1 300 10 60 10000
  host configuration "appendonly": yes
  multi-thread: no

Latency by percentile distribution:
0.000% <= 0.295 milliseconds (cumulative count 1)
50.000% <= 1.935 milliseconds (cumulative count 50553)
75.000% <= 2.063 milliseconds (cumulative count 75433)
87.500% <= 2.183 milliseconds (cumulative count 88107)
93.750% <= 2.295 milliseconds (cumulative count 94008)
96.875% <= 2.423 milliseconds (cumulative count 96987)
98.438% <= 2.599 milliseconds (cumulative count 98449)
99.219% <= 2.775 milliseconds (cumulative count 99233)
99.609% <= 2.951 milliseconds (cumulative count 99615)
99.805% <= 3.239 milliseconds (cumulative count 99806)
99.902% <= 3.599 milliseconds (cumulative count 99903)
99.951% <= 3.815 milliseconds (cumulative count 99953)
99.976% <= 3.943 milliseconds (cumulative count 99977)
99.988% <= 4.031 milliseconds (cumulative count 99988)
99.994% <= 4.151 milliseconds (cumulative count 99994)
99.997% <= 4.247 milliseconds (cumulative count 99997)
99.998% <= 4.279 milliseconds (cumulative count 99999)
99.999% <= 4.319 milliseconds (cumulative count 100000)
100.000% <= 4.319 milliseconds (cumulative count 100000)

Cumulative distribution of latencies:
0.000% <= 0.103 milliseconds (cumulative count 0)
0.001% <= 0.303 milliseconds (cumulative count 1)
0.003% <= 0.407 milliseconds (cumulative count 3)
0.004% <= 0.503 milliseconds (cumulative count 4)
0.007% <= 0.607 milliseconds (cumulative count 7)
0.010% <= 0.703 milliseconds (cumulative count 10)
0.013% <= 0.807 milliseconds (cumulative count 13)
0.020% <= 0.903 milliseconds (cumulative count 20)
0.032% <= 1.007 milliseconds (cumulative count 32)
0.045% <= 1.103 milliseconds (cumulative count 45)
0.066% <= 1.207 milliseconds (cumulative count 66)
0.109% <= 1.303 milliseconds (cumulative count 109)
0.159% <= 1.407 milliseconds (cumulative count 159)
0.262% <= 1.503 milliseconds (cumulative count 262)
1.223% <= 1.607 milliseconds (cumulative count 1223)
7.069% <= 1.703 milliseconds (cumulative count 7069)
22.178% <= 1.807 milliseconds (cumulative count 22178)
43.490% <= 1.903 milliseconds (cumulative count 43490)
66.038% <= 2.007 milliseconds (cumulative count 66038)
80.502% <= 2.103 milliseconds (cumulative count 80502)
99.735% <= 3.103 milliseconds (cumulative count 99735)
99.992% <= 4.103 milliseconds (cumulative count 99992)
100.000% <= 5.103 milliseconds (cumulative count 100000)

Summary:
  throughput summary: 12548.63 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        1.962     0.288     1.935     2.335     2.719     4.319
====== LRANGE_500 (first 500 elements) ======
  100000 requests completed in 12.36 seconds
  50 parallel clients
  3 bytes payload
  keep alive: 1
  host configuration "save": 900 1 300 10 60 10000
  host configuration "appendonly": yes
  multi-thread: no

Latency by percentile distribution:
0.000% <= 0.255 milliseconds (cumulative count 1)
50.000% <= 3.015 milliseconds (cumulative count 50411)
75.000% <= 3.239 milliseconds (cumulative count 75230)
87.500% <= 3.415 milliseconds (cumulative count 87859)
93.750% <= 3.559 milliseconds (cumulative count 93910)
96.875% <= 3.703 milliseconds (cumulative count 96997)
98.438% <= 3.847 milliseconds (cumulative count 98498)
99.219% <= 3.999 milliseconds (cumulative count 99222)
99.609% <= 4.143 milliseconds (cumulative count 99614)
99.805% <= 4.271 milliseconds (cumulative count 99808)
99.902% <= 4.687 milliseconds (cumulative count 99903)
99.951% <= 5.079 milliseconds (cumulative count 99952)
99.976% <= 5.303 milliseconds (cumulative count 99976)
99.988% <= 5.551 milliseconds (cumulative count 99988)
99.994% <= 5.719 milliseconds (cumulative count 99994)
99.997% <= 5.815 milliseconds (cumulative count 99997)
99.998% <= 5.983 milliseconds (cumulative count 99999)
99.999% <= 6.119 milliseconds (cumulative count 100000)
100.000% <= 6.119 milliseconds (cumulative count 100000)

Cumulative distribution of latencies:
0.000% <= 0.103 milliseconds (cumulative count 0)
0.001% <= 0.303 milliseconds (cumulative count 1)
0.002% <= 0.407 milliseconds (cumulative count 2)
0.003% <= 0.607 milliseconds (cumulative count 3)
0.004% <= 0.703 milliseconds (cumulative count 4)
0.005% <= 0.807 milliseconds (cumulative count 5)
0.007% <= 0.903 milliseconds (cumulative count 7)
0.010% <= 1.007 milliseconds (cumulative count 10)
0.011% <= 1.103 milliseconds (cumulative count 11)
0.015% <= 1.207 milliseconds (cumulative count 15)
0.019% <= 1.303 milliseconds (cumulative count 19)
0.024% <= 1.407 milliseconds (cumulative count 24)
0.034% <= 1.503 milliseconds (cumulative count 34)
0.047% <= 1.607 milliseconds (cumulative count 47)
0.065% <= 1.703 milliseconds (cumulative count 65)
0.079% <= 1.807 milliseconds (cumulative count 79)
0.096% <= 1.903 milliseconds (cumulative count 96)
0.117% <= 2.007 milliseconds (cumulative count 117)
0.177% <= 2.103 milliseconds (cumulative count 177)
61.082% <= 3.103 milliseconds (cumulative count 61082)
99.531% <= 4.103 milliseconds (cumulative count 99531)
99.957% <= 5.103 milliseconds (cumulative count 99957)
99.999% <= 6.103 milliseconds (cumulative count 99999)
100.000% <= 7.103 milliseconds (cumulative count 100000)

Summary:
  throughput summary: 8093.23 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        3.038     0.248     3.015     3.607     3.935     6.119
====== LRANGE_600 (first 600 elements) ======
  100000 requests completed in 14.52 seconds
  50 parallel clients
  3 bytes payload
  keep alive: 1
  host configuration "save": 900 1 300 10 60 10000
  host configuration "appendonly": yes
  multi-thread: no

Latency by percentile distribution:
0.000% <= 0.343 milliseconds (cumulative count 1)
50.000% <= 3.543 milliseconds (cumulative count 50367)
75.000% <= 3.783 milliseconds (cumulative count 75043)
87.500% <= 3.967 milliseconds (cumulative count 87636)
93.750% <= 4.127 milliseconds (cumulative count 93949)
96.875% <= 4.255 milliseconds (cumulative count 96924)
98.438% <= 4.383 milliseconds (cumulative count 98478)
99.219% <= 4.527 milliseconds (cumulative count 99220)
99.609% <= 4.671 milliseconds (cumulative count 99618)
99.805% <= 4.863 milliseconds (cumulative count 99807)
99.902% <= 4.999 milliseconds (cumulative count 99903)
99.951% <= 5.191 milliseconds (cumulative count 99952)
99.976% <= 5.679 milliseconds (cumulative count 99976)
99.988% <= 6.735 milliseconds (cumulative count 99988)
99.994% <= 7.647 milliseconds (cumulative count 99994)
99.997% <= 8.119 milliseconds (cumulative count 99997)
99.998% <= 8.431 milliseconds (cumulative count 99999)
99.999% <= 8.591 milliseconds (cumulative count 100000)
100.000% <= 8.591 milliseconds (cumulative count 100000)

Cumulative distribution of latencies:
0.000% <= 0.103 milliseconds (cumulative count 0)
0.001% <= 0.407 milliseconds (cumulative count 1)
0.002% <= 0.607 milliseconds (cumulative count 2)
0.003% <= 0.703 milliseconds (cumulative count 3)
0.004% <= 1.103 milliseconds (cumulative count 4)
0.005% <= 1.303 milliseconds (cumulative count 5)
0.006% <= 1.503 milliseconds (cumulative count 6)
0.007% <= 1.607 milliseconds (cumulative count 7)
0.009% <= 1.703 milliseconds (cumulative count 9)
0.010% <= 1.807 milliseconds (cumulative count 10)
0.011% <= 1.903 milliseconds (cumulative count 11)
0.012% <= 2.007 milliseconds (cumulative count 12)
0.015% <= 2.103 milliseconds (cumulative count 15)
6.974% <= 3.103 milliseconds (cumulative count 6974)
93.230% <= 4.103 milliseconds (cumulative count 93230)
99.938% <= 5.103 milliseconds (cumulative count 99938)
99.980% <= 6.103 milliseconds (cumulative count 99980)
99.990% <= 7.103 milliseconds (cumulative count 99990)
99.996% <= 8.103 milliseconds (cumulative count 99996)
100.000% <= 9.103 milliseconds (cumulative count 100000)

Summary:
  throughput summary: 6884.68 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        3.568     0.336     3.543     4.167     4.487     8.591
====== MSET (10 keys) ======
  100000 requests completed in 1.70 seconds
  50 parallel clients
  3 bytes payload
  keep alive: 1
  host configuration "save": 900 1 300 10 60 10000
  host configuration "appendonly": yes
  multi-thread: no

Latency by percentile distribution:
0.000% <= 0.247 milliseconds (cumulative count 1)
50.000% <= 0.727 milliseconds (cumulative count 51622)
75.000% <= 0.831 milliseconds (cumulative count 76254)
87.500% <= 0.903 milliseconds (cumulative count 88015)
93.750% <= 0.959 milliseconds (cumulative count 94086)
96.875% <= 1.015 milliseconds (cumulative count 97144)
98.438% <= 1.071 milliseconds (cumulative count 98485)
99.219% <= 1.191 milliseconds (cumulative count 99230)
99.609% <= 1.415 milliseconds (cumulative count 99613)
99.805% <= 3.551 milliseconds (cumulative count 99805)
99.902% <= 4.839 milliseconds (cumulative count 99903)
99.951% <= 26.463 milliseconds (cumulative count 99952)
99.976% <= 26.943 milliseconds (cumulative count 99976)
99.988% <= 27.663 milliseconds (cumulative count 99988)
99.994% <= 27.807 milliseconds (cumulative count 99995)
99.997% <= 27.871 milliseconds (cumulative count 99997)
99.998% <= 27.967 milliseconds (cumulative count 99999)
99.999% <= 27.999 milliseconds (cumulative count 100000)
100.000% <= 27.999 milliseconds (cumulative count 100000)

Cumulative distribution of latencies:
0.000% <= 0.103 milliseconds (cumulative count 0)
0.009% <= 0.303 milliseconds (cumulative count 9)
0.877% <= 0.407 milliseconds (cumulative count 877)
1.306% <= 0.503 milliseconds (cumulative count 1306)
17.989% <= 0.607 milliseconds (cumulative count 17989)
44.995% <= 0.703 milliseconds (cumulative count 44995)
71.440% <= 0.807 milliseconds (cumulative count 71440)
88.015% <= 0.903 milliseconds (cumulative count 88015)
96.825% <= 1.007 milliseconds (cumulative count 96825)
98.841% <= 1.103 milliseconds (cumulative count 98841)
99.280% <= 1.207 milliseconds (cumulative count 99280)
99.479% <= 1.303 milliseconds (cumulative count 99479)
99.607% <= 1.407 milliseconds (cumulative count 99607)
99.671% <= 1.503 milliseconds (cumulative count 99671)
99.708% <= 1.607 milliseconds (cumulative count 99708)
99.718% <= 1.703 milliseconds (cumulative count 99718)
99.724% <= 1.807 milliseconds (cumulative count 99724)
99.726% <= 1.903 milliseconds (cumulative count 99726)
99.731% <= 2.007 milliseconds (cumulative count 99731)
99.736% <= 2.103 milliseconds (cumulative count 99736)
99.797% <= 3.103 milliseconds (cumulative count 99797)
99.850% <= 4.103 milliseconds (cumulative count 99850)
99.924% <= 5.103 milliseconds (cumulative count 99924)
99.950% <= 6.103 milliseconds (cumulative count 99950)
99.976% <= 27.103 milliseconds (cumulative count 99976)
100.000% <= 28.111 milliseconds (cumulative count 100000)

Summary:
  throughput summary: 58823.53 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.756     0.240     0.727     0.975     1.135    27.999

redis-benchmark -h localhost -a aKr84vEm5Ekd  38.91s user 16.67s system 98% cpu 56.356 total
```

## Redis 的高并发与高可用

+ 高并发: Redis 主从架构的集群，通过配置读写分离，可以支持上百万的读QPS。
+ 高可用：Redis 通过配置 Sentinel(哨兵)，可以在 Master 挂掉之后及时切换主备，让集群的故障时间尽可能短。

## 解决异步复制和脑裂导致的数据丢失

```
+ min-slaves-to-write 1
+ min-slaves-max-lag 10
```

要求至少有1个 slave，数据复制和同步的延迟不能超过10秒

+ 一旦所有的slave，数据复制和同步的延迟都超过了10秒钟，那么这个时候，master就不会再接收写请求了

`min-slaves-max-lag`配置后，master 一旦认为 slave 的数据和自己的差别太大的话(salve ack 的时间超过了10秒钟)，就停止接收写请求。在脑裂情景下，空集群没有 slave 向 master 请求数据，master 也会停止写请求。这样集群恢复后，小集群的 master 也不会有数据，造成数据丢失。

`min-slaves-to-write`配置后，一旦发生了脑裂，小集群的 master 认为 slave 节点数不够，也会停止接收写请求。
