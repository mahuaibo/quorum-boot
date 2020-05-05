# boot-cli
Quorum 网络引导服务

## 技术选型
```
Golang开发语言
CouchDB数据库
```

## 逻辑架构
![逻辑架构](./img.png)
```
主要三个组件：boot-service、boot-client、wallet
boot-service
收集联盟链启动的配置信息
boot-client
创建私钥信息、从boot-service下载节点信息，分发到服务器并且启动节点
wallet
存储私钥信息、只供boot-client调用
```

