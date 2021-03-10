package main


import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/broker"
	"github.com/micro/go-micro/config/cmd"
	"github.com/micro/go-micro/server"
	_ "github.com/micro/go-plugins/registry/consul"
	log "github.com/sirupsen/logrus"
	"gomicrosrv/handler"
	pb "gomicrosrv/proto"
	"gomicrosrv/conf"
)

func main() {
	service := micro.NewService(
		micro.Name("aiops.micro.srv.v2.switch"),
		micro.Version("latest"),
		// micro.RegisterTTL(30*time.Second),      // 服务被发现之后注册的信息存在多长时间，然后过期并被删除
		// micro.RegisterInterval(15*time.Second), // 服务应该重新注册的时间间隔
	)

	service.Init()
	conf.Init()
	pb.RegisterUserServiceHandler(service.Server(), handler.NewUserRpcHandler(), server.InternalHandler(true))

	// broker 初始化
	cmd.Init()
	if err := broker.Init(); err != nil {
		log.Fatal("Broker 初始化失败:", err)
	}
	if err := broker.Connect(); err != nil {
		log.Fatal("Broker 连接失败:", err)
	}
	defer broker.Disconnect()

	if err := service.Run(); err != nil {
		log.Fatal("启动失败：", err)
	}
}
