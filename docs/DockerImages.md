# iamges 清理

# 所有有用的 images 都在 run 时，可以删除无用的不在 run 的 images

```
docker images -q | xargs docker rmi
```

# 长时间运行，硬盘爆满时，删除容量暴涨的 docker

```
du -d 1 -hm /var/lib/docker/containers | sort -t $'\t' -k 1rn,1
```
