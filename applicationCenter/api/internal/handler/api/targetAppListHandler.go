package api

import (
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/timex"
	"net/http"
	"qbq-open-platform/common"
	"qbq-open-platform/common/errorsEnums"
	"qbq-open-platform/common/result"
	"qbq-open-platform/common/utils"

	"github.com/zeromicro/go-zero/rest/httpx"
	"qbq-open-platform/applicationCenter/api/internal/logic/api"
	"qbq-open-platform/applicationCenter/api/internal/svc"
	"qbq-open-platform/applicationCenter/api/internal/types"
)

func TargetAppListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := timex.Now()
		var req types.ApiTargetAppListReq
		//解析入参
		if err := httpx.Parse(r, &req); err != nil {
			result.HttpResult(r, w, nil, errorsEnums.NewErrCodeMsg(errorsEnums.PARAM_MISS, err.Error()), true, startTime)
			return
		}
		//参数验证
		if err := common.Validate(r.Context(), &req); err != nil {
			result.HttpResult(r, w, nil, err, true, startTime)
			return
		}
		//执行业务逻辑
		jsonStr, _ := utils.Struct2json(&req)
		logx.WithContext(r.Context()).Infof("[--- 请求参数: %s ---]", jsonStr)

		l := api.NewTargetAppListLogic(r.Context(), svcCtx)
		resp, err := l.TargetAppList(&req)
		result.HttpResult(r, w, resp, err, true, startTime)
	}
}
