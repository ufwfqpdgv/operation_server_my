package main

import (
	"fmt"

	. "models"
	"samh_common_lib/base"
	"utils"
	"utils/log"

	"github.com/davecgh/go-spew/spew"
)

func test() {
	// temp()
	dbTest()
	// redisClientTest()
	// redisClusterTest()
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
	// spew.Dump(activity)
	log.Debug(activity)
}

func redisClientTest() {
	var err error
	o1 := Object{
		Foo: "111",
		Ans: 123,
	}
	s, err := utils.Json.MarshalToString(o1)
	if err != nil {
		panic(err)
	}
	key := Config.Redis_item.Prefix + "object"
	err = RedisClient.Do("json.set", key, ".", s).Err()
	if err != nil {
		panic(err)
	}
	sta := "gg"
	// err = RedisClient.Do("json.set", key, ".foo", `"abcaa"`).Err()
	err = RedisClient.Do("json.set", key, ".foo", fmt.Sprintf(`"%v"`, sta)).Err()

	if err != nil {
		panic(err)
	}

	v, err := RedisClient.Do("json.get", key, ".").String()
	if err != nil {
		panic(err)
	}
	log.Debug(v)
	o := &Object{}
	err = utils.Json.UnmarshalFromString(v, o)
	if err != nil {
		panic(err)
	}
	log.Debugf("%v____%+v", key, o)
}

func redisClusterTest() {
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
