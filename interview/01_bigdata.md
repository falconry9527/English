# java

# Collection
```
list：有序（元素有插入顺序）、允许重复元素。
ArrayList：基于数组实现，查询快，增删慢，线程不安全。
LinkedList：基于链表实现，查询慢，增删快，线程不安全。
Vector：基于数组实现，线程安全，效率较低。

set: 无序（HashSet），不允许重复元素
HashSet：基于哈希表实现，元素无序且不可重复。
LinkedHashSet：在 HashSet 的基础上维护了插入顺序，也支持迭代。
TreeSet：基于红黑树实现，元素按自然顺序或自定义比较器排序。

map 接口用于存储键值对，即映射关系。
HashMap：基于哈希表实现，键值对无序，线程不安全。
LinkedHashMap：在 HashMap 的基础上维护了插入顺序，也支持迭代。
TreeMap：基于红黑树实现，键（Key）按自然顺序或自定义比较器排序。
```
# sql
## sql优化
```
1.分区分桶
2.先过滤/distnct 再jioin
3.减少不必要的数据插入： 微批架构
```

## in/exist （ɪɡˈzɪst）
```
in： 某个字段是否存在某些值
exist ： 是否存在某个字段

in ：集合操作：先内后外 ： 拿着集合去便利每一条数据 ；适合大表 join 小表
exist ：存在性检查， 先外后内 ：拿着基础表的数据去逐条匹配  ； 适合大表 join 大表

一般在开发中都不用，用join 

```
## mysql b+tree
```
```

## in/exist （ɪɡˈzɪst）
```
in： 某个字段是否存在某些值
exist ： 是否存在某个字段

in ：一般适合小表
exist ： 在基础表里面逐条数据
```

# flink
## 游戏行业常用指标
```
roi（Return on Investment）： 投资回报率
ltv（Lifetime Value）： 用户生命周期价值（充值，广告收益）
retain：留存
ARPU：平均每用户收益 (Average Revenue Per User)
ARPPU：平均每付费用户收益 (Average Revenue Per Paying User)
```

## flink 时间和水位线
```
Flink 提供了三种时间语义，用于确定窗口计算的时间基础：
处理时间（Process Time）：:以执行计算的机器的本地时间作为事件时间。
摄入时间（Ingestion Time）：:以Flink 接收数据流时的时间作为事件时间。
事件时间（Event Time）：:以事件实际发生的原始时间作为事件时间，通常从数据中提取，并在数据流中进行标记。

水位线是基于事件时间语义的关键机制。它解决了以下问题：
乱序数据处理：
在分布式系统中，数据在网络传输过程中可能会因为延迟而改变到达顺序，导致数据产生乱序。

```
