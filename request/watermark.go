package request

import (
	"github.com/gin-gonic/gin"
	"wTools/param"
	"wTools/response"
)

type Watermark struct {
}

func ParamUrl(ctx *gin.Context) param.ParamUrl {
	var url param.ParamUrl
	err := ctx.Bind(&url)
	if err != nil {
		response.FailWithMessage(err.Error(), ctx)
	}
	return url
}
