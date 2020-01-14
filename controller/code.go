package controller

import (
	"io"
	"net/http"
	// "encoding/json"
	
	"redis"
	"tool"
)


//Visitor 游客获取验证码并且设置MyCookie
func Visitor(w http.ResponseWriter, r *http.Request) {

	// haveCookie,err := r.Cookie("MyCookie")
	// if err!=nil{
	// 	return
	// }
	code := getcode()
	
	// if haveCookie!=nil{
	// 	redisdb.UpDo(haveCookie.Name,code)
	// 	return
	// }
	mycookie := tool.GetMyCookie()
	redisdb.SetDo(mycookie,code)
	cookie := &http.Cookie{
		Name : "Visitor",
		Value : mycookie,
		MaxAge : 30*60,
	}
	http.SetCookie(w, cookie)
	//返回json
	io.WriteString(w, code)
}

func getcode()(code string){
	code = "code"
	return code
}