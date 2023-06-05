# 使用 golang 1.20 作为基础镜像
FROM golang:1.20

# 将工作目录设置为 /app
WORKDIR /app

# 将本地文件复制到容器中
COPY . .

# 构建 Go 应用程序
RUN go build -o main .

# 暴露端口 8080
EXPOSE 8080

# 运行应用程序
CMD ["./main"]
