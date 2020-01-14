//Package pojo 对应数据库的实体类
package pojo

// User 表users的实体类
type User struct{
	ID		int		`json:"ID"`
	Name	string	`json:"name"`
	Email	string	`json:"email"`
	Phone	string	`json:"phone"`
	Pwd		string	`json:"password"`
}

// NewAUser 获取一个实体类
func NewAUser() *User{
	return &User{}
}

// NewUser 获取一个实体类
func NewUser(id int,name,email,phone,pwd string) (*User){
	return &User{
		id,
		name,
		email,
		phone,
		pwd,
	}
}
