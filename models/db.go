/*************************************************************************
> File Name: db.go
> Author: sgs921107
> Mail: 757513128@gmail.com
> Created Time: 2024-11-21 12:21:02 星期四
> Content: This is a desc
*************************************************************************/

package models

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sgs921107/glogging"
	"github.com/sgs921107/go_framework/common"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	dbInstance *gorm.DB
	dbOnce     sync.Once
)

// DbLogger 自定义gorm的logger
// 使用全局的logger替换gorm的默认logger
type DbLogger struct {
	logger.Config
	Logger *glogging.LogrusLogger
}

// Init 初始化gorm的日志等级
func (l *DbLogger) Init() {
	switch l.Logger.Level {
	case logrus.WarnLevel:
		l.LogMode(logger.Warn)
	case logrus.ErrorLevel:
		l.LogMode(logger.Error)
	case logrus.TraceLevel, logrus.PanicLevel:
		l.LogMode(logger.Silent)
	default:
		l.LogMode(logger.Info)
	}
}

// LogMode 设置gorm的日志等级
func (l *DbLogger) LogMode(level logger.LogLevel) logger.Interface {
	l.LogLevel = level
	return l
}

// Info 普通日志
func (l *DbLogger) Info(ctx context.Context, format string, args ...interface{}) {
	l.Logger.WithContext(ctx).Infof(format, args...)
}

// Info 警告日志
func (l *DbLogger) Warn(ctx context.Context, format string, args ...interface{}) {
	l.Logger.WithContext(ctx).Warnf(format, args...)
}

// Info 错误日志
func (l *DbLogger) Error(ctx context.Context, format string, args ...interface{}) {
	l.Logger.WithContext(ctx).Errorf(format, args...)
}

func (l *DbLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	if l.Logger.Level <= 1 {
		return
	}
	elapsed := time.Since(begin)
	entry := l.Logger.WithContext(ctx).WithField("elapsed", float64(elapsed.Nanoseconds())/1e6)
	switch {
	case err != nil && l.LogLevel >= logger.Silent && (!errors.Is(err, logger.ErrRecordNotFound) || !l.IgnoreRecordNotFoundError):
		sql, rows := fc()
		entry.WithFields(glogging.LogrusFields{
			"sql":  sql,
			"rows": rows,
			"err":  err.Error(),
		}).Trace("Error Execute Sql!")
	case elapsed > l.SlowThreshold && l.SlowThreshold != 0 && l.LogLevel >= logger.Warn:
		sql, rows := fc()
		entry.WithFields(glogging.LogrusFields{
			"sql":           sql,
			"slowThreshold": l.SlowThreshold,
			"rows":          rows,
		}).Warn("Slow Sql")
	case l.LogLevel >= logger.Info:
		sql, rows := fc()
		entry.WithFields(glogging.LogrusFields{
			"sql":  sql,
			"rows": rows,
		}).Info("Execute Sql Succeed!")
	}
}

type MysqlClient struct {
	dsn string
}

func (c *MysqlClient) DSN() string {
	if c.dsn == "" {
		c.dsn = fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
			common.Setting.Mysql.UserName,
			common.Setting.Mysql.Password,
			common.Setting.Mysql.Host,
			common.Setting.Mysql.Port,
			common.Setting.Mysql.DB,
			common.Setting.Mysql.Charset,
		)
	}
	return c.dsn
}

func (c *MysqlClient) Migrator() error {
	err := c.DB().AutoMigrate(&Users{})
	if err != nil {
		common.Logger.WithField("err", err.Error()).Fatal("Failed To Auto Migrate!")

	} else {
		common.Logger.Info("Succeed To Migrate Tables")
	}
	return err
}

// 以单例的形式获取mysql客户端实例
func (c *MysqlClient) DB() *gorm.DB {
	dbOnce.Do(func() {
		dsn := c.DSN()
		entry := common.Logger.WithField("dsn", dsn)
		var err error
		// 实例化自定义的gorm日志管理器
		dbLogger := DbLogger{Logger: common.Logger}
		dbLogger.Init()
		dbInstance, err = gorm.Open(
			mysql.New(mysql.Config{DSN: dsn}), &gorm.Config{Logger: &dbLogger},
		)
		if err != nil {
			entry.WithField("err", err.Error()).Fatal("Failed To Connect Mysql!")
		} else {
			entry.Info("Succeed To Connect Mysql!")
		}
		sqlDB, err := dbInstance.DB()
		// 配置连接池
		if err != nil {
			common.Logger.WithField("err", err.Error()).Error("Failed To Got SQL DB!")
		} else {
			// 最大限制链接的数量
			sqlDB.SetMaxIdleConns(common.Setting.Mysql.MaxIdleConns)
			// 最多打开的链接数
			sqlDB.SetMaxOpenConns(common.Setting.Mysql.MaxOpenConns)
			// 链接的最大保持时间
			sqlDB.SetConnMaxLifetime(time.Duration(common.Setting.Mysql.ConnMaxLifeTime))
			common.Logger.WithFields(glogging.LogrusFields{
				"MaxIdleConns":    common.Setting.Mysql.MaxIdleConns,
				"MaxOpenConns":    common.Setting.Mysql.MaxOpenConns,
				"ConnMaxLifetime": common.Setting.Mysql.ConnMaxLifeTime,
			}).Info("Configured To Mysql Conn Pool")
		}
	})
	return dbInstance
}

// 获取mysql客户端实例
func NewMysqlDB() *gorm.DB {
	client := MysqlClient{}
	return client.DB()
}
