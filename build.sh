#!/bin/bash

export GOOS=linux
export GOARCH=amd64
export CGO_ENABLED=0
export GOPATH=${GOPATH}:$(pwd)

env=local
server_name=samh-operation
server_path=src/server/main/main.go
server_name_test=${server_name}"-test"
server_path_test=src/process/main/main.go
pack_path=bin/${server_name}.tar.gz

if [[ $1 == "b" ]]
then
    go build -o bin/${server_name} ${server_path}
    go build -o bin/${server_name_test} ${server_path_test}
elif [[ $1 == "t" ]]
then
    ./bin/$server_name_test ${env}
elif [[ $1 == "r" ]]
then
    ./bin/${server_name} ${env}
elif [[ $1 == "p" ]]
then
    mkdir -p ${server_name}/bin
    cp bin/${server_name} bin/${server_name_test} ${server_name}/bin
    cp -r config run.sh ${server_name}
    tar -zcvf ${pack_path} ${server_name}
    rm -rf ${server_name}
elif [[ $1 == "s" ]]
then
    scp ${pack_path} root@47.99.39.114:/home/lingyuanwei/xndm_work/server/pack
    scp run2.sh root@47.99.39.114:/home/lingyuanwei/xndm_work/server
elif [[ $1 == "st" ]]
then
    scp ${pack_path} root@118.31.21.181:/home/lingyuanwei/xndm_work/server/pack
    scp run2.sh root@118.31.21.181:/home/lingyuanwei/xndm_work/server
elif [[ $1 == "pt" ]]
then
    mkdir -p ${server_name}/bin
    cp bin/${server_name} bin/${server_name_test} ${server_name}/bin
    cp -r config ${server_name}
    now_time=$(date "+%Y-%m-%d_%H:%M:%S")
    tar -zcvf bin/${server_name}"."${now_time}.tar.gz ${server_name}
    rm -rf ${server_name}
else
    echo "b:构建,t:本地测试,r:本地跑服务,p:打包成部署用的压缩包到bin里"
fi

