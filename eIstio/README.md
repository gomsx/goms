# Istio

## istio 功能

- Connect (连接)  
    智能控制服务之间的调用流量，能够实现灰度升级、AB 测试和红黑部署等功能.
    Request Routing (请求路由),Fault Injection (故障注入),Traffic Shifting (流量转移),TCP Traffic Shifting (TCP 流量转移),Request Timeouts (请求超时),Circuit Breaking (限流),Mirroring (流量镜像),Ingress (流量入口),Egress (流量出口).

- Secure (安全加固)  
    自动为服务之间的调用提供认证、授权和加密.
    Certificate Management (证书管理),Authentication (鉴权),Authorizaition (授权).

- Control (控制)  
    应用用户定义的 policy，保证资源在消费者中公平分配.

- Observe (观察)  
    查看服务运行期间的各种数据，比如日志、监控和 tracing，了解服务的运行情况.
    Metrics,Logs,Tracing.

## istio 组件

- Envoy - Istio 的**数据面**  
    Sidecar 代理微服务处理进出流量，来自集群内的服务间，或者集群内和集群外的服务间.这些代理组成安全的服务网格，提供丰富的空能，如discovery, rich layer-7 routing, circuit breakers, policy enforcement and telemetry recording/reporting functions.

> 服务网格不是 overlay 网络.它简化并加强应用内的微服务的通讯，通过底层框架提供的网络.

- Istiod - Istio 的**控制面**  
    提供 service discovery (服务发现), configuration (配置) 和 certificate management (证书管理).

    - Pilot 动态配置代理.

    - Citadel 证书的分发和 rotation.

    - Galley 负责 validating, ingesting, aggregating, transforming 和 distributing config.

- Operator  
    提供用户友好的选项来操作服务网格.

## istio API 资源/对象  

- VirtualService  
    配置影响流量路由的参数,定义了对特定目标服务的一组流量规则.形式上是一个虚拟服务,将满足条件的流量转发到对应的服务后端(一个服务,或者是 DestinationRule 中定义的服务子集).
- DestinationRule  
    配置目标规则,定义了在路由发生后,应用于服务流量的策略.具体指定了负载平衡的配置,边车的连接池大小以及离群检测值,以从负载平衡池中检测并清除不正常的后端服务.
- Gateway  
    服务网关,描述了负载均衡器,在网格的边缘运行,来接收传入或传出的 HTTP/TCP 连接.
- ServiceEntry  
    外部服务配置,在网格的服务注册表中添加条目,以便网格中的服务可以自动发现/访问/路由到这些条目指定的服务.

## istio 客户端

istioctl

## debug
```
kubectl logs pod/user-deploy-79df86bcf-qd6dm istio-init
kubectl logs pod/user-deploy-79df86bcf-qd6dm istio-proxy
kubectl logs pod/user-deploy-79df86bcf-qd6dm user
```

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

## 常用简写

```
$ kubectl api-resources

virtualservices                   vs         
destinationrules                  dr         
gateways                          gw         
serviceentries                    se         
workloadentries                   we         
sidecars                                     
```

## 参考
https://github.com/istio/istio  
https://istio.io  
http://www.uml.org.cn/wfw/201909063.asp  

