package rpc

import (
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
	"log"
	"test.com/project-common/discovery"
	"test.com/project-common/logs"
	"test.com/project-grpc/user/login"

	"test.com/project-user/config"
)

var LoginServiceClient login.LoginServiceClient

func InitRpcUserClient() {
	etcdRegister := discovery.NewResolver(config.C.EtcdConfig.Addrs, logs.LG)
	// 让 gRPC 知道 etcd 是服务发现后端
	resolver.Register(etcdRegister)
	// 确定使用etcd来实现服务发现,触发 etcd 解析器去发现并解析服务实例的具体地址，最终建立连接。
	conn, err := grpc.Dial(
		"etcd:///user",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()),
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	LoginServiceClient = login.NewLoginServiceClient(conn)
}
