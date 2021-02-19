#!/bin/bash
set -x

# cd ..
cd ..
pwd

# 1, log...succ/succeed... ==> log...succeeded...
key="log\."
fs=$(grep -rl "${key}" --exclude-dir=.git | grep .go)
echo "${fs}"
old1="\bsucc\b"
old2="\bsucceed\b"
new="succeeded"
sed -i "/${key}/{s/${old1}/${new}/g;s/${old2}/${new}/g}" ${fs}

# 2, when...succ/succeeded to... ==> when...succeed to... 
fs=$(find -name "*_test.go")
echo "${fs}"

sub=".*when.*"
old1="\bsucc to\b"
old2="\bsucceeded to\b"
new="succeed to"
sed -i "s/\(${sub}\)\(${old1}\)/\1${new}/g" ${fs}
sed -i "s/\(${sub}\)\(${old2}\)/\1${new}/g" ${fs}

old="\bfailed to\b"
new="fail to"
sed -i "s/\(${sub}\)\(${old}\)/\1${new}/g" ${fs}

# 3, ...
old="to finished"
new="to finish"
sed -i "s/${old}/${new}/g" ${fs}

old1="\bshould be succ\b"
old2="\bshould be succeeded\b"
new="should be succeed"
sed -i "s/${old1}/${new}/g" ${fs}
sed -i "s/${old2}/${new}/g" ${fs}

old1="\bshould existed\b"
old2="\bshould not existed\b"
new1="should exist"
new2="should not exist"
sed -i "s/${old1}/${new1}/g" ${fs}
sed -i "s/${old2}/${new2}/g" ${fs}

old1="\bshould be write succ\b"
old2="\bshould be read succ\b"
old3="\bshould be succeed\b"
new1="should write succeeded"
new2="should read succeeded"
new3="should succeed"
sed -i "s/${old1}/${new1}/g" ${fs}
sed -i "s/${old2}/${new2}/g" ${fs}
sed -i "s/${old3}/${new3}/g" ${fs}

old1="\bread succ\b"
old2="\bupdate succ\b"
old3="\bping succ\b"
old4="\bdao level succ\b"
new1="read succeeded"
new2="update succeeded"
new3="ping succeeded"
new4="dao do succeed"
sed -i "s/${old1}/${new1}/g" ${fs}
sed -i "s/${old2}/${new2}/g" ${fs}
sed -i "s/${old3}/${new3}/g" ${fs}
sed -i "s/${old4}/${new4}/g" ${fs}

old="\bhandping when\b"
new="Hand ping when"
sed -i "s/${old}/${new}/g" ${fs}

# 4 name: xxx... ==> name: Xxx...
sub1="{name\: *\""
sub2="."
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