package models

import (
	"fmt"

	"samh_common_lib/base"
	"utils"

	log "github.com/cihub/seelog"
	"github.com/fsnotify/fsnotify"
	"github.com/go-redis/redis"
	"github.com/go-xorm/xorm"
)

var (
	Env                string
	Config             *utils.Config
	OperationDB        *xorm.Engine
	SamhDB             *xorm.Engine
	RedisClient        *redis.Client
	RedisClusterClient *redis.ClusterClient
)

func Init() {
	utils.InitConfig(fmt.Sprintf("config/%s.toml", Env))
	Config = utils.ConfigInstance()
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
					return
				}
				// log.Debug("event:", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Debug("modified file:", event.Name)
					Init()
				}

			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Error("error:", err)
			}
		}
	}()

	err = watcher.Add("config")
	if err != nil {
		panic(err)
	}
	<-done
}
