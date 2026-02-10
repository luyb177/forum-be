package client

import (
	pbp "github.com/Muxi-X/forum-be/microservice/post/proto"
	"github.com/Muxi-X/forum-be/pkg/handler"
	"github.com/Muxi-X/forum-be/pkg/identity"
	"github.com/go-micro/plugins/v4/registry/etcd"
	_ "github.com/go-micro/plugins/v4/registry/kubernetes"
	opentracingWrapper "github.com/go-micro/plugins/v4/wrapper/trace/opentracing"
	"github.com/opentracing/opentracing-go"
	"github.com/spf13/viper"
	micro "go-micro.dev/v4"
	"go-micro.dev/v4/registry"
)

var PostClient pbp.PostService

func PostInit() {
	r := etcd.NewRegistry(
		registry.Addrs(viper.GetString("etcd.addr")),
		etcd.Auth(viper.GetString("etcd.username"), viper.GetString("etcd.password")),
	)
	service := micro.NewService(
		micro.Name("forum.cli.post"),
		micro.Registry(r),
		micro.WrapClient(
			opentracingWrapper.NewClientWrapper(opentracing.GlobalTracer()),
		),
		micro.WrapCall(handler.ClientErrorHandlerWrapper()))

	service.Init()

	PostClient = pbp.NewPostService(identity.Prefix()+"forum.service.post", service.Client())
}
