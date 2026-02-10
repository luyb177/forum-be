package client

import (
	pbf "github.com/Muxi-X/forum-be/microservice/feed/proto"
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

var FeedClient pbf.FeedService

func FeedInit() {
	r := etcd.NewRegistry(
		registry.Addrs(viper.GetString("etcd.addr")),
		etcd.Auth(viper.GetString("etcd.username"), viper.GetString("etcd.password")),
	)
	service := micro.NewService(
		micro.Name("forum.cli.feed"),
		micro.Registry(r),
		micro.WrapClient(
			opentracingWrapper.NewClientWrapper(opentracing.GlobalTracer()),
		),
		micro.WrapCall(handler.ClientErrorHandlerWrapper()))

	service.Init()

	FeedClient = pbf.NewFeedService(identity.Prefix()+"forum.service.feed", service.Client())
}
