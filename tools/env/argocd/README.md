# argocd

## deploy argocd

```
deploy-argocd.sh
```

## install argocd cli

```
install-argocd-cli.sh
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
# get admin password
kubectl get pods -n argocd -l app.kubernetes.io/name=argocd-server -o name | cut -d'/' -f 2

# login cmd 1
argocd login localhost:31141 --username admin --password argocd-server-58665666dc-82xvm --insecure

# login cmd 2
argocd login localhost:31141 --username admin --password $(kubectl get pods -n argocd -l app.kubernetes.io/name=argocd-server -o name | cut -d'/' -f 2) --insecure

# login ui https://120.79.33.44:31140/

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
    --repo git@github.com:fuwensun/goms.git  \
    --path eK8s/app/overlays/test  \
    --dest-server https://kubernetes.default.svc  \
    --dest-namespace test-goms  \
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

