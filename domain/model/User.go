package model

//	数据库model的映射关系
type User struct {
	//主键
	ID int64 `gorm:"primary_key;not_null;auto_increment"`
	//用户名
	UserName string `gorm:"unique_index;not_null"`
	//
	FirstName string
	//加密后密码
	HashPassword string
}

// 用户信息序列化json
func (u User) UserInfo() map[string]interface{} {
	data := make(map[string]interface{})
	data["id"] = u.ID
	data["username"] = u.UserName
	data["first"] = u.FirstName
	return data
}
