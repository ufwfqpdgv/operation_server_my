package api

import (
	. "models"
	process_api "process/api"
	"samh_common_lib/base"
	"utils/log"

	"github.com/gin-gonic/gin"
)

func ActivityApi(c *gin.Context) {
	defer base.RecoverFunc(c)

	rq := &ActivityRequest{}
	err := c.ShouldBind(rq)
	log.Infof(base.NowFunc()+"Request:%+v", *rq)
	if err == nil {
		rsp, retCode := process_api.ActivityApi(rq)
		log.Infof(base.NowFunc()+"Response:%v,%+v", retCode, *rsp)
		base.SendResponse(c, retCode, rsp)
	} else {
		log.Warn(base.NowFuncError())
		base.SendResponse(c, base.SamhResponseCode_Param_Less, nil)
	}
}
