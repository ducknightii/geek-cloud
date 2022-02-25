
### CRI
容器运行时(Container Runtime)，运行于集群的每一个节点，负责容器的整个生命周期。

- 一组grpc服务： 运行时服务(RuntimeService), 镜像服务(ImageService)
- OCI(Open Container Initiative) 开放容器计划
- dockershim、 containerd、 CRI-O


## CNI (Container Network Interface)

k8s网络设计原则：
- 所有的Pod能够不通过NAT就能直接访问
- 所有的节点能够不通过NAT就能互相访问
- 容器内能看见的ip地址和外部组件看到的容器IP是一样的

常见插件：
- IPAM: IP地址分配
- 主插件：网卡设置
  - bridge：创建网桥，并把主机端口和容器端口插入网桥
  - ipvlan：为容器添加ipvlan网口
- Mate：附加功能
  - portmap：设置主机端口和容器端口映射
  - bandwidth：利用Linux Traffic Control 限流
  - firewall：通过iptables或firewalld为容器设置防火墙规则

## CSI
存储

#### Rook
开源分布式存储编排系统
