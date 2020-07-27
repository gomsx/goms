# iamges 清理

# 所有的有用的 images 都 run 时，可以删除不在 run 的无用的 images

```
docker images -q | xargs docker rmi
```

<<<<<<< Updated upstream
=======
# 长时间运行，硬盘爆满时，删除重用超预期的 images

```
du -d 1 -hm /var/lib/docker/containers | sort -t $'\t' -k 1rn,1
```
>>>>>>> Stashed changes
