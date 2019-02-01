package utils

import (
	"fmt"
	"os"

	"samh_common_lib/base"
	"utils/config"
	"utils/log"

	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
)

func InitDB(cfg config.DB) (db *xorm.Engine) {
	log.Debug(base.NowFunc())

	db = &xorm.Engine{}
	var connectStr string
	if cfg.Type == "mssql" {
		connectStr = fmt.Sprintf("user id=%s;password=%s;server=%s;port%d;database=%s",
			cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Db_name)
	} else {
		connectStr = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8",
			cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Db_name)
	}

	var err error
	db, err = xorm.NewEngine(cfg.Type, connectStr)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	db.SetMapper(core.GonicMapper{})
	db.ShowSQL(true)
	db.SetMaxIdleConns(cfg.Max_conns)
	db.SetMaxOpenConns(cfg.Max_conns)

	exist, err := pathExists(cfg.Log_path)
	if err != nil {
		panic(err)
	}
	if !exist {
		err = os.Mkdir(cfg.Log_path, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}
	f, err := os.Create(cfg.Log_path + cfg.Log_name)
	if err != nil {
		panic(err)
	}
	db.SetLogger(xorm.NewSimpleLogger(f))

	return
}

// 判断文件夹是否存在
func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
