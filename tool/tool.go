package tool

import (
	"fmt"
	"strconv"
	"net/http"

	"github.com/satori/go.uuid"

	"redis"
)

//GetLoginedID 获取登录的ID
func GetLoginedID(r *http.Request)(id int,err error){
	cookie,err := r.Cookie("MyCookie")
	if(err!=nil){
		fmt.Println("无MyCookie");
		return 0,err
	}
	val,err1 := redisdb.GetDo(cookie.Value)
	id,err2 := strconv.Atoi(val)
	if err1!=nil{
		return 0,err1
	}else if err2!=nil{
		return 0,err2
	}
	return id,nil
}

//GetMyCookie 获取MyCookie
func GetMyCookie()(mycookie string){
	return uuid.Must(uuid.NewV4()).String()
}