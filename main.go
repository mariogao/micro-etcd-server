package main

import (
	"github.com/mariogao/micro-etcd-server/hanlder"
	proto "github.com/mariogao/micro-etcd-server/proto"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
)

func main() {

	// create a new service
	service := micro.NewService(
		micro.Name("helloworld"),
		//etcd注册服务
		micro.Registry(etcd.NewRegistry(registry.Addrs("127.0.0.1:2379"))),
	)

	// initialise flags
	service.Init()
	proto.RegisterHiHandler(service.Server(), &hanlder.MyhelloStruct{})

	// start the service
	service.Run()

}
