package models

import (
	"fmt"

	"samh_common_lib/base"
	"utils"
	"utils/config"
	"utils/log"

	"github.com/davecgh/go-spew/spew"
	"github.com/fsnotify/fsnotify"
	"github.com/getsentry/raven-go"
	"github.com/go-redis/redis"
	"github.com/go-xorm/xorm"
)

var (
	Env                string
	Config             *config.Config
	OperationDB        *xorm.Engine
	SamhDB             *xorm.Engine
	RedisClient        *redis.Client
	RedisClusterClient *redis.ClusterClient
)

func Init() {
	filePath := fmt.Sprintf("config/%s.toml", Env)
	config.Init(filePath)
	spew.Printf("config init succ,filepath:%s\n", filePath)
	Config = config.ConfigInstance()
	log.Init(Config.Log_info_item)
	OperationDB = utils.InitDB(Config.DB_arr["operation"])
	SamhDB = utils.InitDB(Config.DB_arr["samh"])
	RedisClient = utils.InitRedisClient(Config.Redis_item)
	RedisClusterClient = utils.InitRedisCluster(Config.Redis_cluster_item)
	err := raven.SetDSN(Config.Sentry_dsn)
	if err != nil {
		log.Panic(err)
	}
}

func NewConfigWatcher() {
	log.Debug(base.NowFunc())

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		panic(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					log.Error(base.NowFuncError())
					return
				}
				if event.Name == fmt.Sprintf("config/%s.toml", Env) && event.Op == fsnotify.Write {
					log.Debug("modified file:", event.Name)
					Init()
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					log.Error(base.NowFuncError())
					return
				}
				log.Error(err)
			}
		}
	}()

	err = watcher.Add("config")
	if err != nil {
		panic(err)
	}
	<-done
}
