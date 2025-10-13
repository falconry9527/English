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

MySQL InnoDB 支持 多种锁机制，主要是为了保证事务的 隔离性 和 一致性。
脏读：读到未提交的数据
不可重复读：同一事务中两次读取同一行不同
幻读：同一事务中两次查询返回的行数不同

READ UNCOMMITTED （读未提交）：脏读，不可重复读，幻读
READ COMMITTED （读已提交）：不可重复读，幻读
REPEATABLE READ （可重复读）：幻读
SERIALIZABLE （串行化）： 任何问题都不会出现
```

## 死锁
```
死锁是指 两个或多个事务在数据库中互相等待对方持有的锁，从而无法继续执行。

事务 T1 ：
START TRANSACTION;
UPDATE account SET balance = balance - 10 WHERE id = 1;  -- 锁住 id=1
UPDATE account SET balance = balance + 10 WHERE id = 2;  -- 尝试锁住 id=2

事务 T2 ：
START TRANSACTION;
UPDATE account SET balance = balance - 20 WHERE id = 2;  -- 锁住 id=2
UPDATE account SET balance = balance + 20 WHERE id = 1;  -- 尝试锁住 id=1

相互等待 → 死锁

解决方法：
1. 事务顺序一致
2. 捕获死锁异常重试
3. 加索引 ： 避免全表扫描导致锁多行，提高锁粒度为行锁

```


## mysql优化 
```
1. 加索引
2. sql 优化
3. 分库分表
4. 读写分离
5. 冷热分离
```
