package dao

import (
	"fmt"
	"github.com/jhonnli/go-base/initial"
	"github.com/jhonnli/go-base/initial/config"
	"github.com/jinzhu/gorm"
	//_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"os"
)

var (
	db  *gorm.DB
	err error
)

func Init() {
	db, err = gorm.Open("postgres", fmt.Sprintf(config.Config.DB.Dsn, config.Config.DB.Pwd))
	//db, err = gorm.Open("mysql", fmt.Sprintf(config.Config.DB.Dsn, config.Config.DB.Pwd))
	if err != nil {
		initial.Log.Error(fmt.Sprintf("【initPublishDB.NewEngine】ex:%s\n", err.Error()))
		os.Exit(0)
		return
	}
	err = db.DB().Ping()
	if err != nil {
		initial.Log.Error(fmt.Sprintf("【initPublishDB.Ping】ex:%s\n", err.Error()))
		os.Exit(0)
		return
	}
	db.DB().SetMaxIdleConns(config.Config.DB.MaxIdleConn)
	db.DB().SetMaxOpenConns(config.Config.DB.MaxOpenConn)
}
