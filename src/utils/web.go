package utils

import (
	"time"

	"samh_common_lib/base"

	log "github.com/cihub/seelog"
	"github.com/davecgh/go-spew/spew"
	resty "gopkg.in/resty.v1"
)

func HttpGet(url string, rq map[string]string, rsp interface{}, timeout int) (retCode base.SamhResponseCode) {
	log.Debug(base.NowFunc())

	log.Debug(spew.Sprintf("Request:%+v", rq))
	resty.SetTimeout(time.Duration(timeout) * time.Second)
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

func HttpPost(url string, rq interface{}, rsp interface{}, timeout int) (retCode base.SamhResponseCode) {
	log.Debug(base.NowFunc())

	log.Debug(spew.Sprintf("Request:%+v", rq))
	resty.SetTimeout(time.Duration(timeout) * time.Second)
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

//旧版本的只支持这样form的并把rq先转成json串的样式
func HttpPost2(url string, rq interface{}, rsp interface{}, timeout int) (retCode base.SamhResponseCode) {
	log.Debug(base.NowFunc())

	log.Debug(spew.Sprintf("Request:%+v", rq))
	resty.SetTimeout(time.Duration(timeout) * time.Second)
	b, err := Json.Marshal(rq)
	if err != nil {
		log.Error(err.Error())
		retCode = base.SamhResponseCode_Param_Invalid
		return
	}
	resp, err := resty.R().
		SetHeader("Content-Type", "application/x-www-form-urlencoded").
		SetBody(string(b)).
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
