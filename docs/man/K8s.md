# K8s

## patch pv finalizers to null for delete
kubectl patch persistentvolume/pv-local-mysql -p '{"metadata":{"finalizers":null}}' -n default
kubectl patch persistentvolumeclaim/mysql-pv-mysql-sts-0 -p '{"metadata":{"finalizers":null}}' -n default

## remove evicted pod
kubectl get pods | grep Evicted | awk '{print $1}' | xargs kubectl delete pod
