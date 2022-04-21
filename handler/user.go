package handler

import (
	"context"

	log "github.com/micro/micro/v3/service/logger"

	user "user/proto"
)

type User struct{}

// Return a new handler
func New() *User {
	return &User{}
}

// Call is a single request handler called via client.Call or the generated client code
//curl "http://localhost:8080/user/Call"
func (e *User) Call(ctx context.Context, req *user.Request, rsp *user.Response) error {
	log.Info("Received User.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// XiongYao is a single request handler called via client.Call or the generated client code
// curl "http://localhost:8080/user/UserInfo"
// 当访问http://localhost:8080/user/UserInfo的时候调用这个处理函数
func (e *User) UserInfo(ctx context.Context, req *user.Request, rsp *user.Response) error {
	log.Info("Received User.UserInfo request")
	rsp.Msg = req.Name + "是大帅比"
	return nil
}
