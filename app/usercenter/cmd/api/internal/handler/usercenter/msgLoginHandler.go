package usercenter

import (
	"net/http"

	"github/community-online/common/result"

	"github/community-online/app/usercenter/cmd/api/internal/logic/usercenter"
	"github/community-online/app/usercenter/cmd/api/internal/svc"
	"github/community-online/app/usercenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func MsgLoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.MsgLoginReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := usercenter.NewMsgLoginLogic(r.Context(), svcCtx)
		resp, err := l.MsgLogin(&req)
		result.HttpResult(r, w, resp, err)
	}
}
