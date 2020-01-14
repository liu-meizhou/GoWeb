package service

import (
	"mapper"
	"pojo"
)

// MessageService 是
type MessageService interface {
	GetMessageByReceiveID(ReceiveID int,isRead bool) ([]pojo.Message)
	AddMessage(message *pojo.Message) (state int)
	GetNoReadCountMessageByReceiveID(ReceiveID int) (count int)
	SetReadMessage(SendID,ReceiveID int) (count int)
}
// MessageServiceImpl 是
type MessageServiceImpl struct {}


var messageMapper mapper.MessageMapper
var messageService MessageService


func init()  {
	messageMapper = mapper.GetMessageMapper()
	messageService = &MessageServiceImpl{}
}

// GetMessageService 是获取用户操作服务
func GetMessageService() MessageService {
	return messageService
}


//GetMessageByReceiveID 根据接受ID读取已读或未读信息
func (p *MessageServiceImpl) GetMessageByReceiveID(ReceiveID int,isRead bool) ([]pojo.Message){
	sql := `SELECT "sendid", "receiveid", "type", "context",to_char(sendtime,'yyyy-MM-dd HH24:MI:SS'), "isread" 
			FROM message 
			where receiveid = $1 AND isread = $2
			order by sendid;`
	return messageMapper.GetMessageBy(sql,ReceiveID,isRead);
}


//AddMessage 添加一条信息
func (p *MessageServiceImpl) AddMessage(message *pojo.Message) (state int){
	message.SetSendTimeNow()
	sql := `INSERT INTO message(sendid,receiveid,type,context,sendtime,isread) VALUES($1,$2,$3,$4,to_timestamp($5,'yyyy-MM-dd hh24:mi:ss'),$6);`
	return messageMapper.Update(sql,true, message.SendID, message.ReceiveID, message.Type, message.Context, message.SendTime, message.IsRead)
}

//GetNoReadCountMessageByReceiveID 通过ReceiveID获取未读信息数量
func (p *MessageServiceImpl) GetNoReadCountMessageByReceiveID(ReceiveID int) (count int){
	sql := `SELECT count(*) FROM message WHERE receiveid=$1 AND isread=FALSE`
	return messageMapper.SelectRowSql(sql,ReceiveID)
}

//SetReadMessage 设置已读信息
func (p *MessageServiceImpl) SetReadMessage(SendID,ReceiveID int) (count int){
	sql := `UPDATE message SET isread=TRUE WHERE sendid=$1 AND receiveid=$2`
	return messageMapper.Update(sql, false, SendID,ReceiveID)
}