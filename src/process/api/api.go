package api

import (
	. "models"
	"samh_common_lib/base"
	"utils/log"
)

func ActivityApi(rq *ActivityRequest) (rsp *ActivityResponse, retCode base.SamhResponseCode) {
	log.Debug(base.NowFunc())
	defer log.Debug(base.NowFunc() + " end")

	retCode = base.SamhResponseCode_Succ
	rsp = &ActivityResponse{}
	var err error
	switch rq.FetchActivityType {
	case FetchActivityTypeCode_All:
		err = OperationDB.Select("*").Table("activity").
			Where("use_status=? and status=?",
				base.SamhDataStatusCode_Online, base.SamhDataStatusCode_Normal).
			Find(&rsp.ActivityArr)
	case FetchActivityTypeCode_Id:
		err = OperationDB.Select("*").Table("activity").
			Where("activity_id=? and use_status=? and status=?",
				rq.ActivityId, base.SamhDataStatusCode_Online, base.SamhDataStatusCode_Normal).
			Find(&rsp.ActivityArr)
	case FetchActivityTypeCode_Type:
		err = OperationDB.Select("*").Table("activity").
			Where("type=? and use_status=? and status=?",
				rq.ActivityType, base.SamhDataStatusCode_Online, base.SamhDataStatusCode_Normal).
			Find(&rsp.ActivityArr)
	default:
		retCode = base.SamhResponseCode_Param_Invalid
		return
	}
	if err != nil {
		log.Error(err)
		retCode = base.SamhResponseCode_ServerError
		return
	}
	if len(rsp.ActivityArr) == 0 {
		retCode = base.SamhResponseCode_Data_NotExist
		return
	}

	return
}
