package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "grpc-demo/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// 连接到gRPC服务器
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("连接失败: %v", err)
	}
	defer conn.Close()

	// 创建客户端
	client := pb.NewGreeterClient(conn)

	// 1. 测试简单的 SayHello
	log.Println("===== 测试 1: 简单的 SayHello =====")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	resp, err := client.SayHello(ctx, &pb.HelloRequest{Name: "张三"})
	if err != nil {
		log.Fatalf("调用 SayHello 失败: %v", err)
	}
	log.Printf("收到响应: %s (count: %d)", resp.Message, resp.Count)

	// 2. 测试服务器流式 RPC
	log.Println("\n===== 测试 2: 服务器流式 RPC =====")
	stream, err := client.SayHelloMultipleTimes(context.Background(), &pb.HelloRequest{Name: "李四"})
	if err != nil {
		log.Fatalf("调用 SayHelloMultipleTimes 失败: %v", err)
	}

	// 接收流式响应
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			// 流结束
			log.Println("流式响应接收完毕")
			break
		}
		if err != nil {
			log.Fatalf("接收流失败: %v", err)
		}
		log.Printf("收到流式响应: %s (count: %d)", resp.Message, resp.Count)
	}

	log.Println("\n客户端演示完成!")
}
