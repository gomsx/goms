# initdbd debug

## log

kubectl logs initdbd-deploy-58d7fbff4f-vn9qb
kubectl logs initdbd-deploy-58d7fbff4f-vn9qb | wc -l

## patch mysql

```
env:
- name: MYSQL_ALLOW_EMPTY_PASSWORD
    value: "yes"
- name: MYSQL_ROOT_HOST
    value: "%"
```
MYSQL_ALLOW_EMPTY_PASSWORD 设为 yes 使 root 用户可以 EMPTY_PASSWORD 登入
MYSQL_ROOT_HOST 设为 % 使 root 用户可以远程登入
