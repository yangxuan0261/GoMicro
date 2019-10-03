GitHub 地址 - https://github.com/jeckbjy/test-go/tree/master/etcd

---

# 基于etcd v3 构建服务发现以及集中配置分发

- 官网
  - https://github.com/etcd-io/etcd
  - https://etcd.readthedocs.io/en/latest/

- 一些文章
  - API的使用   https://yuerblog.cc/2017/12/12/etcd-v3-sdk-usage/ 
  - 应用场景:   https://www.infoq.cn/article/etcd-interpretation-application-scenario-implement-principle
  - 比较:       https://www.servercoder.com/2018/03/30/consul-vs-zookeeper-etcd/
  - 服务发现:   http://daizuozhuo.github.io/etcd-service-discovery/
  - v3服务发现: https://github.com/moonlong/etcd-discovery
  - API用法:    https://yuerblog.cc/2017/12/12/etcd-v3-sdk-usage/
  - 原理解析:   https://yuerblog.cc/2017/12/10/principle-about-etcd-v3/

- 服务发现原理：
  - 服务的提供者(provider)使用Put注册服务,并使用Lease保持连接
  - 服务的使用者(consumer)使用Get首次获取服务,并使用Watcher监视Key的变化,根据api提供的不同的Action,做相应操作,比如连接新的节点,删除新的节点等
  - 对于某些节点,既可能是服务的提供者也可能是某些服务的使用者

- 运行例子
  - 编译:./build.sh 生成到../bin/discovery
  - cd ../bin 在bin目录执行
  - 启动etcd
  - discovery --role=1 启动provider
  - discovery --role=2 启动comsumer