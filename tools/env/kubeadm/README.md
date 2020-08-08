# k8s 集群 kubeadm 部署

## master
```
k8s-master-install.sh  
k8s-master-init.sh  
```

## node
```
k8s-node-install.sh  
k8s-node-init.sh       
k8s-node-join.sh     
```

## 注意

- kubeadm/master/kubeadm-init-master.sh 中的 apiserver-advertise-address 的 ip 要改成 master 节点的.

- kubeadm/node/kubeadm-join-node.sh 中的 ip\token\discovery-token-ca-cert-hash 要改成对应的:
    - master init 时会生成.
    - token 过期后可以用 kubeadm/master/kubeadm-token-ssl.sh 生成.

