// Copyright 2019 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package db

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"openpitrix.io/tag/pkg/config"
)

/*
* MysqlConnPool
* use gorm
 */

var (
	instance *MysqlConnPool
	once     sync.Once
	db       *gorm.DB
	err      error
)

type MysqlConnPool struct{}

func GetInstance() *MysqlConnPool {
	once.Do(func() {
		instance = &MysqlConnPool{}
	})
	return instance
}

/*
* @fuc init connection
 */
func (m *MysqlConnPool) InitDataPool() (isSucc bool) {
	cfg := config.GetInstance()

	var (
		dbCfg            = cfg.Mysql
		connectionString = fmt.Sprintf(
			"%v:%v@(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local",
			dbCfg.User,
			dbCfg.Password,
			dbCfg.Host,
			dbCfg.Port,
			dbCfg.Database,
		)
	)

	db, err = gorm.Open("mysql", connectionString)
	if err != nil {
		log.Print(err)
		return false
	}

	err = db.DB().Ping()

	if err != nil {
		return false
	}

	db.DB().SetMaxIdleConns(100)
	db.DB().SetMaxOpenConns(100)
	db.DB().SetConnMaxLifetime(10 * time.Second)
	db.LogMode(cfg.Mysql.LogMode)

	// table name should be singular
	db.SingularTable(true)

	if err != nil {
		log.Fatal(err)
		return false
	}
	// share between goroutine, no need to close.
	// defer db.Close()
	return true
}

func (m *MysqlConnPool) GetMysqlDB() *gorm.DB {
	return db
}
