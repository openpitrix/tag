// Copyright 2019 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package config

import (
	"flag"
	"fmt"
	"os"
	"sync"

	"github.com/koding/multiconfig"
	"openpitrix.io/logger"

	"openpitrix.io/tag/pkg/constants"
)

type Config struct {
	Log  LogConfig
	Grpc GrpcConfig

	Mysql struct {
		Host string `default:"tag-db"`
		Port int    `default:"3306"`
		//Host     string `default:"192.168.0.6"`
		//Port     int    `default:"13306"`
		User     string `default:"root"`
		Password string `default:"password"`
		Database string `default:"tag"`
		Disable  bool   `default:"false"`
		//LogMode  bool   `default:"false"`
		LogMode bool `default:"true"`
	}

	App struct {
		//Host string `default:"127.0.0.1"`
		//Port int    `default:"9201"`
		Host string `default:"tag-manager"`
		Port int    `default:"9201"`

		//ApiHost string `default:"127.0.0.1"`
		//ApiPort int    `default:"9200"`
		ApiHost string `default:"tag-manager"`
		ApiPort int    `default:"9200"`

		MaxWorkingNotifications int `default:"5"`
		MaxWorkingTasks         int `default:"5"`
	}

	Websocket struct {
		MessageTypes string `default:"nf,event"`
	}
}

var instance *Config

var once sync.Once

func GetInstance() *Config {
	once.Do(func() {
		instance = &Config{}
	})
	return instance
}

type LogConfig struct {
	Level string `default:"error"` // debug, info, warn, error, fatal
}

type GrpcConfig struct {
	ShowErrorCause bool `default:"false"` // show grpc error cause to frontend
}

func (c *Config) PrintUsage() {
	fmt.Fprintf(os.Stdout, "Usage of %s:\n", os.Args[0])
	flag.PrintDefaults()
	fmt.Fprint(os.Stdout, "\nSupported environment variables:\n")
	e := newLoader(constants.ServiceName)
	e.PrintEnvs(new(Config))
	fmt.Println("")
}

func (c *Config) GetFlagSet() *flag.FlagSet {
	flag.CommandLine.Usage = c.PrintUsage
	return flag.CommandLine
}

func (c *Config) ParseFlag() {
	c.GetFlagSet().Parse(os.Args[1:])
}

func (c *Config) LoadConf() *Config {
	c.ParseFlag()
	config := instance

	m := &multiconfig.DefaultLoader{}
	m.Loader = multiconfig.MultiLoader(newLoader(constants.ServiceName))
	m.Validator = multiconfig.MultiValidator(
		&multiconfig.RequiredValidator{},
	)
	err := m.Load(config)
	if err != nil {
		panic(err)
	}

	logger.Infof(nil, "LoadConf: %+v", config)

	return config
}
