# Git

## git 跟踪分支和变基

将本地的 main 分支改成跟踪 upstream/main：
```
git fetch upstream
git branch --set-upstream-to=upstream/main main
```
一旦设置了跟踪分支，一定要变基（rebase）你的 main 分支，使它与上游仓库的任何新变化保持一致：
```
git remote update
git checkout main
git rebase
```

## git commit test\chore

test 包含 ut\bt\it 和 tests 目录下
chore 包含 build(makefile)\cicd(workflow)\tools\git-hook\...

## git reset\git checkout -- .

git reset 从 history 区重置 cached 区
git checkout -- . 从 cached 区重置 work 区

## git 批量处理改动的文件

1. 列出改动的文件
git st | awk '{ print $2 }'

2. grep 关注的文件
git st | awk '{ print $2 }' | grep _test.go

## git 分支管理

git br | grep m.bk[2-9] | xargs git br -D

## git co 某个分支的某个目录

git co dev -- docs

## 远程库管理

1, 删除分支
git push origin :m.bk1
2, 伤处tag
git push origin :v1.10.0 --tags

## stash

1, stash pop 冲突解决后要 stash drop,没有 stash --continue

## 查看连个分支的 commit 差异

git log -s main...dev

## 查看工作区文件的改动

git diff --name-status | grep -v "^D" | awk '{ print $2; }
