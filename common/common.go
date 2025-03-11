/*************************************************************************
> File Name: common.go
> Author: sgs921107
> Mail: 757513128@gmail.com
> Created Time: 2024-11-21 13:06:20 星期四
> Content: This is a desc
*************************************************************************/

package common

import (
	"github.com/sgs921107/glogging"
	"strconv"
)

var (
	Setting = NewSetting()
	Logger  = glogging.NewLogrusLogging(glogging.Options{}).GetLogger()
)

func Str2Uint(s string) (uint, error) {
	i, err := strconv.Atoi(s)
	if err != nil {
		return 0, ErrValue
	}
	if i < 0 {
		return 0, ErrValue
	}
	return uint(i), nil
}
