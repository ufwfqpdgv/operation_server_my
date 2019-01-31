package api

import (
	. "models"
	process_api "process/api"
	"samh_common_lib/base"

	log "github.com/cihub/seelog"
	"github.com/gin-gonic/gin"
)

func ActivityApi(c *gin.Context) {
	log.Debug(base.NowFunc() + "Start")
	defer base.RecoverFunc(c)

	rq := &ActivityRequest{}
	err := c.ShouldBind(rq)
	log.Debugf(base.NowFunc()+"Request:%+v", *rq)
	if err == nil {
		rsp, retCode := process_api.ActivityApi(rq)
		log.Debugf(base.NowFunc()+"Response:%v,%+v", retCode, *rsp)
		base.SendResponse(c, retCode, rsp)
	} else {
		log.Warn(base.NowFuncError())
		base.SendResponse(c, base.SamhResponseCode_Param_Less, nil)
	}
}
