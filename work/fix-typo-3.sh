#!/bin/bash
set -x
set -e

# cd ..
cd ..
pwd

# pre="\"\.*"
# suffi="\.*\""

# go file
fs=$(find -name "*.go")
echo "${fs}"

key="\".*\"" # 匹配到字符串
old="\bsucc\b"
new="succeed"
sed -n "/${old}/p" ${fs}
sed -i "/${key}/{ s/${old}/${new}/g }" ${fs}

old="to finished"
new="to finish"
sed -i "s/${old}/${new}/g" ${fs}

# 1, log...succ/succeed... ==> log...succeeded...
key="log\."
fs=$(grep -rl "${key}" --exclude-dir=.git | grep .go)
echo "${fs}"
old="\bsucceed\b"
new="succeeded"
sed -i "/${key}/{ s/${old}/${new}/g }" ${fs}

# go test file
fs=$(find -name "*_test.go")
echo "${fs}"

# 1 level --> do
old="\bdao level\b"
new="dao do"
sed -i "s/${old}/${new}/g" ${fs}

# 2 Crate --> Create
old="\Crate\b"
new="Create"
sed -i "s/${old}/${new}/g" ${fs}

# 3 handping --> Hand ping
old="\bhandping when\b"
new="Hand ping when"
sed -i "s/${old}/${new}/g" ${fs}


actions=("create" "read" "write" "update" "ping" "do" "succeed" "fail")
present=("\bsucceed\b" "\bfail\b")
past=("succeeded" "failed")

# 4 do xx --> do xxed
for action in "${actions[@]}";do
    sed -i "s/${action} ${present[0]}/${action} ${past[0]}/g" ${fs}
    sed -i "s/${action} ${present[1]}/${action} ${past[1]}/g" ${fs}
done

# 5 should be do --> should do
for action in "${actions[@]}";do
    sed -i "s/should be ${action}/should ${action}/g" ${fs}
done

old1="\bshould existed\b"
old2="\bshould not existed\b"
new1="should exist"
new2="should not exist"
sed -i "s/${old1}/${new1}/g" ${fs}
sed -i "s/${old2}/${new2}/g" ${fs}

old1="\bshould be succeed\b"
old2="\bshould be succeeded\b"
sed -i "s/${old1}/${new}/g" ${fs}
sed -i "s/${old2}/${new}/g" ${fs}

# 6 whend xxed to --> when xx to
sub=".*when.*"
old="\bsucceeded to\b"
new="succeed to"
sed -i "s/\(${sub}\)\(${old}\)/\1${new}/g" ${fs}
old="\bfailed to\b"
new="fail to"
sed -i "s/\(${sub}\)\(${old}\)/\1${new}/g" ${fs}

# 7 name: xxx... ==> name: Xxx...
sub1="{name\: *\""
sub2="." # 任意字符
sed -n "/$sub1/p" ${fs}
sed -i "s/\($sub1\)\($sub2\)/\1\U\2/" ${fs}

# check
set +x
fs=$(find -name "*.go")

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