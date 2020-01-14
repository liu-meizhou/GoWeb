package controller

import (
	"encoding/json"
	"strconv"
	"fmt"
	"net/http"

	"tool"
)


//GetNoReadCount 获取未读条数
func GetNoReadCount(w http.ResponseWriter, r *http.Request){
	v,err := tool.GetLoginedID(r)
	if err!=nil{
		fmt.Println("执行redis错误")
		return
	}
	fmt.Fprint(w, messageService.GetNoReadCountMessageByReceiveID(v)) 
}

//GetEveryNoRead 获取未读条数
func GetEveryNoRead(w http.ResponseWriter, r *http.Request){
	v,err := tool.GetLoginedID(r)
	if err!=nil{
		fmt.Println("执行redis错误")
		return
	}
	message := messageService.GetMessageByReceiveID(v,false)
	b , err := json.Marshal ( message ) 
    if err != nil { 
		fmt. Println ( "error:" , err ) 
		return
	} 
	fmt.Fprintln(w, string(b))
}

//SendReadBySendID 把所有SendID和ReceiveID匹配设置为已读信息
func SendReadBySendID(w http.ResponseWriter, r *http.Request){
	receiveID,err := tool.GetLoginedID(r)
	if err!=nil{
		fmt.Println("执行redis错误")
		return
	}
	err = r.ParseForm()
	if err != nil {
		fmt.Println("解析表单数据失败!")
		fmt.Fprintln(w, "别乱发数据啊，格式都不对")
		return
	}
	sendID,err := strconv.Atoi(r.PostForm.Get("SendID"))
	if err!=nil{
		fmt.Println("转换字符串错误")
		return
	}
	messageService.SetReadMessage(sendID,receiveID)
}