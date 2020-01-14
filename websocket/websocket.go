package socket

import (
	"strconv"
	"fmt"
	"encoding/json"
	"net/http"

	"github.com/gorilla/websocket"
    
	"pojo"
	"redis"
	"service"
)

// ClientManager 客户端管理
type ClientManager struct {
	//客户端 map 储存并管理所有的长连接client，在线的为true，不在的为false
	clients map[*Client]bool
	//web端发送来的的message我们用broadcast来接收，并最后分发给所有的client
	broadcast chan []byte
	//新创建的长连接client
	register chan *Client
	//新注销的长连接client
	unregister chan *Client
}

//Client 客户端 
type Client struct {
	//用户id
	id string
	//连接的socket
	socket *websocket.Conn
	//发送的消息
	send chan []byte
}

//会把Message格式化成json pojo.message

//创建客户端管理者
var manager = ClientManager{
	broadcast:  make(chan []byte),
	register:   make(chan *Client),
	unregister: make(chan *Client),
	clients:    make(map[*Client]bool),
}
//GetManager 获取客户端管理器
func GetManager()(*ClientManager){
	return &manager
}

//Start 启动manage
func (manager *ClientManager) Start() {
	for {
		select {
		//如果有新的连接接入,就通过channel把连接传递给conn
		case conn := <-manager.register:
			//把客户端的连接设置为true
			manager.clients[conn] = true
			fmt.Println("/A new socket has connected.")
			//把返回连接成功的消息json格式化
			// jsonMessage, _ := json.Marshal(pojo.NewMessageContext("/A new socket has connected."))
			// //调用客户端的send方法，发送消息
			// manager.send(jsonMessage, conn)
		//如果连接断开了
		case conn := <-manager.unregister:
			//判断连接的状态，如果是true,就关闭send，删除连接client的值
			if _, ok := manager.clients[conn]; ok {
				close(conn.send)
				delete(manager.clients, conn)
				fmt.Println("A socket has disconnected.")
				// jsonMessage, _ := json.Marshal(pojo.NewMessageContext("/A socket has disconnected."))
				// manager.send(jsonMessage, conn)
			}
			//消息存入数据库， 用户在线则再加上转发
		case message := <-manager.broadcast:
			//解析消息
			var data pojo.Message
			json.Unmarshal(message, &data)
			data.SetSendTimeNow()
			data.IsRead=false
			receiveID := strconv.Itoa(data.ReceiveID)
			message,_ = json.Marshal(data)
			//遍历已经连接的客户端，把消息发送给他
			for conn := range manager.clients {
				if conn.id==receiveID{
					data.IsRead=true
					select {
					case conn.send <- message:
					default:
						data.IsRead=false
						close(conn.send)
						delete(manager.clients, conn)
					}
				}
				
			}
			//存入数据库
			service.GetMessageService().AddMessage(&data)
		}
	}
}

//定义客户端管理的send方法
func (manager *ClientManager) send(message []byte, ignore *Client) {
	for conn := range manager.clients {
		//不给屏蔽的连接发送消息
		if conn != ignore {
			conn.send <- message
		}
	}
}

//定义客户端结构体的read方法
func (c *Client) read() {
	defer func() {
		manager.unregister <- c
		c.socket.Close()
	}()

	for {
		//读取消息
		_, message, err := c.socket.ReadMessage()
		//如果有错误信息，就注销这个连接然后关闭
		if err != nil {
			manager.unregister <- c
			c.socket.Close()
			break
		}
		//如果没有错误信息就把信息放入broadcast 
		id,err := strconv.Atoi(c.id)
		if err!=nil{
			fmt.Println("传入格式错误")
			continue
		}
		//解析消息
		var data pojo.Message
		json.Unmarshal(message, &data)
		data.SendID=id
		jsonMessage, _ := json.Marshal(data)
		manager.broadcast <- jsonMessage
	}
}

func (c *Client) write() {
	defer func() {
		c.socket.Close()
	}()

	for {
		select {
		//从send里读消息
		case message, ok := <-c.send:
			//如果没有消息
			if !ok {
				c.socket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			//有消息就写入，发送给web端
			c.socket.WriteMessage(websocket.TextMessage, message)
		}
	}
}

//WsHandler 长连接处理器
func WsHandler(res http.ResponseWriter, req *http.Request) {
	//将http协议升级成websocket协议
	conn, err := (&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}).Upgrade(res, req, nil)
	if err != nil {
		http.NotFound(res, req)
		return
	}
	//获取Cokkie进行身份认证
	//读取消息
	_, message, err := conn.ReadMessage()
	//如果有错误信息，就注销这个连接然后关闭
	if err != nil {
		conn.Close()
		return
	}
	cookieValue := string(message[:])
	val,err := redisdb.GetDo(cookieValue)
	if err!=nil{
		conn.Close()
		return
	}
	//每一次连接都会新开一个client，client.id通过uuid生成保证每次都是不同的
	client := &Client{id: val, socket: conn, send: make(chan []byte)}
	//注册一个新的链接
	manager.register <- client

	//启动协程收web端传过来的消息
	go client.read()
	//启动协程把消息返回给web端
	go client.write()
}


