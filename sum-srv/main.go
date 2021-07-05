package main

import (
	"context"
	"github.com/asim/go-micro/plugins/registry/etcd/v3"
	"github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/logger"
	"github.com/asim/go-micro/v3/registry"
	"github.com/isfk/go-micro-plugins/registry/nacos/v3"
	"go-micro-learning/micro-demo/proto/sum"
	"go-micro-learning/micro-demo/sum-srv/handler"
)

var etcdReg registry.Registry

func init() {
	// 注册中心修改为etcd
	etcdReg = etcd.NewRegistry(registry.Addrs("127.0.0.1:2379"))
}

const nacosNamespace = "dev"

func main() {

	nacosRegistry := nacos.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{"123.56.239.121:8848"}
		options.Context = context.WithValue(context.Background(), &nacos.NacosNamespaceContextKey{}, nacosNamespace)

	})
	service := micro.NewService(
		micro.Name("sum-srv"),
		micro.Registry(nacosRegistry),
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
