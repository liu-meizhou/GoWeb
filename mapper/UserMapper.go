//Package mapper 实现CRUD操作
package mapper

import (
	"fmt"
	"database/sql"

	"pojo"
	"DBconn"
)

//UserMapper 用户CRUD操作接口
type UserMapper interface {
	Update(sql string, isInsert bool, args... interface{}) (int)
	GetUserBy(sql string, args... interface{})([]pojo.User)
}

//UserMapperImpl 用户CRUD操作实现类
type UserMapperImpl struct {}

var db *sql.DB
var userMapper UserMapper
func init(){
	db=database.GetConnect()
	userMapper=&UserMapperImpl{}
}

// GetUserMapper 获取操作CRUD的已实现接口
func GetUserMapper() UserMapper {
	return userMapper
}

func getAllU(rows *sql.Rows) ([]pojo.User){
	var user pojo.User
    var users []pojo.User
    for rows.Next(){
        rows.Scan(&user.ID, &user.Name, &user.Email, &user.Phone,&user.Pwd)
        users = append(users, user)
	}
	return users
}

//GetUserBy Byxxx获取uesrs
func (p *UserMapperImpl) GetUserBy(sql string, args... interface{})([]pojo.User){
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
	users := getAllU(rows)
	return users
}


// Update 增删改,操作
func (p *UserMapperImpl) Update(sql string, isInsert bool, args... interface{}) (state int){
    stmt, err := db.Prepare(sql)
	defer stmt.Close()
	if err!=nil{
		fmt.Println("SQL语句设置失败")
		return -1
	}
	if isInsert{
		err = stmt.QueryRow(args...).Scan(&state)
		if err!=nil{
			fmt.Println("参数添加失败")
			return -2
		}
		return 
	}
	result , err := stmt.Exec(args...)
	if err!=nil{
		fmt.Println("参数添加失败")			
		return -2
	}
	num, err := result.RowsAffected()
	if err!=nil{
		fmt.Println("修改失败")
		return -3
	}
	fmt.Println("修改成功，修改行数",num)
	return 0
}
