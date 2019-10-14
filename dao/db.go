package dao

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jhonnli/go-base/initial/config"
	"github.com/jinzhu/gorm"
	"log"
	"os"
)

var (
	db  *gorm.DB
	err error
)

func Init() {
	db, err = gorm.Open("mysql", fmt.Sprintf(config.Config.DB.Dsn, config.Config.DB.Pwd))
	if err != nil {
		log.Fatalf("【initPublishDB.NewEngine】ex:%s\n", err.Error())
		os.Exit(0)
		return
	}
	err = db.Ping()
	if err != nil {
		log.Fatalf("【initPublishDB.Ping】ex:%s\n", err.Error())
		os.Exit(0)
		return
	}
	db.SetMaxIdleConns(config.Config.DB.PublishDb.MaxIdleConn)
	db.SetMaxOpenConns(config.Config.DB.PublishDb.MaxOpenConn)
}