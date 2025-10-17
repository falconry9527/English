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