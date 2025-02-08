/*************************************************************************
> File Name: response.go
> Author: sgs921107
> Mail: 757513128@gmail.com
> Created Time: 2024-11-23 22:23:29 星期六
> Content: This is a desc
*************************************************************************/

package response

type BaseResponse struct {
	Code   int    `json:"code" default:"0"`
	Msg    string `json:"msg" default:"OK"`
	ErrMsg string `json:"errMsg" default:""`
}

func (r *BaseResponse) Error(code Code, errMsg string) {
	r.Code = code.Value()
	r.Msg = code.String()
	r.ErrMsg = errMsg
}

func (r *BaseResponse) Ok() {
	r.Code = OK.Value()
	r.Msg = OK.String()
	r.ErrMsg = ""
}
