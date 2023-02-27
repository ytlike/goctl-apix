package api

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"qbq-open-platform/applicationCenter/api/internal/logic/api"
	"qbq-open-platform/applicationCenter/api/internal/svc"
	"qbq-open-platform/applicationCenter/api/internal/types"
)

func TestGetQueryHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.TestGetQueryReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := api.NewTestGetQueryLogic(r.Context(), svcCtx)
		err := l.TestGetQuery(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
