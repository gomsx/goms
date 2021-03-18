# k8s 集群 kubeadm 部署

## 配置

要配置 kubeadm/master/config.yaml 中的版本和 master 节点 IP
```
version: v1.18.1
master_ip: 172.18.194.179
```

## 安装

1, master

```
k8s-master-install.sh
k8s-master-init.sh
```

2, node

```
k8s-node-install.sh
k8s-node-init.sh
```

## 建立集群

1, master 生成 node 节点使用的加入 cluster 命令

```
k8s-master-gen-join-cmd.sh
```

输出如下:
```
sudo kubeadm join 172.18.194.179:6443 --token db0znh.g2bjcipculsqygf --discovery-token-ca-cert-hash sha256:68c464257baaf208235e2bf1bceddccf6d9dbd1e5bb1c60f9302641f97b9414
```

2, node 使用加入 cluster 命令

必要时先 reset

```
sudo kubeadm reset
```

然后 join

```
sudo kubeadm join 172.18.194.179:6443 --token db0znh.g2bjcipculsqygf --discovery-token-ca-cert-hash sha256:68c464257baaf208235e2bf1bceddccf6d9dbd1e5bb1c60f9302641f97b9414
```
