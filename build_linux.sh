#!/bin/bash

# 设置目标操作系统和架构
export GOOS=linux
export GOARCH=amd64

# 二进制文件名
BINARY_NAME="lh-tool"
# 压缩包文件名
ARCHIVE_NAME="lh-tool-linux-amd64.tar.gz"

echo "正在构建 Linux 版本 ($BINARY_NAME)..."
go build -ldflags "-s -w" -o $BINARY_NAME .

if [ $? -eq 0 ]; then
    echo "构建成功: $BINARY_NAME"
    
    echo "正在打包..."
    tar -czvf $ARCHIVE_NAME $BINARY_NAME
    
    if [ $? -eq 0 ]; then
        echo "打包成功: $ARCHIVE_NAME"
        # 打包后删除原始二进制文件，保持目录整洁
        rm $BINARY_NAME
    else
        echo "打包失败"
        exit 1
    fi
else
    echo "构建失败"
    exit 1
fi
