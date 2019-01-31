## 项目
这个项目提供了飒漫画的运营活动服务

## 项目说明
- build.sh为构建文件，./build.sh b:构建,t:本地测试,r:本地跑服务,p:打包成部署用的压缩包到bin里
- bin为可执行文件及部署包存放
- config为服务及log配置
- doc文件夹为db脚本等
- src为源码，内process为主要逻辑处理，server为http服务搭建
- glide.yaml为glide管理的相关依赖，如到新环境可直接用glide下载全部依赖
