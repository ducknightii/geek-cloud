### 题目
- 构建本地镜像 
- 编写 Dockerfile 将练习 2.2 编写的 httpserver 容器化 
- 将镜像推送至 docker 官方镜像仓库 
- 通过 docker 命令本地启动 httpserver 
- 通过 nsenter 进入容器查看 IP 配置

### 答案

[Dockerfile](../module2/httpserver/Dockerfile)
[Makefile](../module2/httpserver/Makefile)

```
TAG=1.0.0 make docker-build
TAG=1.0.0 make docker-push

docker run --name httpserver -d ducknight/httpserver:1.0.0

docker inspect httpserver

nsenter -n -t{PID}
```

```
root@VM-4-12-ubuntu:/home/ubuntu/github.com/ducknightii/geek-cloud/module2/httpserver# ip addr
1: lo: <LOOPBACK,UP,LOWER_UP> mtu 65536 qdisc noqueue state UNKNOWN group default qlen 1000
    link/loopback 00:00:00:00:00:00 brd 00:00:00:00:00:00
    inet 127.0.0.1/8 scope host lo
       valid_lft forever preferred_lft forever
20: eth0@if21: <BROADCAST,MULTICAST,UP,LOWER_UP> mtu 1500 qdisc noqueue state UP group default
    link/ether 02:42:ac:11:00:02 brd ff:ff:ff:ff:ff:ff link-netnsid 0
    inet 172.17.0.2/16 brd 172.17.255.255 scope global eth0
       valid_lft forever preferred_lft forever
```