export install_path = /opt/soft
export go_version = go1.17.5.linux-amd64
export GOROOT ?= $(install_path)/$(go_version)
export GOPATH ?= $(install_path)/gopath
export GOBIN ?= $(GOPATH)/bin
export PATH := $(GOROOT)/bin:$(GOBIN):$(PATH)
export GOPROXY := https://goproxy.io,direct

BIN_NAME = a.out

all: clean $(BIN_NAME)

# Just for local testing
check:$(BIN_NAME)
	./$(BIN_NAME) ./input/auntdemo.csv ./input/orderdemo.csv ./result.csv ./useTime.txt

# 如果未使用 cgo, 可以去掉这里的 -a 选项以加快编译速度
$(BIN_NAME):
	go build -a -o $(BIN_NAME) main.go solution.go

clean:
	rm -rf $(BIN_NAME)
