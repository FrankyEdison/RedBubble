# Makefile不能在Windows下使用，mac和linux都行
# 声明这些目标
.PHONY: all build run gotool clean help

# 定义项目编译后的名字
BINARY="RedBubble"

# 在终端只输入make，就会执行这条命令
all: gotool build

# 在linux编译出可执行文件
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o ./bin/${BINARY}

# 运行，命令前面有@则执行时不在终端打印该命令
run:
	@go run ./main.go

gotool:
    # 格式化代码
	go fmt ./
    # 检查
	go vet ./

# 如果当前目录下已有编译出的可执行文件，则删除它
clean:
	@if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

# 打印提示
help:
	@echo "make - 格式化 Go 代码, 并编译生成二进制文件"
	@echo "make build - 编译 Go 代码, 生成二进制文件"
	@echo "make run - 直接运行 Go 代码"
	@echo "make clean - 移除二进制文件和 vim swap files"
	@echo "make gotool - 运行 Go 工具 'fmt' and 'vet'"
