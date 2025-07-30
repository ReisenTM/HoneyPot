# HoneyPot


# 项目搭建

**为什么要用两个消息队列**

Kafka:高吞吐
rabbitMQ:消息准确率高

# 表结构搭建

## 索引
唯一ID不能是纯数字id
- 安全问题

占领空闲ip->需要一张网卡->探针ip

## 什么时候分表
数据量很大的时候
最重要的是按照什么分

## 实体外键
> **在数据库中真正生成外键约束，保障两个实体间的关系一致性。**

追求性能不要开，追求一致性必须开

# 设计模式
## 单例模式
适用于全局只实例化一次的场景
- 配置管理、日志记录、数据库连接池



# jwt
明文的，别传重要信息
## 双令牌


# 镜像服务搭建
**Internal**：内部包，只能在内部被调用


## Docker配置代理
首先找到虚拟机的默认网关,然后再看主机vpn的端口

尝试
```bash
curl [虚拟机ip]:[vpn端口]
```

没报错就是成功

然后,因为docker pull /push 的代理被 systemd 接管，所以需要设置 systemd…
```bash
sudo mkdir -p /etc/systemd/system/docker.service.d
sudo vim /etc/systemd/system/docker.service.d/http-proxy.conf
```

```
[Service]
Environment="HTTP_PROXY=http://[GateWay IP]:[vpn地址]"
Environment="HTTPS_PROXY=http://[GateWay IP]:[vpn地址]"
```

然后重启服务
```bash
sudo systemctl daemon-reload
sudo systemctl restart docker
```

## Go环境搭建

下载go编译器
```
cd /opt

wget [https://golang.google.cn/dl/]

tar -xvf解压 
```

配置代理
vim /etc/profile 在文件后追加以下内容
```go
export GOPROXY=https://goproxy.cn 
export GOROOT=/opt/go 
export PATH=$PATH:$GOROOT/bin export 
GOPATH=/opt/go/pkg 
export PATH=$PATH:$GOPATH/bin 
```
退出并保存，刷新环境变量
` source /etc/profile`



## 虚拟机开放端口
服务要想让外部应用访问，需要开放端口
ubuntu举例:
```sh
ufw allow 8080
```


## 保存镜像
docker save -o [filename] imageName


## 制作诱捕镜像
编写DockerFile
如
```go
FROM golang:alpine AS builder  
  
# 构建可执行文件  
ENV CGO_ENABLED 0  
ENV GOPROXY https://goproxy.cn,direct  
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories  
  
WORKDIR /build  
ADD . .  
RUN go mod tidy  
RUN go build -o main  
  
#FROM alpine  
FROM scratch  
WORKDIR /app  
COPY --from=builder /build/main /app  
# RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories  
ADD . .  
CMD ["./main"]
```

然后
在目录下运行`docker build -t [image_name] .`

再保存为tar`docker save -o [name] xxx:latest`
> latest少了就解析不出标签了，不能少


**Docker API**:
go get github.com/docker/docker@v24.0.6
