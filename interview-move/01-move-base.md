## move 语言的数据类型
```
基础类型 : bool,address,u8,u128
复合类型: string,struct,option,

集合类型：
可以遍历的：Vector(动态数组，内存连续，小容量)， VecMap (Key是动态数组，内存连续，小容量)
不能遍历的: Table(key-vaule 格式，无顺，查找快),linked_table (key-vaule 格式，有序，查找慢)

table : value 通常是基础类型，没有key的普通struct 实例 

dynamic_field：（父Object的） 子field 的  Key-vauleV索引系统（value 必须是必须有 store能力 ），子Object 的owner 必须是父Object
dynamic_object_field: （父Object的） 子Object 的 Key-vauleV 索引系统（value 的 strcut 必须有 key 和 store能力），子 Object 可独立拥有 。

object_table : 是 dynamic_object_field 的封装，但是 子Object 的 object owner 是父Object
add方法的时候，绑定了 子Object 的owner 是父Object 

```


##  Resource 
```
Resource：具有线性类型（没有 copy 和 drop 能力，不能被复制和隐式丢弃，必须显式转移或销毁）语义的 struct
Object = Resource + UID + 系统管理

资源安全模型：
Sui 的 Resource 安全模型通过 Move 的 线性资源类型，能力系统，对象所有权 ，
把 资产安全、并发安全  前移到语言和系统层，而不是依赖合约逻辑兜底。

```


##  能力系统
```
一. 概念
struct 具有的 4 种能力（Ability）：
copy：允许类型被复制；大多数原生数值类型默认具备该能力。
drop：允许类型在作用域结束时被隐式丢弃；没有 drop 的类型 需要显式处理（例如 object::delete）。
key：表示链上对象，有全局唯一 ID，可放在全局存储里作为对象；不能与 copy 或 drop 共存。
store： 表示它可以作为元素被存到持久化结构里，并且 可以跨模块操作。

transfer 和 public_transfer
transfer<T: key> : 模块内 （LP凭证）
public_transfer<T: key + store> : 跨模块 （coin）

二. 特殊情况
1. Hot Potato 结构体（Move）
没有任何能力的 struct，只能在同一个交易中被创建 和 强制销毁
例如闪电贷款：
- 借款时向 用户 返回 Hot Potato
- 用户 必须在同一个交易中 完成还款(消耗该结构体), 若未还款(消耗该结构体)，交易将直接失败

2. OTW
OTW 是一种特殊的见证类型，只能使用一次。它不能手动创建，并且保证每个模块中的 OTW 都是唯一的。
仅具备drop能力，没有字段，以模块名称命名，所有字母均为大写。

public struct MY_TOKEN has drop {}
coin::create_currency方法中要求使用 OTW ，从而确保 coin::TreasuryCap 只被创建一次

3. Resource ： 没有 copy 和 drop 能力的struct，就是 Resource。 

4 . OBJECT
具有 key 能力的 + UID 的struct实例,被称为Object。Object被sui系统管理。

```


##  OBJECT
```
一. 定义：
具有 key 能力 + UID，且 被sui系统管理的 struct实例 被称为Object。
sui 系统管理负责：
UID :  Object 的唯一身份标识(object::new(ctx))
Ownership ：
生命周期 : 
并发控制 和 Version 

二. 详解
1. OBJECT UID 
coin.split 的时候，更改原来coin代币的余额，生成新uid的 coin
coin.join 的时候，delete 合并coin，把余额合并到主coin

2. OBJECT 的 Owner 四种类型
Owned address: transfer::transfer(obj, recipient);
Shared Object（共享对象）: transfer::share_object(obj); 
Immutable Object（不可变对象） : transfer::freeze_object(obj);
Owned Object : dynamic_field / dynamic_object_field / object_table  的add 方法

3. 生命周期
创建UID（object::new(ctx)）-> 进入对象系统(transfer,share_object,freeze_object,dynamic_field add) -> 业务操作（以coin为例：split，cjoin ，transfer） -> delete 

4. 并发控制
Owned Object (Fast Path)：Owned Object 在任一时刻只有唯一 owner 拥有写权限，类似于单线程，没有并发修改 , 因此无需全局排序，不走共识机制 （sui的共识机制是 pos）
Shared Object (Consensus Path)：允许所有用户同时修改，类似于多现场，存在并发修改冲突，通过 shared version + 共识锁 对交易进行全局排序 ，走共识机制 

```


## 范型 和 反射  
```
1. 范型
Sui 的范型 是为了把 类型判断 前移到编译期，用类型系统保证资产安全。
phantom ：T不占存储，只用于类型区分

2. 反射
Sui Move 没有运行时反射，而是 用范型 把 类型判断  前移到编译期。 只有类型反射 （type_name）

类型反射 :
let defining_type_name: TypeName = type_name::with_defining_ids<T>();
let original_type_name: TypeName = type_name::with_original_ids<T>();

// Returns only "ID" of the package.
let defining_package: address = type_name::defining_id<T>();
let original_package: address = type_name::original_id<T>();

```

##  字段跨module访问
```
sui的字段都不能跨module 直接访问，需要提供对应的get 或者 set 方法
```

