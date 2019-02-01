package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	. "models"
	"samh_common_lib/base"
	local_router "server/router"
	"utils"
	"utils/log"

	"github.com/davecgh/go-spew/spew"
)

func init() {
	Env = *flag.String("env", "local", "env:local,dev,test,official")
	flag.Parse()
	Init()
	go NewConfigWatcher()
	test()
}

func main() {
	s := &http.Server{
		Addr:           fmt.Sprintf(":%v", Config.Base_info.Port),
		Handler:        local_router.InitRouter(),
		ReadTimeout:    2 * time.Second,
		WriteTimeout:   2 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}

func test() {
	// temp()
	// dbTest()
	// redisTest()
	// httpGetTest()
	// httpPostTest()
}

func temp() {
}

func dbTest() {
	activity := &Activity{}
	_, err := OperationDB.Select("*").Get(activity)
	if err != nil {
		panic(err)
	}
	spew.Dump(activity)

	v, err := RedisClient.Do("json.get", "object", ".").String()
	if err != nil {
		panic(err)
	}
	spew.Dump(v)
	o := &Object{}
	err = utils.Json.UnmarshalFromString(v, o)
	if err != nil {
		panic(err)
	}
	spew.Dump(o)
}

func redisTest() {
	v, err := RedisClusterClient.Do("get", "lywob").Result()
	if err != nil {
		panic(err)
	}
	spew.Dump(v)
}

type Object struct {
	Foo string `json:"foo"`
	Ans int    `json:"ans"`
}

func httpGetTest() {
	rq := ShowRequest{
		Uid:           1,
		DeviceId:      "1",
		FetchShowType: FetchShowTypeCode_All,
		ShowType:      4,
	}
	log.Debugf("%+v", rq)
	rsp := &ShowResponse{}
	retCode := utils.HttpGet("http://test.samh.xndm.tech/api/v1/operation/show",
		utils.Struct2Map(rq), rsp, Config.Web.Http_request_timeout)
	log.Debugf("code:%v\nrsp:%v", retCode, spew.Sprintf("%+v", rsp))
	// log.With(zap.string("key","value").Debug(retCode)
}

func httpPostTest() {
	rq := &JoinActivityRequest{
		SamhBaseRequest: base.SamhBaseRequest{
			Uid:      1,
			DeviceId: "1",
		},
		ActivityId:    1,
		RewardRuleId:  1,
		PayType:       1,
		IsPaypal:      0,
		ClientType:    "android",
		ClientVersion: "2.0.4",
	}
	rsp := &JoinActivityResponse{}
	retCode := utils.HttpPost("http://test.samh.xndm.tech/api/v1/operation/join/activity",
		rq, rsp, Config.Web.Http_request_timeout)
	log.Debug(retCode, rsp.Status)
}
