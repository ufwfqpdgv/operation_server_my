package utils

import (
	"samh_common_lib/base"

	log "github.com/cihub/seelog"
	"github.com/davecgh/go-spew/spew"
	resty "gopkg.in/resty.v1"
)

func HttpGet(url string, rq map[string]string, rsp interface{}) (retCode base.SamhResponseCode) {
	log.Debug(base.NowFunc())

	log.Debug(spew.Sprintf("Request:%+v", rq))
	resp, err := resty.R().
		SetQueryParams(rq).
		SetResult(rsp).
		Get(url)
	if err != nil {
		retCode = base.SamhResponseCode_ServerError
		log.Error(err, resp)
		return
	}
	retCode = base.SamhResponseCode_Succ
	log.Debug(spew.Sprintf("Response:%+v", rsp))

	return
}

func HttpPost(url string, rq interface{}, rsp interface{}) (retCode base.SamhResponseCode) {
	log.Debug(base.NowFunc())

	log.Debug(spew.Sprintf("Request:%+v", rq))
	resp, err := resty.R().
		SetBody(rq).
		SetResult(rsp).
		Post(url)
	if err != nil {
		retCode = base.SamhResponseCode_ServerError
		log.Error(err, resp)
		return
	}
	spew.Dump(resp)
	spew.Dump(rsp)
	retCode = base.SamhResponseCode_Succ
	log.Debug(spew.Sprintf("Response:%+v", rsp))

	return
}

func HttpPost2(url string, rq interface{}, rsp interface{}) (retCode base.SamhResponseCode) {
	log.Debug(base.NowFunc())

	log.Debug(spew.Sprintf("Request:%+v", rq))
	resp, err := resty.R().
		SetBody(rq).
		SetResult(rsp).
		Post(url)
	if err != nil {
		retCode = base.SamhResponseCode_ServerError
		log.Error(err, resp)
		return
	}
	retCode = base.SamhResponseCode_Succ
	log.Debug(spew.Sprintf("Response:%+v", rsp))

	return
}

/* func HttpRequest(rq *VipRechargeRequest) (rsp *VipRechargeResponse, retCode base.SamhResponseCode) {
	seelog.Debug(base.NowFunc())
	seelog.Debugf(base.NowFunc()+"Request:%+v", *rq)

	rsp = &VipRechargeResponse{}
	addr := ServiceConfig.GetString("VIP_server.Url") + "activity_vip_products/"
	timeOut := time.Duration(ServiceConfig.GetInt("VIP_server.Time_out")) * time.Millisecond
	resty.SetTimeout(timeOut)
	b, err := json.Marshal(rq)
	if err != nil {
		seelog.Error(err.Error())
		retCode = base.SamhResponseCode_Param_Invalid
		return
	}
	resp, err := resty.R().
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetBody(string(b)).
		SetResult(rsp).
		Post(addr)
	if err == nil {
		retCode = base.SamhResponseCode_Succ
		if rsp.Code == base.SamhResponseCode_Succ {
			seelog.Debugf(base.NowFunc()+"Response:%v,%+v", rsp.Code, *rsp)
		} else {
			seelog.Errorf(base.NowFuncError(), *rsp, resp)
		}
		return
	} else {
		seelog.Error(err.Error(), resp)
		retCode = base.SamhResponseCode_ServerError
	}

	return
} */
