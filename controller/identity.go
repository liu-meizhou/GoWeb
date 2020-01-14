package controller

import (
	"strconv"
	"io"
	"fmt"
	"net/http"

	"pojo"
	"service"
	"redis"
	"tool"
	"filter"
)


//Login 验证登录
func Login(w http.ResponseWriter, r *http.Request){
	err := r.ParseForm()
	if err != nil {
		fmt.Println("解析表单数据失败!")
		fmt.Fprintln(w, "别乱发数据啊，格式都不对")
		return
	}
	if !identityCode(){
		fmt.Println("验证码输入错误!")
		fmt.Fprintln(w, "验证码输入错误")
		return
	}
	id := identityUser(r.PostForm)
	if id==-1{
		fmt.Println("账号或密码输入错误!")
		fmt.Fprintln(w, "账号或密码输入错误")
		return
	}
	//登录成功
	//修改redis
	cookie,err := r.Cookie("Visitor")
	if err!=nil{
		fmt.Fprintln(w, "你想干啥!!!")
		return
	}
	MyCookie := http.Cookie{Name: "MyCookie", Value: tool.GetMyCookie(), MaxAge: 30*60}
	http.SetCookie(w, &MyCookie)
	http.SetCookie(w, &http.Cookie{Name: "Visitor", Path: "/", MaxAge: -1})
	
	redisdb.DelDo(cookie.Value)
	redisdb.SetDo(MyCookie.Value,strconv.Itoa(id))

	io.WriteString(w, "OK")
}


//Register 注册登录
func Register(w http.ResponseWriter, r *http.Request){
	cookie,err:=r.Cookie("Visitor")
	if cookie!=nil && err==nil{
		register(w, r)
		return
	}
	fn := filter.SafeHandler(AddAUser)
	fn(w,r)
}

func register(w http.ResponseWriter, r *http.Request){
	if !identityCode(){
		fmt.Println("验证码输入错误!")
		fmt.Fprintln(w, "验证码输入错误")
		return
	}
	id:=AddUser(r)
	if _,err:=strconv.Atoi(id);err!=nil{
		io.WriteString(w, id)
		return
	}
	//修改redis
	cookie,err := r.Cookie("Visitor")
	if err!=nil{
		fmt.Println("Cookie已过期或不可以，请重新登录")
		fmt.Fprintln(w, "Cookie已过期或不可以，请重新登录")
		return
	}
	MyCookie := http.Cookie{Name: "MyCookie", Value: cookie.Value, MaxAge: 30*60}
	http.SetCookie(w, &MyCookie)
	http.SetCookie(w, &http.Cookie{Name: "Visitor", Path: "/", MaxAge: -1})
	redisdb.UpDo(cookie.Value,id)

	io.WriteString(w, id)
}

//验证验证码
func identityCode()(istrue bool){
	return true
}

//验证账号密码
func identityUser(umap map[string][]string) (id int){
	userService := service.GetUserService()
	index := umap["index"][0]
	var users []pojo.User
	if index=="1" {
		users = userService.GetUserByName(umap["username"][0])
	}else if index=="2"{
		users = userService.GetUserByEmail(umap["email"][0])
	}else{
		users = userService.GetUserByPhone(umap["phone"][0])
	}
	for user := range users{
		if users[user].Pwd==umap["password"][0]{
			return users[user].ID
		}
	}
	return -1
}





