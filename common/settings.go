/*************************************************************************
> File Name: settings.go
> Author: sgs921107
> Mail: 757513128@gmail.com
> Created Time: 2024-11-20 20:19:43 星期三
> Content: This is a desc
*************************************************************************/

package common

import (
	"github.com/sgs921107/gcommon"
	"sync"
)

var (
	setting = &Settings{}
	once    sync.Once
)

type Settings struct {
	// log
	Log struct {
		Level string `default:"DEBUG"`
	}
	// mysql
	Mysql struct {
		UserName        string
		Password        string
		Host            string `default:"127.0.0.1"`
		Port            int    `default:"3306"`
		DB              string
		Charset         string `default:"utf8mb4"`
		MaxIdleConns    int    `default:"10"`
		MaxOpenConns    int    `default:"60"`
		ConnMaxLifeTime int    `default:"3600"`
	}
	// 雪花ID
	Snow struct {
		NodeID uint16 `lt:"1024"`
	}
	// jwt
	Jwt struct {
		Key string `eq:"32"`
	}
}

func NewSetting() *Settings {
	once.Do(func() {
		gcommon.LoadEnvFile(ENV_PATH, true)
		gcommon.EnvIgnorePrefix()
		gcommon.EnvFill(setting)
	})
	return setting
}
