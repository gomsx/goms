#!/bin/bash
set -x
set -e

# 项目目录 pro dir
PD="$(cd ../../ && pwd)"
echo "--> PD:${PD}"

# set GOMS_TEST_DIR
TEST_SN="$(date +%N)"
TD="/tmp/goms-test${TEST_SN}"
echo "--> TEST DIR: ${TD}"
[ -d "${TD}" ] || mkdir -p "${TD}"
export GOMS_TEST_DIR="${TD}"
ls ${TD}

# set TEST_IT
TEST_IT="${TD}/integration"
[ -d "${TEST_IT}" ] || mkdir -p "${TEST_IT}"
cp -r ${PD}/tests/mock/client/script ${TEST_IT}
ls ${TEST_IT}

# test log
log="$(cd ./ && pwd)/test.log"
echo "log file: $log"

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
	(cd ${TEST_IT}/script && bash test-api-http.sh "0.01" "${ver}")
	(cd ${TEST_IT}/script && bash test-api-grpc.sh "0.01" "${ver}")
	return
}

test_by_one() {
	app="$1"
	cmd="$2"
	ver="$3"

	echo "--> 清理" | tee -a ${log}
	kill_app_must "${app}"

	echo "--> 运行" | tee -a ${log}
	make compile
	./${app} &
	sleep 3s

	echo "--> 测活"
	if [ "true" != "$(service_running "$cmd" 1 3s)" ]; then
		echo "--> 测活失败: 清理&退出" | tee -a ${log}
		kill_app_must "${app}"
		make clean
		sleep 2s
		return 255
	fi
	echo "--> 测活成功" | tee -a ${log}

	echo "--> 测试" | tee -a ${log}
	(do_test "${ver}")

	echo "--> 清理" | tee -a ${log}
	kill_app_must "${app}"
	make clean
	sleep 2s
	return 0
}

# PD/...
dirs=("eApi" "eTest" "eRedis")
apps=("eapi" "etest" "eredis")
vers=("v1" "" "")
pingv=("curl -o /dev/null -w %{http_code} localhost:8080/v1/ping")
ping=("curl -o /dev/null -w %{http_code} localhost:8080/ping")
pings=("${pingv}" "${ping}" "${ping}")
# pings=("\$pingv" "\$ping" "\$ping")

echo -e "测试编号: ${TEST_SN}" | tee ${log}
echo -e "测试时间: $(date -R)" | tee -a ${log}
echo -e "\n清理:" | tee -a ${log}

for ((i = 0; i < ${#apps[*]}; i++)); do
	echo "==> 清理 ${dirs[i]}" | tee -a ${log}
	kill_app_must "${apps[i]}"
done

echo -e "\n测试:" | tee -a ${log}

for ((i = 0; i < ${#apps[*]}; i++)); do
	echo -e "\n${i}.测试 ${dirs[i]}" | tee -a ${log}

	echo "==> 测试开始" | tee -a ${log}
	cd ${PD}/${dirs[i]}/build
	test_by_one "${apps[i]}" "${pings[i]}" "${vers[i]}"
	ret=$?
	if [ "0" != "$ret" ]; then
		echo "==< 测试失败" | tee -a ${log}
		continue
	fi
	echo "==< 测试成功" | tee -a ${log}
done

echo -e "\n\n"
cat "${log}"
