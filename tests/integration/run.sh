#!/bin/bash
set -x
set -e

# 项目 root
prox="$(cd ../../ && pwd)"
echo "--> prox:${prox}"

# set GOMS_TEST_ROOT
test_root="/tmp/goms-test$(date +%N)"
echo "--> test root: ${test_root}"
[ -d "${test_root}" ] || mkdir -p "${test_root}"
ls ${test_root}
export GOMS_TEST_ROOT="${test_root}"

# set testx
testx="${test_root}/mock/client/script"
[ -d "${testx}" ] || mkdir -p "${testx}"
cp -r ${prox}/tests/mock/client/script ${testx}/..
ls ${testx}

# function
service_running() {
	pingx=$1
	limit_times=$2
	delay=$3

	for ((try_times = 0; try_times < limit_times; try_times++)); do
		result_code="$($pingx)"
		if [ "200" == "${result_code}" ]; then
			echo "true"
			return
		fi
		sleep ${delay}
	done

	echo "false"
	return
}

kill_app_must() {
	app="$1"
	pid="$(ps -u "${USER}" | grep "${app}" | grep -v grep | awk '{print $1}')"
	[ -n "${pid}" ] && kill -9 "${pid}"
	# must 逻辑,此处捕获错误码,阻止上传
	echo "--> exit code: $?"
	return
}

do_test() {
	ver="$1"
	cd ${testx}
	bash ./test-api-http.sh "0.01" "${ver}"
	bash ./test-api-grpc.sh "0.01" "${ver}"
	return
}

test_by_one() {
	app="$1"
	cmd="$2"
	ver="$3"

	echo "--> 清理"
	kill_app_must "${app}"

	echo "--> 运行"
	make compile
	./${app} &
	sleep 2s

	echo "--> 测活"
	if [ "true" != "$(service_running "$cmd" 1 3s)" ]; then
		echo "--> 测活失败: $?"
		kill_app_must "${app}"
		make clean
		sleep 2s
		return 255
	fi
	echo "--> 测活成功: $?"

	echo "--> 测试"
	(do_test "${ver}")

	echo "--> 清理"
	kill_app_must "${app}"
	make clean
	sleep 2s
	return 0
}

# prox/...
dirs=("eApi" "eTest" "eRedis")
apps=("eapi" "etest" "eredis")
# pingv="curl -w %{http_code} localhost:8080/v1/ping"
# ping="curl -w %{http_code} localhost:8080/ping"
# pings=("\$pingv" "\$ping" "\$ping")
pingv=("curl -o /dev/null -w %{http_code} localhost:8080/v1/ping")
ping=("curl -o /dev/null -w %{http_code} localhost:8080/ping")
pings=("${pingv}" "${ping}" "${ping}")
vers=("v1" "" "")

# test.log
log="$(pwd)/test.log"
echo -e "测试时间:\n==> $(date -R)" >${log}
echo -e "测试结果:" >>${log}

for ((i = 0; i < ${#apps[*]}; i++)); do
	kill_app_must "${apps[i]}"
done

for ((i = 0; i < ${#apps[*]}; i++)); do
	echo "==> 测试 ${dirs[i]}" >>${log}

	cd ${prox}/${dirs[i]}/build
	test_by_one "${apps[i]}" "${pings[i]}" "${vers[i]}"
	ret=$?
	if [ "0" != "$ret" ]; then
		echo "==< 测试失败" >>${log}
		continue
	fi
	echo "==< 测试成功" >>${log}
done

cat "${log}"
