package main

import (
	// "database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"serverDemo/serviceAPI"
	"serverDemo/sql"

	// "Users/starpensive/GO-Language-1/serverDemo/SQL"

	_ "github.com/go-sql-driver/mysql" //go get -u github.com/go-sql-driver/mysql
	"github.com/gorilla/mux"           //go get -u github.com/gorilla/mux
)

func main() {
	sql.SQLinit()
	serviceAPI.SayHello()
	router := mux.NewRouter()
	// 设置路由
	// http.HandleFunc("/checkUsernameAndPassword", checkUsernameAndPasswordHandler)
	router.HandleFunc("/", controllerPage)
	// 设置路由和处理函数
	router.HandleFunc("/registerInfo", checkUsernameAndPasswordHandler).Methods("POST")
	router.HandleFunc("/checkUsernameAndPassword", checkUsernameAndPasswordHandler).Methods("POST")
	router.HandleFunc("/items", getAllItems).Methods("GET")
	router.HandleFunc("/items", createItem).Methods("POST")
	router.HandleFunc("/items/{id}", getItem).Methods("GET")
	router.HandleFunc("/items/{id}", updateItem).Methods("PUT")
	router.HandleFunc("/items/{id}", deleteItem).Methods("DELETE")

	// 注册处理静态资源的路由
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	router.PathPrefix("/css/").Handler(http.StripPrefix("/css/", http.FileServer(http.Dir("static/css"))))
	router.PathPrefix("/js/").Handler(http.StripPrefix("/js/", http.FileServer(http.Dir("static/js"))))
	router.PathPrefix("/img/").Handler(http.StripPrefix("/img/", http.FileServer(http.Dir("static/img"))))
	// 启动服务器并监听指定端口
	log.Fatal(http.ListenAndServe(":8080", router))
}

type LoginResponse struct {
	Exists   bool   `json:"exists"`
	Message  string `json:"message"`
	Redirect string `json:"redirect"`
}

func controllerPage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/index.html")
}
func getAllItems(w http.ResponseWriter, r *http.Request) {
	// 处理获取所有项的逻辑
}

func getItem(w http.ResponseWriter, r *http.Request) {
	// 处理获取单个项的逻辑
}

func createItem(w http.ResponseWriter, r *http.Request) {
	// 处理创建项的逻辑
}

func updateItem(w http.ResponseWriter, r *http.Request) {
	// 处理更新项的逻辑
}

func deleteItem(w http.ResponseWriter, r *http.Request) {
	// 处理删除项的逻辑
}

func checkUsernameAndPasswordHandler(w http.ResponseWriter, r *http.Request) {
	// 解析表单数据
	var formData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	err := json.NewDecoder(r.Body).Decode(&formData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println("username=" + formData.Username)
	fmt.Println("password=" + formData.Password)

	// 做一些处理逻辑，比如检查用户名是否已存在
	exists := false
	message := ""
	redirect := ""

	if exists {
		// 用户名已存在的处理逻辑
		message = "Username already exists: " + formData.Username
	} else {
		// 用户名可用，继续注册流程
		message = "Registration successful for username: " + formData.Username
		redirect = "/success" // 跳转到成功页面的URL
	}

	// 构造登录响应
	loginResponse := LoginResponse{
		Exists:   exists,
		Message:  message,
		Redirect: redirect,
	}

	// 将响应转换为JSON字符串
	jsonResponse, err := json.Marshal(loginResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// 设置响应的内容类型为JSON
	w.Header().Set("Content-Type", "application/json")

	// 发送响应
	w.Write(jsonResponse)
}
