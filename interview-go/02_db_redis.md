## redis 的数据结构
```
1. String（字符串）
rdb.Set(ctx, "key", "hello", 0)
val, _ := rdb.Get(ctx, "key").Result()

2. hash
rdb.HSet(ctx, "user:1", "name", "Alice", "age", 30)
name, _ := rdb.HGet(ctx, "user:1", "name").Result()
all, _ := rdb.HGetAll(ctx, "user:1").Result()

3. List（列表）： 有序列表，支持从两端 push/pop
rdb.LPush(ctx, "mylist", "task1")
rdb.RPush(ctx, "mylist", "task2")
items, _ := rdb.LRange(ctx, "mylist", 0, -1).Result()
rdb.LPop(ctx, "mylist")

4. Set（集合）
rdb.SAdd(ctx, "myset", "a", "b", "c")
rdb.SRem(ctx, "myset", "b")
members, _ := rdb.SMembers(ctx, "myset").Result()
rdb.SIsMember(ctx, "myset", "a")

5. Sorted Set（有序集合）
rdb.ZAdd(ctx, "leaderboard", &redis.Z{Score: 100, Member: "Alice"})
rdb.ZAdd(ctx, "leaderboard", &redis.Z{Score: 200, Member: "Bob"})
rank, _ := rdb.ZRangeWithScores(ctx, "leaderboard", 0, -1).Result()
rdb.ZRem(ctx, "leaderboard", "Alice")

```

## Redis 持久化原理
```
Redis 支持两种主要的持久化方式：RDB快照 和 AOF（Append Only File），也可以混合使用。
RDB快照 （默认）:
Redis 会在指定时间间隔生成数据快照，写入磁盘
过程： 父进程 fork() 出一个子进程,子进程将内存数据写入 .rdb 文件,父进程继续处理客户端请求，写操作不阻塞
redis.conf： 
save 900 1：在 900 秒（15 分钟）内至少有 1 次写操作，就触发一次 RDB 快照
save 300 10：在 300 秒（5 分钟）内至少有 10 次写操作，就触发一次 RDB 快照
save 60 10000：在 60 秒内至少有 10000 次写操作，就触发一次 RDB 快照

AOF :
Redis 将每条写命令以追加的方式记录到日志文件

```


##  Redis 淘汰策略
```
内存使用接近 maxmemory 配置时，如果还要写入新数据，就需要淘汰（删除）一些旧数据，这就是 内存淘汰策略

noeviction：不淘汰任何 key，当内存超限时写操作会报错。适合关键数据不可丢失，需要严格控制内存的场景。
LRU：只淘汰设置了过期时间的 key，淘汰 最近最少使用 的数据 。
LFU：只淘汰设置了过期时间的 key，淘汰 访问次数最少的数据。

```

##  Redis 穿透,雪崩,击穿
```
穿透：查询的数据本身 根本不存在，请求每次都会绕过缓存直接访问数据库。 → 用 布隆过滤器 / null 缓存
雪崩:大量缓存同时过期，在短时间内大量请求直接访问数据库，造成数据库压力骤增。
  缓存过期时间加随机值：
击穿：热点数据失效瞬间大量请求,大量请求直接访问数据库， 
  热点永不过期+ 异步更新：
```

##  Redis 的哨兵模式和集群模式区别
```
哨兵模式：
解决 主节点高可用 问题
数据不分片，写不能水平拓展（读可以）
适合 中小数据量缓存

集群模式：
解决 主节点高可用 + 水平扩展问题
数据分片，每个 Master 独立写, 写可以水平拓展
适合 大规模系统，高并发读写，容量大

数据分片： 
把数据分成多个部分（Shard），分别存储
每个 Master 节点负责一部分 slot → 数据分片

```

