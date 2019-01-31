#!/bin/bash

server_name=samh-operation

if [[ $1 == "b" ]]
then
    rm -rf ${server_name}
    tar -zxvf pack/${server_name}.tar.gz
elif [[ $1 == "r" ]]
then
    cd ${server_name}
    ./run.sh restart
else
    echo "b:重新部署,r:重新运行"
fi

