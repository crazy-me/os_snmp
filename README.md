## 基于grpc 的snmp服务

#### 安装Consul

* 从[官网](https://www.consul.io)下载安装 consul_1.5.3_linux_amd64.zip
* 解压 unzip consul_1.5.3_linux_amd64.zip
* 查看版本 consul -v
* 开发模式启动 consul agent -dev -ui -client 0.0.0.0
* 访问http://127.0.0.1:8500

#### docker安装Consul

* docker 拉取 consul 镜像 docker pull consul

* 创建并启动容器 默认是以开发模式启动，数据保存在内存中

```
docker run -d --name=consul -e CONSUL_BIND_INTERFACE=eth0 consul:1.7.3
```

#### docker部署

* 下载并编译

```sh
go get github.com/crazy-me/os_snmp
cd os_snmp && go build
```

* 修改配置文件

```sh

# System configuration
system:
  env: 'public'
  port: 8800

zap:
  log-prefix: '[OS-SNMP]'   # 日志前缀
  log-path: 'logs'          # 日志目录名称
  log-max-size: 10          # 在进行切割之前，日志文件的最大大小（以MB为单位）
  log-max-backups: 5        # 保留旧文件的最大个数
  log-max-age: 30           # 保留旧文件的最大天数
  log-compress : false      # 是否压缩/归档旧文件

consul:
  address: '127.0.0.1:8500' # consul 地址
  interval: 5               # consul 健康检查间隔
  expire-time: 2            # consul 服务异常过期时间
```

* Dockerfile文件

```sh
FROM golang:1.15
WORKDIR /go/src/os_snmp
COPY app.yaml .
COPY os_snmp .
COPY conf ./conf
COPY script ./script
EXPOSE 8800
#RUN go env -w GOPROXY=https://goproxy.cn,direct 
#RUN go build -o os_snmp
CMD ["/bin/bash", "/go/src/os_snmp/script/build.sh"]
```

* 构建镜像

```sh
docker build -t os_snmp .
```

* 运行并启动容器

```sh
docker run -d -p 8800:8800 --name os_snmp os_snmp
```





