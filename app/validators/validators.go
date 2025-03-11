/*************************************************************************
> File Name: validation.go
> Author: sgs921107
> Mail: 757513128@gmail.com
> Created Time: 2025-03-10 18:06:29 星期一
> Content: This is a desc
*************************************************************************/

package validators

import (
	"regexp"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func ValidatorMobile(level validator.FieldLevel) bool {
	mobile := level.Field().String()
	return regexp.MustCompile(`^1[3-9]\d{9}$`).MatchString(mobile)
}

func RegisterValidators() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("mobile", ValidatorMobile)
	}
}
