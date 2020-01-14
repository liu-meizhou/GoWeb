package controller

import (
	"io"
	"strconv"
	"fmt"
	"net/http"
	"encoding/json"

	"pojo"
	"redis"
	"service"
)

var userService service.UserService
var messageService service.MessageService

func init(){
	userService=service.GetUserService()
	messageService=service.GetMessageService()
}


//Index 注册登录
func Index(w http.ResponseWriter, r *http.Request){
	http.Redirect(w, r, "/Logined/index.html", http.StatusPermanentRedirect)
}

//Logout 注销登录
func Logout(w http.ResponseWriter, r *http.Request){
	cookie,_ := r.Cookie("MyCookie")
	//销毁redis
	redisdb.DelDo(cookie.Name)

	//跳转页面
	http.Redirect(w, r, "/Login.html", http.StatusPermanentRedirect)

}

//GetUserMgr 获取所有用户管理列表
func GetUserMgr(w http.ResponseWriter, r *http.Request){
	users := userService.GetAllUsers()
	b , err := json. Marshal ( users ) 
    if err != nil { 
        fmt. Println ( "error:" , err ) 
	} 
	fmt.Fprintln(w, string(b))
}

//AddUser 添加用户
func AddUser(r *http.Request)(id string){
	err := r.ParseForm()
	if err != nil {
		return "数据格式不对"
	}
	umap := r.PostForm

	return strconv.Itoa(service.GetUserService().AddUser(pojo.NewUser(0,umap["username"][0],umap["email"][0],umap["phone"][0],umap["password"][0]))) 
}

//AddAUser 添加用户(用于用户管理)
func AddAUser(w http.ResponseWriter, r *http.Request){
	id:=AddUser(r)
	io.WriteString(w, id)
}

//UpdateByID 修改用户
func UpdateByID(w http.ResponseWriter, r *http.Request){
	err := r.ParseForm()
	if err != nil {
		fmt.Println("解析表单数据失败!")
		fmt.Fprintln(w, "别乱发数据啊，格式都不对")
	}
	umap := r.PostForm
	id,_ := strconv.Atoi(umap["ID"][0])
	userService.UpdateUser(pojo.NewUser(id, umap["username"][0],umap["email"][0],umap["phone"][0],umap["password"][0]))
}

//DeleteByID 修改用户
func DeleteByID(w http.ResponseWriter, r *http.Request){
	err := r.ParseForm()
	if err != nil {
		fmt.Println("解析表单数据失败!")
		fmt.Fprintln(w, "别乱发数据啊，格式都不对")
	}
	umap := r.PostForm
	id,_ := strconv.Atoi(umap["ID"][0])
	userService.DeleteUserByID(id)
}