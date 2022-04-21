package handler

import (
	"context"
	log "github.com/micro/micro/v3/service/logger"
	"google.golang.org/protobuf/types/known/structpb"
	user "user/proto"
)

type User struct{}

func New() *User {
	return &User{}
}

func (e *User) UserRegister(ctx context.Context, req *user.UserRegisterRequest, rsp *user.CommonResponse) error {
	log.Info("Received User.Call request")
	rsp.Code = "0"       //登陆成功
	rsp.Message = "登陆成功" //登陆成功
	rsp.Data, _ = structpb.NewValue(make(map[string]interface{}))
	return nil
}

func (e *User) UserLogin(ctx context.Context, req *user.UserLoginRequest, rsp *user.CommonResponse) error {
	log.Info("Received User.UserInfo request")
	rsp.Code = "0"       //登陆成功
	rsp.Message = "登陆成功" //登陆成功
	rsp.Data, _ = structpb.NewValue(1)
	return nil
}

// 当访问get h
//ttp://localhost:8080/user/UserInfo的时候调用这个处理函数
func (e *User) GetUserInfo(ctx context.Context, req *user.GetUserInfoRequest, rsp *user.CommonResponse) error {
	rsp.Code = "0"       //登陆成功
	rsp.Message = "登陆成功" //登陆成功
	rsp.Data, _ = structpb.NewValue(make(map[string]interface{}))
	return nil
}
