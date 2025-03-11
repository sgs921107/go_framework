/*************************************************************************
> File Name: testCommon.go
> Author: sgs921107
> Mail: 757513128@gmail.com
> Created Time: 2025-03-10 15:29:53 星期一
> Content: This is a desc
*************************************************************************/

package test

import (
	"testing"

	"github.com/sgs921107/go_framework/common"
)

type str2UintWant struct {
	Val uint
	Err error
}

func TestStr2Uint(t *testing.T) {
	var data = map[string]str2UintWant{
		"0":   {Val: 0, Err: nil},
		"2":   {Val: 2, Err: nil},
		"-10": {Val: 0, Err: common.ErrValue},
		"x":   {Val: 0, Err: common.ErrValue},
	}
	for k, want := range data {
		have, err := common.Str2Uint(k)
		if have != want.Val || err != want.Err {
			t.Errorf("Any2Uint('%v') want (%v, %v) have (%v, %v)", k, want.Val, want.Err, have, err)
		}
	}
}
