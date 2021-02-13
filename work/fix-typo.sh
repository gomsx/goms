#!/bin/bash
set -x

cd ..
pwd

# log...succ... ==> log...succeed...
key="log\."
old="\bsucc\b"
new="succeed"
fs=$(grep -rl "${key}" --exclude-dir=.git | grep ".go")
echo "$fs"
for f in $fs; do
	sed -i "/${key}/{s/${old}/${new}/g}" "${f}"
done

# Convey(...is succ... ==> Convey(...should succeed...
key="Convey("
old1="\bis succ\b"
new1="should succeed"
old2="\bis exist\b"
new2="should existed"
old3="\bis not exist\b"
new3="should not existed"
fs=$(grep -rl "${key}" --exclude-dir=.git | grep ".go")
echo "$fs"
for f in $fs; do
	sed -i "/${key}/{s/${old1}/${new1}/g}" "${f}"
	sed -i "/${key}/{s/${old2}/${new2}/g}" "${f}"
	sed -i "/${key}/{s/${old3}/${new3}/g}" "${f}"
	sed -i "s/\bthe the\b/the/g" "${f}"
	sed -i "s/\bCrate\b/Create/g" "${f}"
	sed -i "s/\bthe result is\b/the result should be/g" "${f}"
done

# 单词边界 https://www.cnblogs.com/gaara0305/p/10027343.html
