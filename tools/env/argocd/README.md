# argocd

## deploy argocd

```
deploy-argocd.sh
```

## install argocd cli

```
install-argocd-cli.sh
```

## argocd manage app

```
# create namespace
kubectl create namespace test-goms

# create app
argocd app create test-goms   \
    --repo https://github.com/fuwensun/goms.git  \
    --path eK8s/app/overlays/test  \
    --dest-server https://kubernetes.default.svc  \
    --dest-namespace test-goms  \
    --revision HEAD 

# ok
argocd app create test-goms   \
    --repo https://github.com/fuwensun/goms.git  \
    --path eK8s/app/overlays/test  \
    --dest-server https://kubernetes.default.svc  \
    --dest-namespace test-goms  \
    --revision master \
    --server-crt /root/.ssh/id_rsa

# get admin password
kubectl get pods -n argocd -l app.kubernetes.io/name=argocd-server -o name | cut -d'/' -f 2

# login ui https://120.79.33.44:31140/
# login cmd
argocd login localhost:31141 --username admin --password argocd-server-58665666dc-82xvm --insecure

# app
argocd app get test-goms
argocd app sync test-goms
argocd app set test-goms --sync-policy automated
```

## argocd cmd

```
argocd --help

account     Manage account settings
app         Manage applications
cert        Manage repository certificates and SSH known hosts entries
cluster     Manage cluster credentials
completion  output shell completion code for the specified shell (bash or zsh)
context     Switch between contexts
help        Help about any command
login       Log in to Argo CD
logout      Log out from Argo CD
proj        Manage projects
relogin     Refresh an expired authenticate token
repo        Manage repository connection parameters
repocreds   Manage repository connection parameters
version     Print version information
```

