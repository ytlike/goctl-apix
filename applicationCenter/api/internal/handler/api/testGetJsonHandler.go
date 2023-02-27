package api

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"qbq-open-platform/applicationCenter/api/internal/logic/api"
	"qbq-open-platform/applicationCenter/api/internal/svc"
	"qbq-open-platform/applicationCenter/api/internal/types"
)

func TestGetJsonHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.TestGetJsonReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := api.NewTestGetJsonLogic(r.Context(), svcCtx)
		err := l.TestGetJson(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
