package unary

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"qbq-open-platform/common/global"
)

func ServerTokenInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		if len(md.Get(global.CACHE_AUTHENTICATION_KEY)) > 0 {
			authentication := md.Get(global.CACHE_AUTHENTICATION_KEY)[0]
			ctx = context.WithValue(ctx, global.CACHE_AUTHENTICATION_KEY, authentication)
		}
	}
	return handler(ctx, req)
}
