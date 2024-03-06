package main

import (
	// "database/sql"

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
	router := mux.NewRouter()
	// 设置路由
	// http.HandleFunc("/checkUsernameAndPassword", checkUsernameAndPasswordHandler)
	router.HandleFunc("/", serviceAPI.ControllerPage)
	http.Handle(
		"/web/",
		http.StripPrefix("/web/", http.FileServer(http.Dir("./static"))),
	)
	// 设置路由和处理函数
	router.HandleFunc("/registerInfo", serviceAPI.CheckUsernameAndPasswordHandler).Methods("POST")
	router.HandleFunc("/checkUsernameAndPassword", serviceAPI.CheckUsernameAndPasswordHandler).Methods("POST")
	router.HandleFunc("/items", serviceAPI.GetAllItems).Methods("GET")
	router.HandleFunc("/items", serviceAPI.CreateItem).Methods("POST")
	router.HandleFunc("/items/{id}", serviceAPI.GetItem).Methods("GET")
	router.HandleFunc("/items/{id}", serviceAPI.UpdateItem).Methods("PUT")
	router.HandleFunc("/items/{id}", serviceAPI.DeleteItem).Methods("DELETE")

	// 注册处理静态资源的路由
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	router.PathPrefix("/css/").Handler(http.StripPrefix("/css/", http.FileServer(http.Dir("static/css"))))
	router.PathPrefix("/js/").Handler(http.StripPrefix("/js/", http.FileServer(http.Dir("static/js"))))
	router.PathPrefix("/img/").Handler(http.StripPrefix("/img/", http.FileServer(http.Dir("static/img"))))
	// 启动服务器并监听指定端口
	log.Fatal(http.ListenAndServe(":8080", router))
}
