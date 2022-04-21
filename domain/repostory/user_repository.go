package repostory

//对db的操作全部放在这里 相当于dao层
import (
	"github.com/jinzhu/gorm"
	"user/domain/model"
)

type IUserRespository interface {
	//InitTable 初始化数据表
	InitTable() error
	//FindUserByName 根据用户名称查找用户信息
	FindUserByName(string) (*model.User, error)
	//FindUserByID 根据用户ID查找用户信息
	FindUserByID(int64) (*model.User, error)
	//CreateUser 创建用户
	CreateUser(*model.User) (int64, error)
	//DeleteUserID 根据用户id删除用户
	DeleteUserID(int64) error
	//UpdateUser 更新用户信息
	UpdateUser(*model.User) error
	//FindAll 查询所有用户
	FindAll() ([]model.User, error)
}

type UserRepository struct {
	mysqlDB *gorm.DB
}

// InitTable 初始化表结构
func (u *UserRepository) InitTable() error {
	return u.mysqlDB.CreateTable(&model.User{}).Error
}

//FindUserByName 根据用户名称查找用户信息
func (u *UserRepository) FindUserByName(name string) (user *model.User, err error) {
	user = &model.User{}
	return user, u.mysqlDB.Where("user_name = ?", name).Find(user).Error
}

//FindUserByID 根据用户ID查找用户信息
func (u *UserRepository) FindUserByID(userID int64) (user *model.User, err error) {
	user = &model.User{}
	return user, u.mysqlDB.First(user, userID).Error
}

//CreateUser 创建用户
func (u *UserRepository) CreateUser(user *model.User) (userID int64, err error) {
	return user.ID, u.mysqlDB.Create(user).Error
}

//DeleteUserID 根据用户id删除用户
func (u *UserRepository) DeleteUserID(userID int64) error {
	return u.mysqlDB.Where("id = ?", userID).Delete(&model.User{}).Error
}

//UpdateUser 更新用户信息
func (u *UserRepository) UpdateUser(user *model.User) error {
	return u.mysqlDB.Model(user).Update(&user).Error
}

//FindAll 查询所有用户
func (u *UserRepository) FindAll() (userAll []model.User, err error) {
	return userAll, u.mysqlDB.Find(&userAll).Error
}
