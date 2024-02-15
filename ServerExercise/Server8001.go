package main

import (
	"fmt"
	"net/http"
)

func main() {
	// 設置路由和處理函數
	http.HandleFunc("/", handler)

	// 指定監聽的端口
	port := 8001
	address := fmt.Sprintf(":%d", port)

	// 啟動Web服務
	fmt.Printf("Web服務正在監聽端口 %d...\n", port)
	err := http.ListenAndServe(address, nil)
	if err != nil {
		fmt.Println("錯誤：", err)
	}
}

// 處理請求的函數
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, " <h1>8001-歡迎使用Go語言創建的Web服務！ </h1>")
}
