#!/bin/bash
set -e
set -x
set -u

# .sh所在目录
DIR=$(cd "$(dirname "$0")";pwd)
echo $DIR
PRODIR=$(cd "$DIR"/../;pwd)
echo $PRODIR;   

# 查看环境变量
echo $HOME; echo $PATH; which go;   
  
cd $PRODIR/eMysql/docker; chmod +x ./build.sh; ./build.sh  # eMysql
cd $PRODIR/eConf/docker; chmod +x ./build.sh; ./build.sh   # eConf
cd $PRODIR/eGrpc/docker; chmod +x ./build.sh; ./build.sh   # eGrpc
cd $PRODIR/eHttp/docker; chmod +x ./build.sh; ./build.sh   # eHttp
cd $PRODIR/eFlag/docker; chmod +x ./build.sh; ./build.sh   # eFlag
cd $PRODIR/eYaml/docker; chmod +x ./build.sh; ./build.sh   # eYaml
cd $PRODIR/eDocker/docker; chmod +x ./build.sh; ./build.sh # eDocker 

cd $PRODIR/eModule/cmd; go run .                           # eModule

cd $PRODIR/build/cicd/docker; chmod +x ./build.sh; ./build.sh    # ubuntu




