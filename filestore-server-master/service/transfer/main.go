package main

import (
	"fmt"
	"log"
	"time"

	"filestore-server/common"
	"filestore-server/config"
	"filestore-server/mq"
	dbproxy "filestore-server/service/dbproxy/client"
	"filestore-server/service/transfer/process"

	//_ "github.com/micro/go-plugins/registry/kubernetes"
	cli "github.com/urfave/cli/v2"
	micro "go-micro.dev/v4"

	"github.com/go-micro/plugins/v4/registry/consul"
	"go-micro.dev/v4/registry"
)

func startRPCService() {
	r := consul.NewRegistry( // 用于服务注册
		registry.Addrs("127.0.0.1:8500"), // 默认地址，可以不写
	)
	service := micro.NewService(
		micro.Name("go.micro.service.transfer"), // 服务名称
		micro.RegisterTTL(time.Second*10),       // TTL指定从上一次心跳间隔起，超过这个时间服务会被服务发现移除
		micro.RegisterInterval(time.Second*5),   // 让服务在指定时间内重新注册，保持TTL获取的注册时间有效
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

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}

func startTranserService() {
	if !config.AsyncTransferEnable {
		log.Println("异步转移文件功能目前被禁用，请检查相关配置")
		return
	}
	log.Println("文件转移服务启动中，开始监听转移任务队列...")

	// 初始化mq client
	mq.Init()

	mq.StartConsume(
		config.TransOSSQueueName,
		"transfer_oss",
		process.Transfer)
}

func main() {
	// 文件转移服务
	go startTranserService()

	// rpc 服务
	startRPCService()
}
