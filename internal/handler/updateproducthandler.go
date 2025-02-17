package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"goapi/internal/logic"
	"goapi/internal/svc"
	"goapi/internal/types"
)

func UpdateProductHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateProductRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewUpdateProductLogic(r.Context(), svcCtx)
		resp, err := l.UpdateProduct(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
