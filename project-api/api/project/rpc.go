package project

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
	"log"
	"test.com/project-common/discovery"
	"test.com/project-common/logs"
	"test.com/project-grpc/project"
	"test.com/project-grpc/task"

	"test.com/project-user/config"
)

var ProjectServiceClient project.ProjectServiceClient
var TaskServiceClient task.TaskServiceClient

func InitRpcProjectClient() {
	etcdRegister := discovery.NewResolver(config.C.EtcdConfig.Addrs, logs.LG)
	// 让 gRPC 知道 etcd 是服务发现后端
	resolver.Register(etcdRegister)
	// 确定使用etcd来实现服务发现,触发 etcd 解析器去发现并解析服务实例的具体地址，最终建立连接。
	conn, err := grpc.Dial("etcd:///project", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	ProjectServiceClient = project.NewProjectServiceClient(conn)
	TaskServiceClient = task.NewTaskServiceClient(conn)
}
