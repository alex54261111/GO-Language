package serviceAPI

import (
	"fmt"
	"net/http"
)

func SayHello1() {
	fmt.Println("Hello")
}

func ControllerPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/index.html")
}
func GetAllItems(w http.ResponseWriter, r *http.Request) {
	// 处理获取所有项的逻辑
}

func GetItem(w http.ResponseWriter, r *http.Request) {
	// 处理获取单个项的逻辑
}

func CreateItem(w http.ResponseWriter, r *http.Request) {
	// 处理创建项的逻辑
}

func UpdateItem(w http.ResponseWriter, r *http.Request) {
	// 处理更新项的逻辑
}

func DeleteItem(w http.ResponseWriter, r *http.Request) {
	// 处理删除项的逻辑
}
