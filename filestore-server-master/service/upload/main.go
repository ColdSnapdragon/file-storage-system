package main

import (
	"filestore-server/config"
	"filestore-server/mq"
	"fmt"
	"log"
	"os"
	"time"

	//_ "github.com/micro/go-plugins/registry/kubernetes"
	cli "github.com/urfave/cli/v2"
	micro "go-micro.dev/v4"

	"filestore-server/common"
	dbproxy "filestore-server/service/dbproxy/client"
	cfg "filestore-server/service/upload/config"
	upProto "filestore-server/service/upload/proto/upload"
	"filestore-server/service/upload/route"
	upRpc "filestore-server/service/upload/rpc"

	"github.com/go-micro/plugins/v4/registry/consul"
	"go-micro.dev/v4/registry"
)

func startRPCService() {
	r := consul.NewRegistry( // 用于服务注册
		registry.Addrs("127.0.0.1:8500"), // 默认地址，可以不写
	)
	service := micro.NewService(
		micro.Name("go.micro.service.upload"), // 服务名称
		micro.RegisterTTL(time.Second*10),     // TTL指定从上一次心跳间隔起，超过这个时间服务会被服务发现移除
		micro.RegisterInterval(time.Second*5), // 让服务在指定时间内重新注册，保持TTL获取的注册时间有效
		micro.Flags(common.CustomFlags...),
		micro.Registry(r),
	)
	service.Init(
		micro.Action(func(c *cli.Context) error {
			// 检查是否指定mqhost
			mqhost := c.String("mqhost")
			if len(mqhost) > 0 {
				log.Println("custom mq address: " + mqhost)
				mq.UpdateRabbitHost(mqhost)
			}
			return nil
		}),
	)

	// 初始化dbproxy client
	dbproxy.Init(service)
	// 初始化mq client
	mq.Init()

	upProto.RegisterUploadServiceHandler(service.Server(), new(upRpc.Upload))
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}

func startAPIService() {
	router := route.Router()
	router.Run(cfg.UploadServiceHost)
	// service := web.NewService(
	// 	web.Name("go.micro.web.upload"),
	// 	web.Handler(router),
	// 	web.RegisterTTL(10*time.Second),
	// 	web.RegisterInterval(5*time.Second),
	// )
	// if err := service.Init(); err != nil {
	// 	log.Fatal(err)
	// }

	// if err := service.Run(); err != nil {
	// 	log.Fatal(err)
	// }
}

func main() {
	os.MkdirAll(config.TempLocalRootDir, 0777)
	os.MkdirAll(config.TempPartRootDir, 0777)

	// api 服务
	go startAPIService()

	// rpc 服务
	startRPCService()
}
