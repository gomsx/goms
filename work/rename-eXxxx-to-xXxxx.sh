#!/bin/bash
set -x
# set -e

# 当前目录路径
pwdx="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
echo "==> pwdx:$pwdx"

# 当前项目路径 prox
prox="$(cd "$pwdx/.." && pwd)"
echo "==> prox:$prox"

# 搜索 eXxxx 目录
dst="e[A-Z]"
edirs="$(ls | grep "$dst")"
edirs=("$edirs")
echo "--> edirs: $edirs"
# 处理文件
for dir in $edirs; do
	old="$dir"
	new="$(echo "$old" | sed "s/^e/x/g")"
	echo "$new"
	mv "$old" "$new"
done
# exit

# 处理文件
hand_file(){
	files="$1"
	for file in $files; do
		
		olds=("eapi" "econf" "edocker" "eflag" "egrpc" "ehttp" "eistio" "ek8s" "emodule" "emysql" "eredis" "etest" "eyaml")
		news=("xapi" "xconf" "xdocker" "xflag" "xgrpc" "xhttp" "xistio" "xk8s" "xmodule" "xmysql" "xredis" "xtest" "xyaml")

		OLDS=("eApi" "eConf" "eDocker" "eFlag" "eGrpc" "eHttp" "eIstio" "eK8s" "eModule" "eMysql" "eRedis" "eTest" "eYaml")
		NEWS=("xApi" "xConf" "xDocker" "xFlag" "xGrpc" "xHttp" "xIstio" "xK8s" "xModule" "xMysql" "xRedis" "xTest" "xYaml")

		for ((i = 0; i < ${#olds[*]}; i++));
		do
			sed -i "s/${olds[i]}/${news[i]}/g" "$file"
			sed -i "s/${OLDS[i]}/${NEWS[i]}/g" "$file"
		done
	done
}

# 搜索 makefile 文件
dst="*makefile"
file_mks="$(find "$prox" -name "$dst" | grep -v  work | grep -v tools)"
file_mks=("$file_mks")
echo "--> file_makefile: $file_mks"
# 处理文件
for file in $file_mks; do
	sed -n "/APP=/p" "$file"
	sed -i "/APP\=/{ s/\=e/\=x/g }" "$file"
done
hand_file "$file_mks"
# exit

# 搜索 dockerfile 文件
dst="*dockerfile"
file_dks="$(find "$prox" -name "$dst")"
file_dks=("$file_dks")
echo "--> file_dockerfile: $file_dks"
# 处理文件
for file in $file_dks; do
	sed -i "s/\/configs\//\/configs/g" "$file"
done
hand_file "$file_dks"
# exit

# 搜索 go 文件
dst="*.go"
file_gos="$(find "$prox" -name "$dst" | grep -v ".pb.")"
file_gos=("$file_gos")
echo "--> file_gofile: $file_gos"
# 处理文件
hand_file "$file_gos"
# exit

# 搜索 sh 文件
dst="*.sh"
file_shs="$(find "$prox" -name "$dst" | grep -v  work | grep -v tools)"
file_shs=("$file_shs")
echo "--> file bash: $file_shs"
# 处理文件
hand_file "$file_shs"
# exit

# 搜索 mk 文件
dst="*.md"
file_mds="$(find "$prox" -name "$dst")"
file_mds=("$file_mds")
echo "--> file markdown: $file_mds"
# 处理文件
hand_file "$file_mds"
# exit
