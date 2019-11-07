package handle

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"qp_web_server/module"
	"qp_web_server/util"
)

const version = "qp_server_v1.0.0"

func Version(context *gin.Context) {
	context.JSON(http.StatusOK, module.ApiResp{ErrorNo: util.SuccessCode, ErrorMsg: "", Data: version})
}
