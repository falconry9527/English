## CI/CD
```
CI/CD 自动化部署流程。

Jenkins: 
1. Jenkins 检测git提交 →  执行部署脚本
2. Jenkins 构建完成镜像 → 推送镜像仓库 → Jenkins 脚本触发 Kubernetes 重新部署。

------构建完成镜像 :
1. go build -o app .
2. docker build -t registry.cn-hangzhou.aliyuncs.com/demo/go-demo:1.0.0 .
3. echo $ALIYUN_PASS | docker login -u $ALIYUN_USER --password-stdin registry.cn-hangzhou.aliyuncs.com
------推送镜像仓库 :
4. docker push registry.cn-hangzhou.aliyuncs.com/demo/go-demo:1.0.0
------k8s部署 :
5. kubectl set image deployment/go-demo go-demo=registry.cn-hangzhou.aliyuncs.com/demo/go-demo:1.0.0 -n prod
6. kubectl rollout status deployment/go-demo -n prod

```

## 曲线图高并发
```
1. 接口层限流
var limiter = rate.NewLimiter(100, 200) // 每秒最多100请求，桶容量200

2. 热点数据缓存 : 最近 N 条价格，或最近 24 小时数据。

3. 历史 预计算 : 历史数据预计算

4. 数据库优化 : 读写分离，冷热分离，分库分表，加索引

```

## 微服务
```
微服务是一种 架构风格，把一个大型应用拆分成多个 小型、独立、可部署的服务。

每个微服务:
聚焦 单一业务能力（Single Responsibility）。
拥有独立的数据存储和业务逻辑。

```