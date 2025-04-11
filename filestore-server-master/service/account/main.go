package main

import (
	"log"
	"time"

	//_ "github.com/micro/go-plugins/registry/kubernetes"
	micro "go-micro.dev/v4"

	// k8s "github.com/micro/kubernetes/go/micro"

	"filestore-server/common"
	"filestore-server/service/account/handler"
	proto "filestore-server/service/account/proto/user"
	dbproxy "filestore-server/service/dbproxy/client"

	"github.com/go-micro/plugins/v4/registry/consul"
	"go-micro.dev/v4/registry"
)

func main() {
	r := consul.NewRegistry( // 用于服务注册
		registry.Addrs("127.0.0.1:8500"), // 默认地址，可以不写
	)
	service := micro.NewService(
		// service := k8s.NewService(
		micro.Name("go.micro.service.user"),
		micro.RegisterTTL(time.Second*10),
		micro.RegisterInterval(time.Second*5),
		micro.Flags(common.CustomFlags...),
		micro.Registry(r),
	)

	// 初始化service, 解析命令行参数等
	service.Init()

	// 初始化dbproxy client
	dbproxy.Init(service)

	proto.RegisterUserServiceHandler(service.Server(), new(handler.User))
	if err := service.Run(); err != nil {
		log.Println(err)
	}
}
