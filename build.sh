#!/bin/bash

# 定义输出的二进制文件名
OUTPUT_FILE="Mbot"

# 设置编译目标为 Linux 系统，架构为常见的 amd64（x86_64），
# 你可以根据实际需求修改 GOARCH 的值，比如 arm64 等
export GOOS=linux
export GOARCH=amd64
# 禁用 CGO，避免依赖外部 C 库，提升在不同 Linux 系统上的兼容性
export CGO_ENABLED=0

# 进入项目根目录（如果当前脚本不在项目根目录下，可按需调整）
cd "$(dirname "$0")"

# 执行编译命令
go build -o $OUTPUT_FILE main.go
zip -r Mbot.zip Mbot
# 检查编译是否成功
if [ $? -eq 0 ]; then
    echo "编译成功！生成的文件名为: $OUTPUT_FILE zip"
else
    echo "编译失败，请检查代码或 Go 环境配置。"
fi