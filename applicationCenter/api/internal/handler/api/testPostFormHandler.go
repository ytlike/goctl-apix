package api

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"qbq-open-platform/applicationCenter/api/internal/logic/api"
	"qbq-open-platform/applicationCenter/api/internal/svc"
	"qbq-open-platform/applicationCenter/api/internal/types"
)

func TestPostFormHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.TestPostFormReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := api.NewTestPostFormLogic(r.Context(), svcCtx)
		err := l.TestPostForm(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
