
# 查看工程目录
PRODIR:=$(shell cd ../;pwd)# 工程目录
$(warning $(PRODIR))   

# 查看环境变量
$(warning $(HOME))
$(warning $(PATH))
$(warning $(shell which go))


all:compile docker push clean

compile:
	cd $(PRODIR)/eRedis/build;make compile 
	cd $(PRODIR)/eMysql/build;make compile 
	cd $(PRODIR)/eConf/build;make compile 
	cd $(PRODIR)/eGrpc/build;make compile 
	cd $(PRODIR)/eHttp/build;make compile 
	cd $(PRODIR)/eFlag/build;make compile 
	cd $(PRODIR)/eYaml/build;make compile 
	cd $(PRODIR)/eModule/build;make compile 
	cd $(PRODIR)/eDocker/build;make compile 

docker: 
	cd $(PRODIR)/eRedis/build;make docker 
	cd $(PRODIR)/eMysql/build;make docker 
	cd $(PRODIR)/eConf/build;make docker 
	cd $(PRODIR)/eGrpc/build;make docker 
	cd $(PRODIR)/eHttp/build;make docker 
	cd $(PRODIR)/eFlag/build;make docker 
	cd $(PRODIR)/eYaml/build;make docker 
	cd $(PRODIR)/eModule/build;make docker 
	cd $(PRODIR)/eDocker/build;make docker 

push: 
	cd $(PRODIR)/eRedis/build;make push
	cd $(PRODIR)/eMysql/build;make push
	cd $(PRODIR)/eConf/build;make push
	cd $(PRODIR)/eGrpc/build;make push
	cd $(PRODIR)/eHttp/build;make push
	cd $(PRODIR)/eFlag/build;make push
	cd $(PRODIR)/eYaml/build;make push
	cd $(PRODIR)/eModule/build;make push
	cd $(PRODIR)/eDocker/build;make push

clean:
	cd $(PRODIR)/eRedis/build;make clean
	cd $(PRODIR)/eMysql/build;make clean
	cd $(PRODIR)/eConf/build;make clean
	cd $(PRODIR)/eGrpc/build;make clean
	cd $(PRODIR)/eHttp/build;make clean
	cd $(PRODIR)/eFlag/build;make clean
	cd $(PRODIR)/eYaml/build;make clean
	cd $(PRODIR)/eModule/build;make clean
	cd $(PRODIR)/eDocker/build;make clean