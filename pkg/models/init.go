package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"openpitrix.io/tag/pkg/config"
	"openpitrix.io/tag/pkg/db/mysql"
)

var (
	DB *gorm.DB
)

func init() {
	initMysql()
}

func initMysql() {
	mysqlCfg := config.GetInstance().Mysql
	url := fmt.Sprintf(
		"%v:%v@(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local",
		mysqlCfg.User,
		mysqlCfg.Password,
		mysqlCfg.Host,
		mysqlCfg.Port,
		mysqlCfg.Database,
	)
	DB = mysql.Connect(url, mysqlCfg.LogMode)
	DB.Set("gorm:table_options", "ENGINE=MyISAM DEFAULT CHARSET=utf8")
}
