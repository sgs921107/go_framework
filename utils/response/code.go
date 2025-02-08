/*************************************************************************
> File Name: code.go
> Author: sgs921107
> Mail: 757513128@gmail.com
> Created Time: 2024-11-23 22:25:55 星期六
> Content: This is a desc
*************************************************************************/

package response

type Code interface {
	Value() int
	String() string
}

type BaseCode struct {
	value int
	msg   string
}

func (c *BaseCode) Value() int {
	return c.value
}

func (c *BaseCode) String() string {
	return c.msg
}

func newCode(value int, msg string) Code {
	return &BaseCode{value: value, msg: msg}
}

var (
	// succeed
	OK               = newCode(0, "OK")
	ErrBindForm      = newCode(400001, "绑定参数错误")
	ErrParams        = newCode(400002, "参数错误")
	ErrTokenExpried1 = newCode(400003, "token已过期, 请刷新token")
	ErrTokenExpried2 = newCode(400004, "token已过期, 请重新登录")
	ErrTokenInvalid  = newCode(400005, "验证token失败")
	NoUnauthorized   = newCode(400006, "用户未登录")
	FailedAddAuthor  = newCode(500001, "添加作者失败")
	ErrGetAuthors    = newCode(500003, "获取作者列表错误")
)
