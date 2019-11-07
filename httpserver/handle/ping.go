package handle

import (
	"basis/module"
	"basis/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Ping(context *gin.Context) {
	context.JSON(http.StatusOK, module.ApiResp{ErrorNo: util.SuccessCode, ErrorMsg: "pong"})
}
