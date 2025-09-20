## 基础语法与概念
```
storage、memory、calldata 的区别？
storage：链上永久存储，memory：函数内部临时存储，calldata：函数外部调用传入只读参数。

public、external、internal、private 区别？
public：链内外可访问；external：链外调用；internal：合约及继承可访问；private：仅当前合约可访问。

view、pure、普通函数区别？
view 只读状态；pure 不读写状态；普通函数可修改状态。

fallback & receive 函数？
receive()：接收 ETH；fallback()：处理未匹配函数或 ETH 调用。

```


## 基础语法与概念
```
构造函数是什么？
在合约部署时执行一次，用于初始化状态变量。


delegatecall vs call
delegatecall 在调用者上下文执行，被调用合约修改调用者状态；call 修改被调用合约状态。

selfdestruct
销毁合约，返还剩余 ETH 给指定地址。

mapping 特点
无法遍历、默认值 0、常用于地址 → 数据。


动态 vs 固定数组
动态可增减长度，固定长度 gas 更低。

重入攻击 ： 攻击者在外部调用前不断重复调用合约，解决：使用ReentrancyGuard。
权限控制 ：使用 Ownable 或 AccessControl 模块限制操作。

```

## uups 和 transparent代理的区别
```

```

## 安全问题
```

```

## gas 优化
```

```


