etcd
---

- 主要功能
  - k-v存储
  - 监听机制
  - key过期与续约机制，用于监控和服务发现
    - 强一致性、高可用的服务存储目录。Raft协议
    - 注册服务以及心跳监控健康状态
  - cas & cad 用于分布式锁和leader选举 