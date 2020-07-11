# Istio

## istio 功能

- Connect (连接): 智能控制服务之间的调用流量，能够实现灰度升级、AB 测试和红黑部署等功能.
    Request Routing (请求路由),Fault Injection (故障注入),Traffic Shifting (流量转移),TCP Traffic Shifting (TCP 流量转移),Request Timeouts (请求超时),Circuit Breaking (限流),Mirroring (流量镜像),Ingress (流量入口),Egress (流量出口).

- Secure (安全加固): 自动为服务之间的调用提供认证、授权和加密.
    Certificate Management (证书管理),Authentication (鉴权),Authorizaition (授权).

- Control (控制): 应用用户定义的 policy，保证资源在消费者中公平分配.

- Observe (观察): 查看服务运行期间的各种数据，比如日志、监控和 tracing，了解服务的运行情况.
    Metrics,Logs,Tracing.

## istio 组件

- Envoy - Istio 的**数据面**. Sidecar 代理微服务处理进出流量，来自集群内的服务间，或者集群内和集群外的服务间.这些代理组成安全的服务网格，提供丰富的空能，如discovery, rich layer-7 routing, circuit breakers, policy enforcement and telemetry recording/reporting functions.

> 服务网格不是 overlay 网络.它简化并加强应用内的微服务的通讯，通过底层框架提供的网络.

- Istiod - Istio 的**控制面**. 提供 service discovery (服务发现), configuration (配置) 和 certificate management (证书管理).

    - Pilot - 动态配置代理.

    - Citadel - 证书的分发和 rotation.

    - Galley - 负责 validating, ingesting, aggregating, transforming 和 distributing config.

- Operator - 提供用户友好的选项来操作服务网格.

## istio 客户端

istioctl

## 常用命令

```
istioctl --help

register        Registers a service instance (e.g. VM) joining the mesh  
deregister      De-registers a service instance
analyze         Analyze Istio configuration and print validation messages

manifest        Commands related to Istio manifests
operator        Commands related to Istio operator controller.
profile         Commands related to Istio configuration profiles

kube-inject     Inject Envoy sidecar into Kubernetes pod resources

proxy-config    Retrieve information about proxy configuration from Envoy [kube only]
proxy-status    Retrieves the synchronization status of each Envoy in the mesh [kube only]
``` 

## 参考
https://github.com/istio/istio  
https://istio.io  
http://www.uml.org.cn/wfw/201909063.asp  

