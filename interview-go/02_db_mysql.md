## mysql 为什么使用 b+tree，不使用hash，二叉树，红黑树
```
hash tree : 只适合单点查询，不适合范围（><）查询
二叉树： 在极端情况下（1，2，3，4，5）会弱化成链，查询很慢
红黑树： 虽然用颜色重构树，但是，树的深度可能很深，查询会比较慢
b+tree ： 多叉树（每个父节点可以有 100个子节点），树的深度不会很深
```

## mysql 事务 和 锁
```
事务是一组 逻辑上必须作为一个整体执行的 SQL 操作
特点：原子性（Atomicity）、一致性（Consistency）、隔离性（Isolation）、持久性（Durability） → ACID

脏读：读到未提交的数据
不可重复读：同一事务中两次读取同一行不同 ，事务隔离性降低
幻读：同一事务中两次查询返回的行数不同 ，事务隔离性降低

事物的隔离性: 事物单独执行，读取期间，别的事物不能参与数据修改
READ UNCOMMITTED （读未提交）：脏读，不可重复读，幻读
READ COMMITTED （读已提交）：不可重复读，幻读
REPEATABLE READ （可重复读）：幻读
SERIALIZABLE （串行化）： 任何问题都不会出现

MySQL InnoDB 支持 多种锁机制，主要是为了保证事务的 隔离性 和 一致性。

```

## 死锁
```
死锁是指 两个或多个事务在数据库中互相等待对方持有的锁，从而导致查询阻塞，无法继续执行。

事务 T1 ：
START TRANSACTION;
UPDATE account SET balance = balance - 10 WHERE id = 1;  -- 锁住 id=1
UPDATE account SET balance = balance + 10 WHERE id = 2;  -- 尝试锁住 id=2

事务 T2 ：
START TRANSACTION;
UPDATE account SET balance = balance - 20 WHERE id = 2;  -- 锁住 id=2
UPDATE account SET balance = balance + 20 WHERE id = 1;  -- 尝试锁住 id=1

相互等待 → 死锁

解决方法：(转账为例)
1. 事务顺序一致
转账逻辑中，修改余额的顺序，先修改小id ，再修改大id 的数据
2. 使用golang 的锁，根据用户初始化锁
注意： 不能使用全局锁，否则并发很低
3.  加索引 ： 避免全表扫描导致锁多行，提高锁粒度为行锁

```


## mysql优化 
```
1. 加索引
2. sql 优化
3. 分库分表
4. 读写分离
5. 冷热分离
```


## in/exist （ɪɡˈzɪst）
```
in: 某个字段是否存在某些值
exist : 是否存在某个字段

in : 集合操作: 先内后外 : 拿着集合去便利每一条数据 ；适合大表 join 小表
exist :存在性检查， 先外后内 :拿着基础表的数据去逐条匹配  ； 适合大表 join 大表

一般在开发中都不用，用join 
```

## 行列互转
```
行转列 （转成大宽表）： case
列转行 ： union
```

## 聚集索引
```
聚集索引：物理存储按照索引排序（主键就是聚集索引）
非聚集索引：物理存储不按照索引排序；
```

## mysql 分页很慢怎么办
```
SELECT * FROM users ORDER BY name LIMIT 1000000, 10;
问题分析： LIMIT offset, size ,其中 OFFSET 越大，扫描行数越多 

1.利用 排序列的索引，比如id，user_id , 避免 OFFSET 扫描
2.改用 条件范围,使用 WHERE + 索引范围, 避免 OFFSET 扫描
SELECT * FROM users WHERE id BETWEEN 1000001 AND 1000010 ;

```

## 乐观锁 和 悲观锁
```
悲观锁 (Pessimistic Lock) ：
场景：数据冲突会频繁发生
思想：先加锁（查询的时候 for update,再访问/修改数据，阻止其他事务修改，保证数据一致性。
1. SELECT balance FROM account WHERE id=? FOR UPDATE
2. UPDATE account SET balance=? WHERE id=?

乐观锁 (Optimistic Lock) ：
场景：数据冲突不常发生
思想： 不加锁，用 version 或者 时间戳 来判断是否是当前版本。
1. SELECT balance, version FROM account WHERE id=?
2. UPDATE account SET balance=?, version=version+1 WHERE id=? AND version=?

```
