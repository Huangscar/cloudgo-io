#
# BUILD 阶段
# 
FROM golang:1.10 AS build

# 设置我们应用程序的工作目录
WORKDIR /go/src/github.com/Huangscar/cloudgo-io



# 添加所有需要编译的应用代码
ADD . .

RUN  go get ./

# 编译一个静态的go应用（在二进制构建中包含C语言依赖库）
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .

# 设置我们应用程序的启动命令
CMD ["./cloudgo-io"]