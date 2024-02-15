package main

import (
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// 创建一个信号接收通道
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	// 启动服务器代码

	// 等待接收退出信号
	<-sigCh

	// 执行退出操作，例如关闭数据库连接、保存数据等

	// 退出服务器
	os.Exit(0)
}
