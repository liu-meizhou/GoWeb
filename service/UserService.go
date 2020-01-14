//Package service 提供对数据库操作的服务
package service

import (
	"mapper"
	"pojo"
)

// UserService 是
type UserService interface {
	GetAllUsers() []pojo.User
	GetUserByName(name string) []pojo.User
	GetUserByEmail(email string) []pojo.User
	GetUserByPhone(phone string) []pojo.User
	AddUser(user *pojo.User) int
	UpdateUser(user *pojo.User) int
	DeleteUserByID(id int)int
}
// UserServiceImpl 是
type UserServiceImpl struct {}


var userMapper mapper.UserMapper
var userService UserService


func init()  {
	userMapper = mapper.GetUserMapper()
	userService = &UserServiceImpl{}
}

// GetUserService 是获取用户操作服务
func GetUserService() UserService {
	return userService
}

// GetAllUsers 连接数据库查询所有人员
func (p *UserServiceImpl) GetAllUsers() []pojo.User{
	sql := `SELECT "id", "username", "email", "phone", "pwd" FROM allusers order by id`
	return userMapper.GetUserBy(sql)
}

// GetUserByName 用户名获取User
func (p *UserServiceImpl) GetUserByName(name string) []pojo.User{
	sql := `SELECT "id", "username", "email", "phone", "pwd" FROM allusers where username = $1 order by id`
	return userMapper.GetUserBy(sql,name)
}

// GetUserByEmail Email获取User
func (p *UserServiceImpl) GetUserByEmail(email string) []pojo.User{
	sql := `SELECT "id", "username", "email", "phone", "pwd" FROM allusers where email = $1 order by id`
	return userMapper.GetUserBy(sql,email)
}

// GetUserByPhone phone获取User
func (p *UserServiceImpl) GetUserByPhone(phone string) []pojo.User{
	sql := `SELECT "id", "username", "email", "phone", "pwd" FROM allusers where phone = $1 order by id`
	return userMapper.GetUserBy(sql,phone)
}

// AddUser 添加一个用户
func (p *UserServiceImpl) AddUser(user *pojo.User)(start int){
	sql := `INSERT INTO allusers(username,email,phone,pwd) VALUES($1,$2,$3,$4) RETURNING id`
	return userMapper.Update(sql,true,user.Name,user.Email,user.Phone,user.Pwd)
}

// UpdateUser 修改一个用户
func (p *UserServiceImpl) UpdateUser(user *pojo.User)(start int){
	sql := `UPDATE allusers SET username=$1,email=$2,phone=$3,pwd=$4 WHERE id=$5;`
	return userMapper.Update(sql,false,user.Name,user.Email,user.Phone,user.Pwd,user.ID)
}

// DeleteUserByID 删除一个用户
func (p *UserServiceImpl) DeleteUserByID(id int)(start int){
	sql := `DELETE FROM allusers WHERE id=$1;`
	return userMapper.Update(sql,false,id);
}

