package main

import (
	"context"
	httpServer "github.com/asim/go-micro/plugins/server/http/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/registry"
	"github.com/asim/go-micro/v3/server"
	"github.com/gin-gonic/gin"
	"github.com/isfk/go-micro-plugins/registry/nacos/v3"
	"github.com/micro/micro/v3/service/logger"
	"go-micro-learning/micro-demo/http-client/handler"
)

var etcdRegistry registry.Registry

const nacosNamespace = "dev"

func main() {

	nacosRegistry := nacos.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{"123.56.239.121:8848"}
		// 支持 namespace
		options.Context = context.WithValue(context.Background(), &nacos.NacosNamespaceContextKey{}, nacosNamespace)

	})
	srv := httpServer.NewServer(
		server.Name("http-client"),
		server.Address(":8888"),
	)

	router := gin.Default()
	// 注册router
	sumHandler := handler.NewSumHandler()
	sumHandler.Getsum(router)
	newHandler := srv.NewHandler(router)
	if err := srv.Handle(newHandler); err != nil {
		logger.Fatal(err)
	}

	// Create service
	service := micro.NewService(
		micro.Server(srv),
		micro.Registry(nacosRegistry),
		//micro.Registry(etcd.NewRegistry()),
	)
	service.Init()

	// Run service
	if err := service.Run(); err != nil {
		logger.Fatal(err)
	}

}
