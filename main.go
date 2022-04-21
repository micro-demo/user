package main

import (
	"fmt"
	"github.com/micro/micro/v3/service"
	"github.com/micro/micro/v3/service/logger"
	user_handler "user/handler"
	user_proto "user/proto"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("user"),
		service.Version("1.0"),
	)

	// Register handler
	err := user_proto.RegisterUserHandler(srv.Server(), user_handler.New())

	if err != nil {
		fmt.Println("错误启动")
	}

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
