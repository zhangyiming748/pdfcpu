# 获取目标操作系统架构
GOOS := $(shell go env GOOS)
GOARCH := $(shell go env GOARCH)

# 根据操作系统设置可执行文件名
ifeq ($(GOOS),windows)
    TARGET = pdfcpu.exe
else
    TARGET = pdfcpu
endif

# 默认目标
all: build

# 构建可执行文件
build:
	go build -o $(TARGET) main.go
	@echo "构建完成: $(TARGET) for $(GOOS)/$(GOARCH)"

# 清理生成的文件
clean:
ifeq ($(GOOS),windows)
	del pdfcpu.exe
else
	rm -f pdfcpu
endif
	rm -f pdfcpu.log

# 安装依赖（如果需要的话）
install:

# 运行程序（示例）
run:
	go run main.go

# 更新和管理依赖
dep:
	go get -u ./...
	go mod tidy
	go mod vendor

.PHONY: all build clean install run dep