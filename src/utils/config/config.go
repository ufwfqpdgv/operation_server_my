package config

import (
	"github.com/jinzhu/configor"
)

var (
	__cfg *Config
)

func Init(configFilePath string) {
	__cfg = &Config{}
	err := configor.Load(__cfg, configFilePath)
	if err != nil {
		panic(err)
	}
}

func ConfigInstance() (cfg *Config) {
	return __cfg
}

type Config struct {
	Sentry_dsn string

	Base_info struct {
		Version string
		Name    string
		Port    int
	}

	Log_info_item   Log_info
	Internal_server map[string]Internal_serverStruct

	DB_arr             map[string]DB
	Redis_item         Redis
	Redis_cluster_item RedisCluster

	Web struct {
		Http_request_timeout int
	}
}

type Log_info struct {
	Level            string
	Encoding         string
	Stdout           bool
	Development_mode bool
	Path_filename    string
	Max_size         int
	Max_backups      int
	Max_age          int
	Compress         bool
}

type Internal_serverStruct struct {
	Url      string
	Time_out int
}

type DB struct {
	Type      string
	Host      string
	Port      int
	User      string
	Password  string
	Db_name   string
	Max_conns int
	Time_out  int
	Log_path  string
	Log_name  string
	//
	Table_name map[string]string
}

type Redis struct {
	Network     string
	Addr        string
	Password    string
	Max_retries int
	Pool_size   int
}

type RedisCluster struct {
	Master_addr_arr []string
	Slave_addr_arr  []string
	Password        string
	Max_retries     int
	Pool_size       int
}
