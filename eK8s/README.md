# eK8s

## k8s 的结构

核心组件组成：  

- master  
    - **etcd** - 保存了整个集群的状态.  
    - **apiserver** - 提供了资源操作的唯一入口，并提供认证、授权、访问控制、API注册和发现等机制.  
    - **controller manager** - 负责维护集群的状态，比如故障检测、自动扩展、滚动更新等.  
    - **scheduler** - 负责资源的调度，按照预定的调度策略将 Pod 调度到相应的机器上.  
- node  
    - **kubelet** - 负责维护容器的生命周期，同时也负责Volume(CSI)和网络(CNI)的管理.    
    - **kube-proxy** - 负责为Service提供cluster内部的服务发现和负载均衡.   
    - **container runtime** - 负责镜像管理以及Pod和容器的真正运行(CRI).  

## k8s API 资源/对象

- **Pod** - 对应于传统的应用程序,应用程序通常包含多个进程,而进程对应于 docker,所以 pod 也通常包含多个 docker.
    Pod 控制器种类:
    - **Deployment** - 控制多个 pod 副本.
    - **DaemonSet** - 作为 daemon 运行的任务.
    - **StatefulSet** - 有状态的任务.
    - **Job** - 运行一次的任务.
    - **CronJob** - 周期性运行的任务.
- **Service** - 反向代理，提供稳定的 IP 或 DNS,种类：  
    - ClusterIP - 默认的 servie 类型,pod 的代理,含 ClusterIP/ClusterPort(虚拟),用于集群内发布服务.
    - NodePort - 依赖 ClusterIP, 连接 NodeIP/NodePort(真实) 和 ClusterIP/ClusterPort(虚拟),用于集群外发布服务.
    - LoadBalancer - 依赖 NodePort 和外部负载均衡器, 连接 client(用户) 和 NodePort,用于集群外发布带负载均衡器的服务.
    - ExternalName - 依赖外部服务,没有 ClusterIP 和 NodePort,以 DNS 方式访问,用于集群内发布集群外部的服务.
- **Volume** - 是 pod 中能被多个容器访问的共享目录。种类：
    - emptyDir - 和 pod 同生命周期，存储介质可能是内存或硬盘.
    - hostPaht - 宿主机上的目录或文件.
    - persistentVolumeClaim - 从 PV(PersistentVolume) 中申请的存储空间，PV 通常是一种网络存储.
    - ...
- **Namespace** - 将物理集群划分为多个虚拟集群.

> Kubernetes的三种IP  
Node IP： Node节点的IP地址  
Cluster IP: Service的IP地址  
Pod IP: Pod的IP地址  

## k8s 客户端

- kubectl 
- dashboard 
- sdk 

## 调试 

log
```
journalctl -f -u kubelet
kubectl logs service/redis-svc  
```

login
```
kubectl exec -it pod/user-deploy-7fc88fdcbf-gn7kl -- /bin/sh  
kubectl exec -it deployment.extensions/mysql-deploy -- /bin/sh  
kubectl exec -it service/user-svc -- /bin/sh  
```

object
```
kubectl get all
kubectl get rs,pod,deploy,sts,svc,ep
kubectl get event --all-namespace
kubectl get pod -o wide
kubectl describe node
kubectl get ns
```

net
```
netstat -s
netstat -s | grep LISTEN
netstat -s | grep timeout
netstat -s | grep dropped
```

## 常用简写
```
$ kubectl api-resources

 no          Node
 ns          Namespace
 cm          ConfigMap
 po          Pod
 deploy      Deployment
 rs          ReplicaSet
 sts         StatefulSet
 ds          DaemonSet
             Job
 cj          CronJob
 rc          ReplicationController
 ing         Ingress
 svc         Service
 ep          Endpoints
 pv          PersistentVolume

 ev          Event
 sa          ServiceAccount
 quota       ResourceQuota
 psp         PodSecurityPolicy
 ```

## 注意

namespace 可以用 kubens 命令锚定

```
# 安装
git clone https://github.com/ahmetb/kubectx /opt/kubectx
ln -s /opt/kubectx/kubectx /usr/local/bin/kubectx
ln -s /opt/kubectx/kubens /usr/local/bin/kubens

# 查看
kubens

# 绑定
kubens goms-test
```
## 参考

https://kubernetes.io  
https://blog.kubernetes.io  
http://www.coderdocument.com/docs/kubernetes/v1.14/index.html  
https://github.com/kubernetes/community/blob/master/sig-list.md#master-sig-list

