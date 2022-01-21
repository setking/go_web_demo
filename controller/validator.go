package controller

import (
	"fmt"
	"myApp/models"
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
)

var trans ut.Translator

func InitTrans(locale string) (err error) {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterStructValidation(SignUpParamStructLevelValidation, models.ParamSignUp{})
		//v.RegisterTagNameFunc(func(field reflect.StructField) string {
		//	name := strings.SplitN(field.Tag.Get("json"), "", 2)[0]
		//	if name == "-" {
		//		return ""
		//	}
		//	return name
		//})

		zhT := zh.New()
		enT := en.New()

		uni := ut.New(enT, zhT, enT)
		var ok bool
		trans, ok = uni.GetTranslator(locale)
		if !ok {
			return fmt.Errorf("uni.GetTranslator(%s) failed", locale)
		}
		switch locale {
		case "en":
			err = enTranslations.RegisterDefaultTranslations(v, trans)
		case "zh":
			err = zhTranslations.RegisterDefaultTranslations(v, trans)
		default:
			err = enTranslations.RegisterDefaultTranslations(v, trans)
		}
		return
	}
	return
}

func RemoveTopStruct(field map[string]string) map[string]string {
	res := map[string]string{}
	for field, err := range field {
		res[field[strings.Index(field, ".")+1:]] = err
	}
	return res
}

// SignUpParamStructLevelValidation 自定义SignUpParam结构体校验函数
func SignUpParamStructLevelValidation(sl validator.StructLevel) {
	su := sl.Current().Interface().(models.ParamSignUp)

	if su.Password != su.RePassword {
		// 输出错误提示信息，最后一个参数就是传递的param
		sl.ReportError(su.RePassword, "re_password", "RePassword", "eqfield", "password")
	}
}
