#!/bin/bash

# 清理旧的编译结果
rm -f sgamer*

# 编译 Windows 64 位版本
export GOOS=windows
export GOARCH=amd64
go build -o sgamer.exe

# 编译 macOS 64 位版本
export GOOS=darwin
export GOARCH=arm64
go build -o sgamer_darwin_amd64

# 编译 Linux 64 位版本
export GOOS=linux
export GOARCH=amd64
go build -o sgamer_linux_amd64