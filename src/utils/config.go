package utils

import (
	"samh_common_lib/base"

	log "github.com/cihub/seelog"
	"github.com/jinzhu/configor"
)

var (
	__cfg *Config
)

func InitConfig(configFilePath string) {
	log.Debug(base.NowFunc())

	__cfg = &Config{}
	err := configor.Load(__cfg, configFilePath)
	if err != nil {
		log.Error(err)
		panic(err)
	}
	log.Debugf("config init succ,filepath:%s", configFilePath)
}

func ConfigInstance() (cfg *Config) {
	return __cfg
}

type Config struct {
	Sentry_dsn string `required:"-"`

	Base_info struct {
		Version   string `required:"-"`
		Name      string `required:"-"`
		Port      int    `required:"-"`
		Log_level string `required:"-"`
	}

	Internal_server map[string]Internal_serverStruct

	DB_arr             map[string]DB
	Redis_item         Redis
	Redis_cluster_item RedisCluster
}

type Internal_serverStruct struct {
	Url      string `required:"-"`
	Time_out int    `required:"-"`
}

type DB struct {
	Type      string `required:"-"`
	Host      string `required:"-"`
	Port      int    `required:"-"`
	User      string `required:"-"`
	Password  string `required:"-"`
	Db_name   string `required:"-"`
	Max_conns int    `required:"-"`
	Time_out  int    `required:"-"`
	Log_path  string `required:"-"`
	Log_name  string `required:"-"`
	//
	Table_name map[string]string `required:"-"`
}

type Redis struct {
	Network     string `required:"-"`
	Addr        string `required:"-"`
	Password    string `required:"-"`
	Max_retries int    `required:"-"`
	Pool_size   int    `required:"-"`
}

type RedisCluster struct {
	Master_addr_arr []string `required:"-"`
	Slave_addr_arr  []string `required:"-"`
	Password        string   `required:"-"`
	Max_retries     int      `required:"-"`
	Pool_size       int      `required:"-"`
}
