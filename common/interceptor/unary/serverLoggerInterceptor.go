package unary

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/timex"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"qbq-open-platform/common/errorsEnums"
	"qbq-open-platform/common/global"
)

func ServerLoggerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	log := logx.WithContext(ctx)
	requestStart := timex.Now()
	log.Infof("[--- 请求开始: %s:%s ---]", global.Config().ApplicationName, info.FullMethod)
	if req != nil {
		if jsonByte, err := json.Marshal(req); err == nil {
			log.Infof("[--- 请求参数: %s ---]", string(jsonByte))
		}
	} else {
		log.Info("[--- 请求参数: 无 ---]")
	}

	resp, err = handler(ctx, req)
	if err != nil {
		platformError, ok := errorsEnums.IsPlatformErr(err)
		if ok {
			log.WithContext(ctx).WithDuration(timex.Since(requestStart)).Errorf("[--- 请求失败: %s:%s ---]，%+v", global.Config().ApplicationName, info.FullMethod, err)
		} else {
			platformError = errorsEnums.NewErrCodeMsg(errorsEnums.INTERNAL_SERVER_ERROR, err.Error())
			log.WithContext(ctx).WithDuration(timex.Since(requestStart)).Errorf("[--- 请求失败: %s:%s ---]，%+v", global.Config().ApplicationName, info.FullMethod, err)
		}
		//转成grpc err
		//TODO 这里可以把rpc调用详情传递回去
		errStatus := status.New(codes.Code(platformError.GetErrCode()), platformError.GetErrMsg())
		return resp, errStatus.Err()
	} else {
		if resp != nil {
			if jsonByte, err := json.Marshal(resp); err == nil {
				log.WithContext(ctx).WithDuration(timex.Since(requestStart)).Infof("[--- 响应参数: %s ---]", string(jsonByte))
			}
		} else {
			log.WithContext(ctx).WithDuration(timex.Since(requestStart)).Info("[--- 响应参数: 无 ---]")
		}
		log.WithContext(ctx).WithDuration(timex.Since(requestStart)).Infof("[--- 请求成功: %s:%s ---]", global.Config().ApplicationName, info.FullMethod)
	}
	return resp, nil
}
