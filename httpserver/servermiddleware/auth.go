package servermiddleware

import (
	redis "basis/cache"
	"basis/module"
	"basis/util"
	"basis/util/jwt"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	//"strconv"
)

//type BaseAuthReq struct {
//	BaseReq
//	Uid int `form:"uid" json:"uid"  binding:"required"`
//}

//token验证
func BaseAuth() gin.HandlerFunc {
	return func(c *gin.Context) {

		var authReq BaseReq
		//err := c.ShouldBindJSON(&authReq)

		req, err := c.GetRawData()
		if err != nil {
			util.Logger.Errorf("BaseAuth  参数绑定 出错 err: %s ", err.Error())
			c.AbortWithStatusJSON(http.StatusOK, module.ApiResp{ErrorNo: http.StatusForbidden, ErrorMsg: err.Error()})
			return
		}
		//传递参数到下个中间件
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(req)) // 关键点

		err = json.Unmarshal(req, &authReq)

		if err != nil {
			util.Logger.Errorf("BaseAuth  参数绑定 出错 err: %s ", err.Error())
			c.AbortWithStatusJSON(http.StatusOK, module.ApiResp{ErrorNo: http.StatusForbidden, ErrorMsg: err.Error()})
			return
		}
		//token := authReq.Token
		//uid := authReq.Uid

		token := c.GetHeader("token")
		//uid := c.GetHeader("uid")

		if token == "" {
			//权限异常
			c.AbortWithStatusJSON(http.StatusOK, module.ApiResp{
				ErrorNo:  http.StatusForbidden,
				ErrorMsg: http.StatusText(http.StatusForbidden),
			})
			return
		}

		et := jwt.EasyToken{}
		valid, tokenUid, _ := et.ValidateToken(token)

		if !valid {
			c.AbortWithStatusJSON(http.StatusOK, module.ApiResp{
				ErrorNo:  http.StatusForbidden,
				ErrorMsg: "token failed, please login again.",
			})
			return
		}

		//验证token是否存在
		//1,查询redis
		tokenRedis, err := redis.GetRedisDb().Get(util.RedisKeyToken + tokenUid + ":").Result()

		if tokenRedis != token {
			c.AbortWithStatusJSON(http.StatusOK, module.ApiResp{
				ErrorNo:  http.StatusForbidden,
				ErrorMsg: "token failed, please login again.",
			})
			return
		}

		c.Set("uid", tokenUid)

	}

}
