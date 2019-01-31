package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	. "models"
	local_router "server/router"
	// "utils"
	// "github.com/davecgh/go-spew/spew"
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
	// activity := &Activity{}
	// _, err := OperationDB.Select("*").Get(activity)
	// if err != nil {
	// panic(err)
	// }
	// spew.Dump(activity)

	// v, err := RedisClient.Do("json.get", "object", ".").String()
	// if err != nil {
	// panic(err)
	// }
	// spew.Dump(v)
	// o := &Object{}
	// err = utils.Json.UnmarshalFromString(v, o)
	// if err != nil {
	// panic(err)
	// }
	// spew.Dump(o)

	// v, err := RedisClusterClient.Do("get", "lywob").Result()
	// if err != nil {
	// panic(err)
	// }
	// spew.Dump(v)
}

type Object struct {
	Foo string `json:"foo"`
	Ans int    `json:"ans"`
}
