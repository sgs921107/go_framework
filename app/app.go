/*************************************************************************
> File Name: web.go
> Author: sgs921107
> Mail: 757513128@gmail.com
> Created Time: 2024-11-20 20:10:12 星期三
> Content: This is a desc
*************************************************************************/

package app

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sgs921107/glogging"
	v1 "github.com/sgs921107/go_framework/app/v1"
	"github.com/sgs921107/go_framework/app/validators"
	"github.com/sgs921107/go_framework/common"
)

// DebugPrintFunc 自定义DebugPrintRouteFunc
func DebugPrintRouteFunc(httpMethod, absolutePath, handlerName string, nuHandlers int) {
	common.Logger.WithFields(glogging.LogrusFields{
		"method":      httpMethod,
		"path":        absolutePath,
		"handlerName": handlerName,
		"nuHandlers":  nuHandlers,
	}).Debug("Add Route")
}

// DebugPrintFunc 自定义debugPrintFunc
func DebugPrintFunc(format string, values ...interface{}) {
	common.Logger.Debugf(format, values...)
}

// LogFormatter 自定义log formatter
func LogFormatter(param gin.LogFormatterParams) string {
	if param.Latency > time.Minute {
		param.Latency = param.Latency.Truncate(time.Second)
	}
	common.Logger.WithFields(glogging.LogrusFields{
		"ReqTime":  param.TimeStamp.Format(common.Time_LAYOUT),
		"status":   param.StatusCode,
		"latency":  param.Latency,
		"clientIP": param.ClientIP,
		"Method":   param.Method,
		"Path":     param.Path,
		"err":      param.ErrorMessage,
	}).Print("Recv A Request")
	return ""
}

func ListenAndServer(addr string, opts ...gin.OptionFunc) error {
	gin.DebugPrintFunc = DebugPrintFunc
	gin.DebugPrintRouteFunc = DebugPrintRouteFunc
	engin := gin.New(opts...)
	validators.RegisterValidators()
	logger := gin.LoggerWithConfig(gin.LoggerConfig{
		Formatter: LogFormatter,
		Output:    common.Logger.Out,
	})
	engin.Use(logger, gin.Recovery())
	apiGroup := engin.Group("/api")
	v1 := v1.Group{Group: apiGroup}
	v1.Register()
	return engin.Run(addr)
}
