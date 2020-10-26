# golang-study

#### 包管理
export GO111MODULE=on
export GOPROXY=https://goproxy.io
golang 1.12的朋友，确保实验目录不在 GOPATH 中。
go mod init [packagename]
<!-- 拉取必须模块，移除不用的模块 -->
go mod tidy
<!-- 只下载依赖包 -->
go mod download
<!-- 添加新依赖包 -->
方法一：
直接修改 go.mod 文件，然后执行 go mod download
方法二：
使用 go get packagename@v1.2.3，会自动更新 go.mod 文件的
方法三：
go run、go build 也会自动下载依赖