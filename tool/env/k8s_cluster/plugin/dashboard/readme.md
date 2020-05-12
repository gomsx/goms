

## 部署

https://blog.csdn.net/networken/article/details/85607593
https://kubernetes.io/docs/tasks/access-application-cluster/web-ui-dashboard/
https://github.com/kubernetes/dashboard
https://www.jianshu.com/p/be2a12a8bc0b


```

sudo kubectl apply -f 2_0_0/recommended-my.yaml

sudo kubectl get rs,pod,deploy,svc,ep -n kubernetes-dashboard

sudo kubectl delete -f 2_0_0/recommended-my.yaml

```