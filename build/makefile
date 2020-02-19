
# 查看工程目录
PRODIR:=$(shell cd ../;pwd)# 工程目录
$(warning $(PRODIR))   

# 查看环境变量
$(warning $(HOME))
$(warning $(PATH))
$(warning $(shell which go))


all:compile docker push clean

compile:
	cd $(PRODIR)/eMysql/build;make compile 
	cd $(PRODIR)/eConf/build;make compile 
	cd $(PRODIR)/eGrpc/build;make compile 
	cd $(PRODIR)/eHttp/build;make compile 
	cd $(PRODIR)/eFlag/build;make compile 

docker: 
	cd $(PRODIR)/eMysql/build;make docker 
	cd $(PRODIR)/eConf/build;make docker 
	cd $(PRODIR)/eGrpc/build;make docker 
	cd $(PRODIR)/eHttp/build;make docker 
	cd $(PRODIR)/eFlag/build;make docker 

push: 
	cd $(PRODIR)/eMysql/build;make push
	cd $(PRODIR)/eConf/build;make push
	cd $(PRODIR)/eGrpc/build;make push
	cd $(PRODIR)/eHttp/build;make push
	cd $(PRODIR)/eFlag/build;make push

clean:
	cd $(PRODIR)/eMysql/build;make clean
	cd $(PRODIR)/eConf/build;make clean
	cd $(PRODIR)/eGrpc/build;make clean
	cd $(PRODIR)/eHttp/build;make clean
	cd $(PRODIR)/eFlag/build;make clean

# 构建

# cd $PRODIR/eMysql/docker; chmod +x ./build.sh; ./build.sh  # eMysql
# cd $PRODIR/eConf/docker; chmod +x ./build.sh; ./build.sh   # eConf
# cd $PRODIR/eGrpc/docker; chmod +x ./build.sh; ./build.sh   # eGrpc
# cd $PRODIR/eHttp/docker; chmod +x ./build.sh; ./build.sh   # eHttp
# cd $PRODIR/eFlag/docker; chmod +x ./build.sh; ./build.sh   # eFlag
# cd $PRODIR/eYaml/docker; chmod +x ./build.sh; ./build.sh   # eYaml
# cd $PRODIR/eDocker/docker; chmod +x ./build.sh; ./build.sh # eDocker 

# cd $PRODIR/eModule/cmd; go run .                           # eModule

# cd $PRODIR/build/cicd/docker; chmod +x ./build.sh; ./build.sh    # ubuntu




