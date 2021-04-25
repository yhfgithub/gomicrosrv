package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/micro/cli"
	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/etcd"
	"github.com/micro/go-micro/web"
	log "github.com/sirupsen/logrus"
	"gomicroweb/conf"
	"gomicroweb/handler"
	pb "gomicroweb/util"
	"net/http"
	"strconv"
)

// 处理跨域请求,支持options访问
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE, PUT")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}

type Gin struct {
}

type Result struct {
	Id string `json:"id"`
	Message interface{} `json:"message"`
	Success bool `json:"success"`
}

var cl pb.UserService

//获取某个信息
func (g *Gin) SqlOneInfo(c *gin.Context) {
	id ,_:= strconv.ParseInt(c.Query("id"),10,64)
	//调用rpc
	response,_ := cl.GetUserById(context.TODO(),&pb.IdInfo{
		Id: id,
	})
	c.JSON(200,response)
}





//启用API作为一个网关或代理，来作为微服务访问的单一入口。它应该在您的基础架构的边缘运行。它将HTTP请求转换为RPC并转发给相应的服务。
//将HTTP请求转换为RPC并转发给相应的服务。以rpc方式调srv。(因为srv提供的只有rpc接口访问方式。)
func main(){
	conf.Init()
	reg := etcd.NewRegistry(func(op *registry.Options) {
		op.Addrs = []string{
			"127.0.0.1:2379",
		}
	})
	//reg := consul.NewRegistry(func(op *registry.Options) {
	//	op.Addrs = []string{
	//		"127.0.0.1:8500",
	//	}
	//})
	service:=web.NewService(
		web.Name("aiops.micro.api.v2.sql"),
		web.Address(":8012"),
		web.Registry(reg),
	)
	service.Init(
		web.Action(func(c *cli.Context) {
			cl = pb.NewUserService("aiops.micro.srv.v2.user", client.DefaultClient)
		}),
	)

	//创建 RestFul handler
	g := new(Gin)
	gHttp := new(handler.GinHttpHandler)


	router := gin.Default()
	//使用中间件进行 token 认证
	router.Use(Cors())
	r1:=router.Group("/v2/sql")
	r1.GET("/info",g.SqlOneInfo)
	r1.POST("/put",gHttp.PutInfo)
fmt.Println()
	//注册 handler
	service.Handle("/", router)

	//运行 api
	err := service.Run()
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Error("服务启动错误!")
		return
	}
}
