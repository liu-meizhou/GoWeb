package main

import (
	"log"
	"net/http"

	
	"websocket"
	"controller"
	"filter"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./resources/static/")))
	http.Handle("/JS/", http.StripPrefix("/JS/", http.FileServer(http.Dir("./resources/js/"))))
	http.Handle("/Logined/", http.StripPrefix("/Logined/", filter.SafeFileServer(http.Dir("./resources/upms/"))))
	//登录注册操作
	//游客
	http.HandleFunc("/visitor", controller.Visitor)
	//登录
	http.HandleFunc("/Login", controller.Login)
	http.HandleFunc("/Logined/Login", filter.SafeHandler(controller.Index))
	//注册
	http.HandleFunc("/Register", controller.Register)
	//注销
	http.HandleFunc("/Logined/Logout", filter.SafeHandler(controller.Logout))

	//用户管理系统
	//显示所有用户信息
	http.HandleFunc("/Logined/GetUserMgr", filter.SafeHandler(controller.GetUserMgr))
	//添加用户 
	//修改用户
	http.HandleFunc("/Logined/Updata", filter.SafeHandler(controller.UpdateByID))
	//删除用户
	http.HandleFunc("/Logined/Delete", filter.SafeHandler(controller.DeleteByID))
	//联系用户

	//获取未读消息数量
	http.HandleFunc("/Message/GetCount", filter.SafeHandler(controller.GetNoReadCount))
	//获取每一个人的未读消息
	http.HandleFunc("/Message/GetEveryNoRead", filter.SafeHandler(controller.GetEveryNoRead))
	//把所有SendID和ReceiveID匹配设置为已读信息
	http.HandleFunc("/Message/SendReadBySendID", filter.SafeHandler(controller.SendReadBySendID))


	//WebSocket
	go socket.GetManager().Start()
	http.HandleFunc("/ws", socket.WsHandler)
	//http.HandleFunc("/ws/identity",filter.SafeHandler(socket.Identity))

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}
