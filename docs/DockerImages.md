
# iamges 清理

# 所有的有用的 images 都 run 时，可以删除不在 run 的无用的 images

```
docker images -q | xargs docker rmi
```