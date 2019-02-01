package main

import (
	"flag"
	"time"

	. "models"
	"process/api"
	"samh_common_lib/base"

	"github.com/davecgh/go-spew/spew"
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
	spew.Dump("ActivityApiTest")
	start := time.Now()
	request := &ActivityRequest{
		SamhBaseRequest:   base.SamhBaseRequest{Uid: 1, DeviceId: "1"},
		FetchActivityType: FetchActivityTypeCode_All,
		ActivityId:        1,
		ActivityType:      ActivityTypeCode_Vip,
	}
	spew.Dump(request)
	rsp, retCode := api.ActivityApi(request)
	spew.Dump(retCode, rsp)
	cost := time.Since(start)
	spew.Dump("cost=", cost)
}
