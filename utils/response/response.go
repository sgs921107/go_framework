/*************************************************************************
> File Name: response.go
> Author: sgs921107
> Mail: 757513128@gmail.com
> Created Time: 2024-11-23 22:23:29 星期六
> Content: This is a desc
*************************************************************************/

package response

type BaseResponse struct {
	Code int    `json:"code" default:"0"`
	Msg  string `json:"errMsg" default:""`
}

func (r *BaseResponse) Error(code Code) {
	r.Code = code.Value()
	r.Msg = code.String()
}

func (r *BaseResponse) Ok(msg string) {
	r.Code = OK.Value()
	if msg == "" {
		r.Msg = OK.String()
	} else {
		r.Msg = msg
	}
}

func NewBaseResponse() *BaseResponse {
	return &BaseResponse{}
}
