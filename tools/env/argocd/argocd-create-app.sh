#!/bin/bash

set -x

argocd app create guestbook   \
    --repo https://github.com/aivuca/argocd-example-apps.git  \
    --path guestbook  \
    --dest-server https://kubernetes.default.svc  \
    --dest-namespace default  \
    --revision HEAD 

argocd app create test-goms   \
    --repo https://github.com/fuwensun/goms.git  \
    --path eK8s/deploy/overlays/test  \
    --dest-server https://kubernetes.default.svc  \
    --dest-namespace test-goms  \
    --revision master \
    --server-crt /root/.ssh/id_rsa


argocd app create test-goms   \
    --repo https://github.com/fuwensun/goms.git  \
    --path eK8s/app/overlays/test  \
    --dest-server https://kubernetes.default.svc  \
    --dest-namespace test-goms  \
    --revision HEAD 

argocd app get guestbook
argocd app sync guestbook
