# eK8s



## k8s 的结构

- **控制平面**  
    - Kubernetes 主控组件
        - kube-apiserver  
        - kube-controller-manager  
        - kube-scheduler
    - kubelet 进程
- **数据平面**

## k8s 资源/对象
- **Pod**  对应于传统的引用程序,应用程序通常包含多个进程,而进程对应于 docker,所以 pod 也通常包含多个 docker.
- **Service**  
    - **ClusterIP**  默认的servie类型,pod 的代理,含ClusterIP/ClusterPort(虚拟),用于集群内发布服务.
    - **NodePort**  依赖 ClusterIP, 连接 NodeIP/NodePort(真实) 和 ClusterIP/ClusterPort(虚拟),用于集群外发布服务.
    - **LoadBalancer**  依赖 NodePort 和外部负载均衡器, 连接 client(用户) 和 NodePort,用于集群外发布带负载均衡器的服务.
    - **ExternalName**  依赖外部服务,没有ClusterIP 和 NodePort,以DNS方式访问,用于集群内发布集群外部的服务.
- **Volume**
- **Namespace**   



## k8s 对象控制器
- Deployment
- DaemonSet   ds
- StatefulSet
- ReplicaSet  rs
- Job

# k8s 使用

k8s 的客户端分为 
- kubectl 
- dashboard 
- sdk




# 常用简写

configmaps cm  
daemonsets ds  
deployments deploy  
endpoints ep  
event ev  
horizontalpodautoscalers hpa  
ingresses ing  
namespaces ns  
nodes no  
persistentvolumeclainms pvc  
persistentvolumes pv  
pods po  
replicasets rs  
replicationcondtrollers  
resourcequotas quota  
serviceaccounts sa  
services svc  

##