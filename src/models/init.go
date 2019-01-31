package models

import (
	"fmt"

	"samh_common_lib/base"
	"utils"

	log "github.com/cihub/seelog"
	"github.com/davecgh/go-spew/spew"
	"github.com/fsnotify/fsnotify"
	"github.com/go-redis/redis"
	"github.com/go-xorm/xorm"
	"go.uber.org/zap"
)

var (
	Env                string
	Config             *utils.Config
	Log                *zap.Logger
	OperationDB        *xorm.Engine
	SamhDB             *xorm.Engine
	RedisClient        *redis.Client
	RedisClusterClient *redis.ClusterClient
)

func Init() {
	utils.InitConfig(fmt.Sprintf("config/%s.toml", Env))
	Config = utils.ConfigInstance()
	Log = utils.InitLog(Config.Log_info)
	OperationDB = utils.InitDB(Config.DB_arr["operation"])
	SamhDB = utils.InitDB(Config.DB_arr["samh"])
	RedisClient = utils.InitRedisClient(Config.Redis_item)
	RedisClusterClient = utils.InitRedisCluster(Config.Redis_cluster_item)
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
