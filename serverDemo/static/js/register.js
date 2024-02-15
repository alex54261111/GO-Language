$(document).ready(function() {
    $("#registerFrome").submit(function(event) {
           event.preventDefault(); // 阻止表单的默认提交行为
           // 获取表单数据
           var formData = {
             username: $("#username").val(),
             password: $("#password").val()
           };
           // 发送AJAX请求
           $.ajax({
               type: "POST",
               url: "/registerInfo",
               data: JSON.stringify(formData), // 将数据转换为JSON字符串
               contentType: "application/json", // 设置请求的内容类型为JSON
               success: function(loginResponse) {
                   if (loginResponse.exists) {
                       // 用户名已存在的处理逻辑
                       console.log("Username already exists: " + formData.username);
                       $("#errorMessage").text(loginResponse.Message);
                   } else {
                       // 用户名可用，继续注册流程
                       console.log("Registration successful for username: " + formData.username);
                       // 提交表单
                       $("form").unbind("submit").submit();
                       $("#errorMessage").text(loginResponse.Message);
                       window.location.href = loginResponse.redirect; // 跳轉到指定的頁面
                   }
               },
               error: function() {
                   console.log("Error occurred during username check.");
               }
           });
       });
    })