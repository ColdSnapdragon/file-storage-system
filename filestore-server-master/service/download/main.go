package main

import (
	"fmt"
	"time"

	//_ "github.com/micro/go-plugins/registry/kubernetes"
	micro "go-micro.dev/v4"

	"filestore-server/common"
	dbproxy "filestore-server/service/dbproxy/client"
	cfg "filestore-server/service/download/config"
	dlProto "filestore-server/service/download/proto/download"
	"filestore-server/service/download/route"
	dlRpc "filestore-server/service/download/rpc"

	"github.com/go-micro/plugins/v4/registry/consul"
	"go-micro.dev/v4/registry"
)

func startRPCService() {
	r := consul.NewRegistry( // 用于服务注册
		registry.Addrs("127.0.0.1:8500"), // 默认地址，可以不写
	)
	service := micro.NewService(
		micro.Name("go.micro.service.download"), // 在注册中心中的服务名称
		micro.RegisterTTL(time.Second*10),
		micro.RegisterInterval(time.Second*5),
		micro.Flags(common.CustomFlags...),
		micro.Registry(r),
	)
	service.Init()

	// 初始化dbproxy client
	dbproxy.Init(service)

	dlProto.RegisterDownloadServiceHandler(service.Server(), new(dlRpc.Download))
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}

func startAPIService() {
	router := route.Router()
	router.Run(cfg.DownloadServiceHost)
}

func main() {
	// api 服务
	go startAPIService()

	// rpc 服务
	startRPCService()
}
