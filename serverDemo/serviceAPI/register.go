package serviceAPI

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type LoginResponse struct {
	Exists   bool   `json:"exists"`
	Message  string `json:"message"`
	Redirect string `json:"redirect"`
}

func CheckUsernameAndPasswordHandler(w http.ResponseWriter, r *http.Request) {
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
