## 智能合约大小大约可以有多大？
```
200kb
解决方法
1. 使用升级合约，减少每次升级的部署量
2. 模块化
3. 编辑器优化 optimizer
减少重复指令 → 减少字节码大小
合并常量和计算 → 减少部署 gas
```

## ETH
```
1 ETH = 10的18次方 gwei
```

## fallback 和 receive 的区别？
```
receive()：接收 ETH ；
fallback()：处理未匹配函数或 ETH 调用。
```

## call方法
```
call 是 Solidity 提供的 底层函数调用方法，用于在运行时动态调用其他合约的函数。
_target.call(abi.encodeWithSelector(Target.setData.selector, _num, _text));


```




