package unary

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"qbq-open-platform/common/global"
)

func ClientTokenInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	if ctx.Value(global.CACHE_AUTHENTICATION_KEY) != nil {
		authentication := ctx.Value(global.CACHE_AUTHENTICATION_KEY).(string)
		var pairs []string
		pairs = append(pairs, global.CACHE_AUTHENTICATION_KEY, authentication)
		ctx = metadata.AppendToOutgoingContext(ctx, pairs...)
	}

	err := invoker(ctx, method, req, reply, cc, opts...)
	return err
}
