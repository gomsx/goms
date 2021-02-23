#!/bin/bash
set -x
set -e

# 工作目录路径
WD="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
echo "--> work dir: ${WD}"

# 项目目录路径
PD="$(cd "$WD/.." && pwd)"
echo "--> pro dir: ${PD}"

# cd 
cd ${PD}

# go file
FS=$(find -name "*.go")
echo "${FS}"

key="\".*\"" # 匹配到字符串
old="\bsucc\b"
new="succeed"
sed -n "/${old}/p" ${FS}
sed -i "/${key}/{ s/${old}/${new}/g }" ${FS}

old="to finished"
new="to finish"
sed -i "s/${old}/${new}/g" ${FS}

# 1, log...succ/succeed... ==> log...succeeded...
key="log\."
FS=$(grep -rl "${key}" --exclude-dir=.git | grep .go)
echo "${FS}"
old="\bsucceed\b"
new="succeeded"
sed -i "/${key}/{ s/${old}/${new}/g }" ${FS}

# go test file
FS=$(find -name "*_test.go")
echo "${FS}"

# 1 level --> do
old="\bdao level\b"
new="dao do"
sed -i "s/${old}/${new}/g" ${FS}

# 2 Crate --> Create
old="\Crate\b"
new="Create"
sed -i "s/${old}/${new}/g" ${FS}

# 3 handping --> Hand ping
old="\bhandping when\b"
new="Hand ping when"
sed -i "s/${old}/${new}/g" ${FS}


actions=("create" "read" "write" "update" "ping" "do" "succeed" "fail")
present=("\bsucceed\b" "\bfail\b")
past=("succeeded" "failed")

# 4 do xx --> do xxed
for action in "${actions[@]}";do
    sed -i "s/${action} ${present[0]}/${action} ${past[0]}/g" ${FS}
    sed -i "s/${action} ${present[1]}/${action} ${past[1]}/g" ${FS}
done

# 5 should be do --> should do
for action in "${actions[@]}";do
    sed -i "s/should be ${action}/should ${action}/g" ${FS}
done

old1="\bshould existed\b"
old2="\bshould not existed\b"
new1="should exist"
new2="should not exist"
sed -i "s/${old1}/${new1}/g" ${FS}
sed -i "s/${old2}/${new2}/g" ${FS}

old="\bresult is {}"
new="result should be {}"
sed -i "s/${old}/${new}/g" ${FS}

old1="\bshould be succeed\b"
old2="\bshould be succeeded\b"
sed -i "s/${old1}/${new}/g" ${FS}
sed -i "s/${old2}/${new}/g" ${FS}

# 6 whend xxed to --> when xx to
sub=".*when.*"
old="\bsucceeded to\b"
new="succeed to"
sed -i "s/\(${sub}\)\(${old}\)/\1${new}/g" ${FS}
old="\bfailed to\b"
new="fail to"
sed -i "s/\(${sub}\)\(${old}\)/\1${new}/g" ${FS}

# 7 name: xxx... ==> name: Xxx...
sub1="{name\: *\""
sub2="." # 任意字符
sed -n "/$sub1/p" ${FS}
sed -i "s/\($sub1\)\($sub2\)/\1\U\2/" ${FS}

# check
set +x
FS=$(find -name "*.go")

key1="\bsucc"
key2="\bfail"
key3="\bshould"
key4="\bexist"
key5="{name:"

keys=("${key1}" "${key2}" "${key3}" "${key4}" "${key5}")

for key in "${keys[@]}";do
    grep -rh "${key}" --exclude-dir=.git | grep -v -E '=|So|return|if|.pb.|//|#' | sed "s/\t//g"
done

echo -e "\n\n----------------"

for key in "${keys[@]}";do
    grep -r "${key}" --exclude-dir=.git | grep ".go:" | grep -v -E '=|So|return|if|.pb.|//|#' | sed "s/\t/ /g"
done