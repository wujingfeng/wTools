package request

import "github.com/gin-gonic/gin"

type BaseParam interface {
	GetParams(ctx *gin.Context) interface{}
}
