package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	pb "grpc-demo/proto"

	"google.golang.org/grpc"
)

// server 是用来实现 greeter.GreeterServer 的类型
type server struct {
	pb.UnimplementedGreeterServer
}

// SayHello 实现简单的问候
func (s *server) SayHello(_ context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Printf("收到请求: name=%s", req.Name)

	return &pb.HelloResponse{
		Message: fmt.Sprintf("你好, %s! 欢迎学习 gRPC 和 Protocol Buffers", req.Name),
		Count:   1,
	}, nil
}

// SayHelloMultipleTimes 实现服务器流式RPC，返回多条问候消息
func (s *server) SayHelloMultipleTimes(req *pb.HelloRequest, stream pb.Greeter_SayHelloMultipleTimesServer) error {
	log.Printf("收到流式请求: name=%s", req.Name)

	// 发送5条消息
	for i := 1; i <= 5; i++ {
		resp := &pb.HelloResponse{
			Message: fmt.Sprintf("第 %d 次问候: 你好, %s!", i, req.Name),
			Count:   int32(i),
		}

		if err := stream.Send(resp); err != nil {
			return err
		}

		log.Printf("发送消息 %d/5", i)
		time.Sleep(time.Second) // 模拟处理时间
	}

	return nil
}

func main() {
	// 监听TCP端口
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("监听失败: %v", err)
	}

	// 创建gRPC服务器
	s := grpc.NewServer()

	// 注册服务
	pb.RegisterGreeterServer(s, &server{})

	log.Printf("gRPC 服务器启动，监听端口 :50051")

	// 开始服务
	if err := s.Serve(lis); err != nil {
		log.Fatalf("服务失败: %v", err)
	}
}
