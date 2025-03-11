#!/bin/bash
# 生成 Swagger 文档的脚本
# 作者: [你的名字]
# 日期: [日期]
# 描述: 这个脚本用于生成 Swagger 文档。

set -ex
# 项目目录
PROJECT_DIR=$(cd "`dirname $0`/../../" && pwd)
# 进入项目目录
cd $PROJECT_DIR
# 当前脚本所在的目录
WEB_DIR=./bin/web
# 存放 Swagger 文档的目录
SWAG_DIR=$PROJECT_DIR/app/v1/swagger
# 生成 Swagger 文档
swag init -g $WEB_DIR/main.go -o $SWAG_DIR/docs