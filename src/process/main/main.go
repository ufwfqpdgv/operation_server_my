package main

import (
	"flag"
	"time"

	. "models"
	"process/api"
	"samh_common_lib/base"
	"utils/log"
)

func init() {
	Env = *flag.String("env", "local", "env:local,dev,test,official")
	flag.Parse()
	Init()
	go NewConfigWatcher()
}

func main() {
	ActivityApiTest()
}

func ActivityApiTest() {
	log.Debug("ActivityApiTest")
	start := time.Now()
	request := &ActivityRequest{
		SamhBaseRequest:   base.SamhBaseRequest{Uid: 1, DeviceId: "1"},
		FetchActivityType: FetchActivityTypeCode_All,
		ActivityId:        1,
		ActivityType:      ActivityTypeCode_Vip,
	}
	log.Debug(request)
	rsp, retCode := api.ActivityApi(request)
	log.Debug(retCode, rsp)
	cost := time.Since(start)
	log.Debug(cost)
}
