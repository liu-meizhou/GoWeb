package pojo

import (
	"time"
)

// Message 表Message的实体类
type Message struct{
	SendID		int		`json:"SendID"`
	ReceiveID	int		`json:"ReceiveID"`
	Type		int		`json:"Type"`
	Context		string	`json:"Context"`
	SendTime	string	`json:"SendTime"`
	IsRead		bool	`json:"IsRead"`
}

// NewAMessage 获取一个实体类
func NewAMessage() *Message{
	return &Message{}
}

// NewMessage 获取一个实体类
func NewMessage(Sendid,ReceiveID,Type int,Context,SendTime string,IsRead bool) (*Message){
	return &Message{
		Sendid,
		ReceiveID,
		Type,
		Context,
		SendTime,
		IsRead,
	}
}

// NewMessageContext 获取一个实体类
func NewMessageContext(Context string) (*Message){
	message := NewAMessage()
	message.Context=Context
	return message
}


//SetSendTimeNow 设置SendTime为当前时间
func (p *Message) SetSendTimeNow(){
	p.SendTime=time.Now().Format("2006-01-02 15:04:05")
}