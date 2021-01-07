
## set pre-commit hook for git
```
cd .git/hooks
ln -s ../../tools/hooks/pre-commit
```

## manual examination check
```
check.sh
check.sh all

tools/hooks/check-doc.sh "$(find . -name "*makefile")"
tools/hooks/check-code-go.sh "$(find . -name "*.go")"
```