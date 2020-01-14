package mapper

import (
	"fmt"
	"database/sql"

	"pojo"
)

//MessageMapper 用户CRUD操作接口
type MessageMapper interface {
	GetMessageBy(sql string, args... interface{})([]pojo.Message)
	Update(sql string, isInsert bool, args... interface{}) (state int)
	SelectRowSql(sql string, args... interface{}) (result int)
}

//MessageMapperImpl 用户CRUD操作实现类
type MessageMapperImpl struct {}


var messageMapper MessageMapper
func init(){
	messageMapper=&MessageMapperImpl{}
}

// GetMessageMapper 获取操作CRUD的已实现接口
func GetMessageMapper() MessageMapper {
	return messageMapper
}

func getAllM(rows *sql.Rows) ([]pojo.Message){
	var message pojo.Message
    var messages []pojo.Message
    for rows.Next(){
        rows.Scan(&message.SendID, &message.ReceiveID, &message.Type, &message.Context,&message.SendTime,&message.IsRead)
        messages = append(messages, message)
	}
	return messages
}

//GetMessageBy Byxxx获取message
func (p *MessageMapperImpl) GetMessageBy(sql string, args... interface{})([]pojo.Message){
	stmt, err := db.Prepare(sql)
	if err != nil {
		fmt.Println("预编译失败，出现异常")
		return nil
    }
	defer stmt.Close()

	rows, err := stmt.Query(args...)
    if err != nil {
		fmt.Println("查询失败，出现异常")
		return nil
    }
	defer rows.Close()
	messages := getAllM(rows)
	return messages
}


// Update 增删改,操作
func (p *MessageMapperImpl) Update(sql string, isInsert bool, args... interface{}) (state int){
    stmt, err := db.Prepare(sql)
	defer stmt.Close()
	if err!=nil{
		fmt.Println("SQL语句设置失败")
		return -1
	}
	result , err := stmt.Exec(args...)
	if err!=nil{
		fmt.Println("参数添加失败")			
		return -2
	}
	if isInsert{
		result.LastInsertId()
		return 0
	}
	num, err := result.RowsAffected()
	if err!=nil{
		fmt.Println("修改失败")
		return -3
	}
	fmt.Println("修改成功，修改行数",num)
	return 0
}

// SelectRowSql 执行一条查询一行的Sql语句
func (p *MessageMapperImpl) SelectRowSql(sql string, args... interface{}) (result int){
	db.QueryRow(sql,args...).Scan(&result);
	return;
}

