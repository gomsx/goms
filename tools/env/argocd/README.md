# argocd

## deploy argocd

```
bash deploy-argocd.sh
```

## install argocd cli

```
bash install-argocd-cli.sh
```

## login cmd

```
argocd login localhost:31141 --username admin --password $(kubectl get pods -n argocd -l app.kubernetes.io/name=argocd-server -o name | cut -d'/' -f 2) --insecure
```

## login ui

UI: http://120.79.1.69:31140
Username: admin
Password: ...

```
# get admin password
kubectl get pods -n argocd -l app.kubernetes.io/name=argocd-server -o name | cut -d'/' -f 2
```

## add ssh

```
argocd cert add-ssh --batch --from ssh_known_hosts_file
```

ssh_known_hosts_file 格式参考 ssh-keyscan www.github.com 执行结果
```
$ ssh-keyscan www.github.com
www.github.com ssh-rsa AAAAB3NzaC1yc2EAAAAB....
```

## add repo

```
argocd repo add git@github.com:fuwensun/goms.git --ssh-private-key-path ~/.ssh/id_rsa --name fuwensun-goms-repo

argocd repo list
```

## add app

```
# create namespace
kubectl create namespace fuwensun-goms-dev

# create app
argocd app create app-fuwensun-goms-dev   \
    --repo git@github.com:fuwensun/goms.git   \
    --path eK8s/app/overlays/dev    \
    --dest-server https://kubernetes.default.svc    \
    --dest-namespace fuwensun-goms-dev    \
    --revision dev    \
    --sync-policy auto    \
    --upsert

# list app
argocd app list
argocd app set app-fuwensun-goms-dev --sync-policy automated
```

## argocd manage ssh

Managing SSH Known Hosts using the CLI
```
# list all configured SSH known host entries 
argocd cert list --cert-type ssh

# adding all available SSH public host keys for a server to ArgoCD, as collected by ssh-keyscan
ssh-keyscan server.example.com | argocd cert add-ssh --batch 

# importing an existing known_hosts file to ArgoCD
argocd cert add-ssh --batch --from /etc/ssh/ssh_known_hosts

```

Managing SSH known hosts data using declarative setup
```
argocd repo add git@github.com:fuwensun/goms.git --ssh-private-key-path ~/.ssh/id_rsa --insecure-ignore-host-key
```

## argocd manage repo

```
argocd repo add git@github.com:fuwensun/goms.git --ssh-private-key-path ~/.ssh/id_rsa.alz --name goms
argocd repo add git@github.com:fuwensun/goms.git --ssh-private-key-path ~/.ssh/id_rsa.alz
argocd repo list
argocd repo rm git@github.com:fuwensun/goms.git
argocd repo goms
```

## argocd manage app

```
# create namespace
kubectl create namespace goms-dev

# create app
argocd app create app-goms-dev   \
    --repo https://github.com/fuwensun/goms.git  \
    --path eK8s/app/overlays/dev  \
    --dest-server https://kubernetes.default.svc  \
    --dest-namespace goms-dev  \
    --revision HEAD 

# app
argocd app get test-goms
argocd app sync test-goms
argocd app set test-goms --sync-policy automated
argocd app list
argocd app list test-goms
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
> https://github.com/argoproj/argo-cd
