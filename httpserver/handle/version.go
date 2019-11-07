package handle

import (
	"basis/module"
	"basis/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

const version = "qp_server_v1.0.0"

func Version(context *gin.Context) {
	context.JSON(http.StatusOK, module.ApiResp{ErrorNo: util.SuccessCode, ErrorMsg: "", Data: version})
}
