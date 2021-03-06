package main

import (
	"context"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/logger"
	"github.com/asim/go-micro/v3/registry"
	"github.com/isfk/go-micro-plugins/registry/nacos/v3"
	"go-micro-learning/micro-demo/sum-srv/handler"
	"go-micro-learning/micro-demo/sum-srv/proto/sum"
	"os"
)

var etcdReg registry.Registry

const (
	defaultNacosAddr      = "123.56.239.121:8848"
	defaultNacosNamespace = "dev"
)

func main() {
	// 从环境变量中获取nacos的ip和port
	var nacosAddr string
	nacosAddr = os.Getenv("NacosAddr")
	if nacosAddr == "" {
		nacosAddr = defaultNacosAddr
	}
	var nacosNamespace string
	nacosNamespace = os.Getenv("NacosNamespace")
	if nacosNamespace == "" {
		nacosNamespace = defaultNacosNamespace
	}

	nacosRegistry := nacos.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{nacosAddr}
		options.Context = context.WithValue(context.Background(), &nacos.NacosNamespaceContextKey{}, nacosNamespace)

	})
	service := micro.NewService(
		micro.Name("sum-srv"),
		micro.Registry(nacosRegistry),
		micro.Address(":8081"),
	)
	//服务初始化
	service.Init(
		micro.BeforeStart(func() error {
			logger.Info("sum-srv服务启动前日志")
			return nil
		}),
		micro.AfterStart(func() error {
			logger.Info("sum-srv服务启动后日志")
			return nil
		}),
	)

	sum.RegisterSumHandler(service.Server(), handler.Handler())

	if err := service.Run(); err != nil {
		panic(err)
	}

}
