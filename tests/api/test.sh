#!/bin/bash
# set -x
set -e

service_running() {
	pingx=$1
	limit_times=$2
	delay=$3

	for ((try_times = 0; try_times < limit_times; try_times++)); do
		result_code="$($pingx)"
		if [ "200" == "$result_code" ]; then
			echo "true"
			return
		fi
		sleep "$delay"
	done

	echo "false"
	return
}

kill_app() {
	app="$1"
	pid="$(pgrep "$app" | grep -v grep | awk '{print $1}')"
	kill -9 "$pid"
	return
}

do_test() {
	ver="$1"
	test=$prox/tests/client/script/test.sh
	chmod +x $test
	eval "$test" "3" "0.01" "$ver"
	return
}

test_by_one() {
	app="$1"
	cmd="$2"
	ver="$3"

	echo "--> 运行"
	make compile
	./"$app" &
	sleep 2s

	echo "--> 测活"
	if [ "true" != "$(service_running "$cmd" 3 3s)" ]; then
		echo "--> $?"
		kill_app "$app"
		make clean
		return 255
	fi

	echo "--> 测试"
	do_test "$ver"

	echo "--> 清理"
	kill_app "$app"
	make clean
	return 0
}

# 当前目录路径
pwdx="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
echo "==> pwdx:$pwdx"

# 当前项目路径 pro
prox="$(cd "$pwdx/../.." && pwd)"
echo "==> pro:$prox"

# args
dirs=("eApi" "eTest" "eRedis")
apps=("eapi" "etest" "eredis")
# pingv="curl -w %{http_code} localhost:8080/v1/ping"
# ping="curl -w %{http_code} localhost:8080/ping"
# pings=("\$pingv" "\$ping" "\$ping")
pingv=("curl  -o /dev/null  -w %{http_code} localhost:8080/v1/ping")
ping=("curl  -o /dev/null  -w %{http_code} localhost:8080/ping")
pings=("$pingv" "$ping" "$ping")
vers=("v1" "" "")

for ((i = 0; i < ${#apps[*]}; i++)); do
	echo "==> 测试 ${dirs[i]}"
	cd $prox/${dirs[i]}/build
	test_by_one "${apps[i]}" "${pings[i]}" "${vers[i]}"
	ret=$?
	if [ "0" != "$ret" ]; then
		echo "==< 失败"
		continue
	fi
	echo "==< 成功"
done
