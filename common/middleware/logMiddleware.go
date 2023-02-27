package middleware

import (
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
	"net/http"
	"qbq-open-platform/common/global"
)

//处理输出日志
func LogMiddleware() rest.Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			logx.WithContext(r.Context()).Infof("[--- 请求开始: %s:%s ---]", global.Config().ApplicationName, r.URL)
			next(w, r)
		}
	}
}
