package result

import (
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/timex"
	"github.com/zeromicro/go-zero/rest/httpx"
	"google.golang.org/grpc/status"
	"net/http"
	"qbq-open-platform/common/errorsEnums"
	"qbq-open-platform/common/global"
	"qbq-open-platform/common/utils"
	"time"
)

type Body struct {
	Code int32       `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

//http返回
func HttpResult(r *http.Request, w http.ResponseWriter, resp interface{}, err error, respJson bool, startTime time.Duration) {
	var body Body
	if err == nil {
		//成功返回
		body.Code = 200
		body.Msg = "OK"
		body.Data = resp

		if respJson {
			jsonStr, _ := utils.Struct2json(body)
			logx.WithContext(r.Context()).Infof("[--- 响应参数: %s ---]", jsonStr)
			httpx.WriteJson(w, http.StatusOK, body)
		} else {
			httpx.Ok(w)
		}
		logx.WithContext(r.Context()).WithDuration(timex.Since(startTime)).Infof("[--- 请求成功: %s:%s ---]", global.Config().ApplicationName, r.URL)
	} else {
		//错误返回
		errCode := errorsEnums.INTERNAL_SERVER_ERROR
		errMessage := errorsEnums.GetMsgByCode(errorsEnums.INTERNAL_SERVER_ERROR)

		causeErr := errors.Cause(err)
		if platformError, ok := errorsEnums.IsPlatformErr(causeErr); ok {
			errCode = platformError.GetErrCode()
			errMessage = platformError.GetErrMsg()
		} else {
			if gstatus, ok := status.FromError(causeErr); ok {
				grpcCode := int32(gstatus.Code())
				if errorsEnums.IsCodeErr(grpcCode) {
					errCode = grpcCode
					errMessage = gstatus.Message()
				}
			}
		}

		logx.WithContext(r.Context()).WithDuration(timex.Since(startTime)).Errorf("[--- 请求失败: %s:%s ---], %+v", global.Config().ApplicationName, r.URL, err)

		body.Code = errCode
		body.Msg = errMessage
		body.Data = nil
		httpx.WriteJson(w, http.StatusOK, body)
	}
}
