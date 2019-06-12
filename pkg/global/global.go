// Copyright 2019 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package global

import (
	"os"
	"sync"

	"github.com/google/gops/agent"
	"github.com/jinzhu/gorm"
	"openpitrix.io/logger"

	"openpitrix.io/tag/pkg/config"
	"openpitrix.io/tag/pkg/db"
)

type GlobalCfg struct {
	cfg      *config.Config
	database *gorm.DB
}

var instance *GlobalCfg
var once sync.Once

func GetInstance() *GlobalCfg {
	once.Do(func() {
		instance = newGlobalCfg()
	})
	return instance
}

func newGlobalCfg() *GlobalCfg {
	cfg := config.GetInstance().LoadConf()
	g := &GlobalCfg{cfg: cfg}

	g.setLoggerLevel()
	g.openDatabase()

	if err := agent.Listen(agent.Options{
		ShutdownCleanup: true,
	}); err != nil {
		logger.Criticalf(nil, "Failed to start gops agent")
	}
	return g
}

func (g *GlobalCfg) openDatabase() *GlobalCfg {
	if g.cfg.Mysql.Disable {
		logger.Debugf(nil, "%+s", "Database setting for Mysql.Disable is true.")
		return g
	}
	isSucc := db.GetInstance().InitDataPool()

	if !isSucc {
		logger.Criticalf(nil, "%+s", "Init database pool failure...")
		os.Exit(1)
	}
	logger.Debugf(nil, "%+s", "Init database pool successfully.")

	db := db.GetInstance().GetMysqlDB()
	g.database = db
	logger.Debugf(nil, "%+s", "Set globalcfg database value.")
	return g
}

func (g *GlobalCfg) setLoggerLevel() *GlobalCfg {
	AppLogMode := config.GetInstance().Log.Level
	logger.SetLevelByString(AppLogMode)
	logger.Debugf(nil, "Set app log level to %+s", AppLogMode)
	return g
}

func (g *GlobalCfg) GetDB() *gorm.DB {
	return g.database
}
