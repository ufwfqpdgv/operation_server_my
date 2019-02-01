package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"models"
	"server/router"
)

func init() {
	models.Env = *flag.String("env", "local", "env:local,dev,test,official")
	flag.Parse()
	models.Init()
	go models.NewConfigWatcher()
	test()
}

func main() {
	s := &http.Server{
		Addr:           fmt.Sprintf(":%v", models.Config.Base_info.Port),
		Handler:        router.InitRouter(),
		ReadTimeout:    2 * time.Second,
		WriteTimeout:   2 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
