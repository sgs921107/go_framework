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
	OK                  = newCode(0, "OK")
	ErrParams           = newCode(400001, "参数错误")
	FailedAuthorize     = newCode(400002, "用户名或密码错误")
	Forbidden           = newCode(400003, "没有权限")
	Unauthorized        = newCode(400004, "用户未登录")
	InvalidToken        = newCode(400005, "无效的token")
	InvalidUserID       = newCode(400006, "不合法的UserID")
	FailedUpdateUser    = newCode(400007, "更新用户信息失败")
	FailedGenToken      = newCode(400008, "生成token失败")
	ErrTokenExpried1    = newCode(400101, "token已过期, 请刷新token")
	ErrTokenExpried2    = newCode(400102, "token已过期, 请重新登录")
	FailedEncryptPasswd = newCode(500001, "加密密码失败")
	FailedInsertUser    = newCode(500002, "插入用户失败")
	FailedQueryUser     = newCode(500003, "查询用户失败")
)
