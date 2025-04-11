package main

import (
	"filestore-server/common"
	"filestore-server/service/dbproxy/config"
	"log"
	"time"

	//_ "github.com/micro/go-plugins/registry/kubernetes"
	cli "github.com/urfave/cli/v2"
	"go-micro.dev/v4"

	dbConn "filestore-server/service/dbproxy/conn"
	dbProxy "filestore-server/service/dbproxy/proto/dbproxy"
	dbRpc "filestore-server/service/dbproxy/rpc"

	"github.com/go-micro/plugins/v4/registry/consul"
	"go-micro.dev/v4/registry"
)

func startRpcService() {
	r := consul.NewRegistry( // 用于服务注册
		registry.Addrs("127.0.0.1:8500"), // 默认地址，可以不写
	)
	service := micro.NewService(
		micro.Name("go.micro.service.dbproxy"), // 在注册中心中的服务名称
		micro.RegisterTTL(time.Second*10),      // 声明超时时间, 避免consul不主动删掉已失去心跳的服务节点
		micro.RegisterInterval(time.Second*5),
		micro.Flags(common.CustomFlags...),
		micro.Registry(r),
	)

	service.Init(
		micro.Action(func(c *cli.Context) error {
			// 检查是否指定dbhost
			dbhost := c.String("dbhost")
			if len(dbhost) > 0 {
				log.Println("custom db address: " + dbhost)
				config.UpdateDBHost(dbhost)
			}
			return nil
		}),
	)

	// 初始化db connection
	dbConn.InitDBConn()

	dbProxy.RegisterDBProxyServiceHandler(service.Server(), new(dbRpc.DBProxy))
	if err := service.Run(); err != nil {
		log.Println(err)
	}
}

func main() {
	startRpcService()
}

// res, err := mapper.FuncCall("/user/UserExist", []interface{}{"haha"}...)
// log.Printf("error: %+v\n", err)
// log.Printf("result: %+v\n", res[0].Interface())

// res, err = mapper.FuncCall("/user/UserExist", []interface{}{"admin"}...)
// log.Printf("error: %+v\n", err)
// log.Printf("result: %+v\n", res[0].Interface())
