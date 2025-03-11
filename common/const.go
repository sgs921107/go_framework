/*************************************************************************
> File Name: conts.go
> Author: sgs921107
> Mail: 757513128@gmail.com
> Created Time: 2024-11-21 13:07:41 星期四
> Content: This is a desc
*************************************************************************/

package common

import (
	"errors"
	"path"
	"runtime"
)

var (
	_, CUR_FILE, _, _ = runtime.Caller(0)
	CUR_DIR           = path.Dir(CUR_FILE)
	PROJECT_DIR       = path.Dir(CUR_DIR)
	ENV_PATH          = path.Join(PROJECT_DIR, ".env")

	// errors
	ErrValue = errors.New("Value Error")
)

const (
	DATE_LAYOUT = "2006-01-02"
	Time_LAYOUT = "2006-01-02 15:04:05-07:00"
)

type Role uint

const (
	GeneralUser Role = iota
	Merchant
	Administrator
)
