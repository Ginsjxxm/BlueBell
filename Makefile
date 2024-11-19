.PHONY: all build run gotool clean help

BINARY=Wang

all: gotool build

# 针对当前 Windows 系统编译
build:
	@go build -o $(BINARY).exe

# 运行程序
run:
	@go run ./

# 格式化和检查代码
gotool:
	@go fmt ./...
	@go vet ./...

# 清理生成的二进制文件
clean:
	@if exist $(BINARY).exe del /f $(BINARY).exe

# 帮助信息
help:
	@echo "make - Format Go code and build the binary file"
	@echo "make build - Compile Go code and generate a Windows binary file"
	@echo "make run - Run the Go code directly"
	@echo "make clean - Remove the generated binary file"
	@echo "make gotool - Use 'fmt' and 'vet' to check the code"

