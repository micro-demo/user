package repostory

import "user/domain/model"

type IUserRespository interface {
	//InitTable 初始化数据表
	InitTable() error
	//FindUserByName 根据用户名称查找用户信息
	FindUserByName(string2 string) (*model.User, error)
	//FindUserByID 根据用户ID查找用户信息
	FindUserByID(int642 int64) (*model.User, error)
	//CreateUser 创建用户
	CreateUser(user *model.User) (int64, error)
	//DeleteUserID 根据用户id删除用户
	DeleteUserID(int642 int64) error
	//UpdateUser 更新用户信息
	UpdateUser(user *model.User) error
	//FindAll 查询所有用户
	FindAll() ([]model.User, error)
}
