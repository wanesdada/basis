package handle

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"qp_web_server/module"
	"qp_web_server/util"
)


func Ping(context *gin.Context) {
	context.JSON(http.StatusOK, module.ApiResp{ErrorNo: util.SuccessCode, ErrorMsg: "pong",})
}
