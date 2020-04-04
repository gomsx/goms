#!/bin/bash
set -x

echo " ==> straring"

bash /bash/init_mysql.sh

bash docker-entrypoint.sh mysqld

bash /bash/init_mysql.sh

# COUNTER=0
# while [ $COUNTER -lt 50 ]
# do
#     bash /bash/init_mysql.sh
#     echo " ===> $COUNTER"
#     COUNTER+=1
# done