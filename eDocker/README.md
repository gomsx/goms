# eDocker

Docker 的使用,通过一个 demo 了解 dockerfile 的编写.还涉及 docker 的一些设计细节.

容器化技术,使用 docker.


## docker

- Docker image  -  数据层的集合  
- Docker file  -  定义各种镜像层  
- Docker engine  -  构建和管理镜像  
- Docker registry  -   存储镜像的地方  
- Container  -  运行中的镜像实例  


## dockerfile


docker build -t image_name -f dockerfile_path  build_root_path
> build 时指定一个构建的根路径,此处的 build_root_path,COPY/ADD时,会在这个路径下查找.

## 注意

1, 把用户加入docker组后就不用 sudo 来执行 docker 命令
```
//设置
sudo groupadd docker
sudo gpasswd -a ${USER} docker
sudo systemctl restart docker

//确认
cat /etc/group | grep docker | grep ${USER}
docker version
```

2, 有些 docker registry 在国外，要设置国内镜像
```
cat /etc/docker/daemon.json
{
        "registry-mirrors": ["http://hub-mirror.c.163.com"]
}
```