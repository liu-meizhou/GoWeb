# GoWeb 此代码是直接在docker环境下拉取，dockerfile创建成image 运行image发布的项目
# 电脑需要有 docker-images: alpine golang:alpine ; git 数据库postgres-alpine

# 1.在docker 上更新项目文件 `sudo git pull`
# 2.进入项目文件(Dockerfile文件目录下) `cd GoWeb` 
#  name=GoWeb
# 3.用Dockerfile创建image 'docker build -t ${name}`
# 4.运行image `docker run -d -p 8888:8888 ${imageID}`
# 5.
