# week 8
1.使用 redis benchmark 工具, 测试 10 20 50 100 200 1k 5k 字节 value 大小，redis get set 性能。  
2.写入一定量的 kv 数据, 根据数据大小 1w-50w 自己评估, 结合写入前后的 info memory 信息  , 分析上述不同 value 大小下，平均每个 key 的占用内存空间。
## 性能测试
### 30w 10B
SET: 107565.44 requests per second, p50=0.247 msec  
GET: 107411.38 requests per second, p50=0.247 msec
### 30w 20B
SET: 109090.91 requests per second, p50=0.247 msec  
GET: 109210.05 requests per second, p50=0.247 msec
### 30w 50B
SET: 106609.81 requests per second, p50=0.247 msec  
GET: 108499.09 requests per second, p50=0.247 msec
### 30w 100B
SET: 108381.51 requests per second, p50=0.247 msec  
GET: 108499.09 requests per second, p50=0.247 msec
### 30w 200B
SET: 105633.80 requests per second, p50=0.255 msec  
GET: 104821.80 requests per second, p50=0.255 msec
### 30w 1KB
SET: 106647.71 requests per second, p50=0.247 msec  
GET: 106007.07 requests per second, p50=0.247 msec
### 30w 5KB
SET: 102986.61 requests per second, p50=0.255 msec  
GET: 102319.24 requests per second, p50=0.263 msec
## key计算
used_memory:1769808  
used_memory_human:1.69M  
used_memory_rss:4947968  
used_memory_rss_human:4.72M  
used_memory_peak:4500512  
used_memory_peak_human:4.29M  
used_memory_peak_perc:39.32%  
used_memory_overhead:1026984  
used_memory_startup:1009360  
used_memory_dataset:742824

### 30w 10B
used_memory:1769808  
used_memory_human:1.69M  
used_memory_rss:4947968  
used_memory_rss_human:4.72M  
used_memory_peak:4500512  
used_memory_peak_human:4.29M  
used_memory_peak_perc:39.32%  
used_memory_overhead:1026984  
used_memory_startup:1009360  
used_memory_dataset:742824  
忽略不记
### 30w 20B
used_memory:1769824  
used_memory_human:1.69M  
used_memory_rss:5132288  
used_memory_rss_human:4.89M  
used_memory_peak:4500512  
used_memory_peak_human:4.29M  
used_memory_peak_perc:39.32%  
used_memory_overhead:1026984  
used_memory_startup:1009360  
used_memory_dataset:742840  
平均每个kv占用(742840-742824)/300000=0.00005B  
# Memory 50
used_memory:1769856  
used_memory_human:1.69M  
used_memory_rss:5218304  
used_memory_rss_human:4.98M  
used_memory_peak:4500512  
used_memory_peak_human:4.29M  
used_memory_peak_perc:39.33%  
used_memory_overhead:1026984  
used_memory_startup:1009360  
used_memory_dataset:742872  
平均每个kv占用(742872-742840)/300000=0.00010B  
# Memory 100
used_memory:1769904  
used_memory_human:1.69M  
used_memory_rss:5267456  
used_memory_rss_human:5.02M  
used_memory_peak:4500512  
used_memory_peak_human:4.29M  
used_memory_peak_perc:39.33%  
used_memory_overhead:1026984  
used_memory_startup:1009360  
used_memory_dataset:742920  
平均每个kv占用(742920-742872)/300000=0.00016B  
# Memory 200
used_memory:1770000  
used_memory_human:1.69M  
used_memory_rss:5332992  
used_memory_rss_human:5.09M  
used_memory_peak:4500512  
used_memory_peak_human:4.29M  
used_memory_peak_perc:39.33%  
used_memory_overhead:1026984  
used_memory_startup:1009360  
used_memory_dataset:743016  
平均每个kv占用(743016-742920)/300000=0.00032B
# Memory 1024
used_memory:1771328  
used_memory_human:1.69M  
used_memory_rss:5419008  
used_memory_rss_human:5.17M  
used_memory_peak:4500512  
used_memory_peak_human:4.29M  
used_memory_peak_perc:39.36%  
used_memory_overhead:1026984  
used_memory_startup:1009360  
used_memory_dataset:744344  
平均每个kv占用(744344-743016)/300000=0.00442B  
# Memory 5120
used_memory:1775424  
used_memory_human:1.69M  
used_memory_rss:5758976  
used_memory_rss_human:5.49M  
used_memory_peak:4500512  
used_memory_peak_human:4.29M  
used_memory_peak_perc:39.45%  
used_memory_overhead:1026984  
used_memory_startup:1009360  
used_memory_dataset:748440  
平均每个kv占用(748440-744344)/300000=0.01365B  