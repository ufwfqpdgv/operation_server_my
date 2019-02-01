#!/bin/bash

export GOPATH=${GOPATH}:$(pwd)

env=local
server_name=samh-operation

if [[ $1 == "start" ]]
then
    nohup ./bin/${server_name} ${env} &
elif [[ $1 == "stop" ]]
then
    pkill ${server_name}
elif [[ $1 == "restart" ]]
then
    pkill ${server_name}
    nohup ./bin/${server_name} ${env} &
elif [[ $1 == "r" ]]
then
    pkill ${server_name}
    ./bin/${server_name} ${env}
else
    echo "start:开始,stop:停止,restart:重新运行"
fi

