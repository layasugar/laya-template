package middlewares

import (
	"context"
	"log"

	"google.golang.org/grpc"
)

// TestInterceptor 测试拦截器
func TestInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Print("this is test interceptor")
	resp, err := handler(ctx, req)
	return resp, err
}
