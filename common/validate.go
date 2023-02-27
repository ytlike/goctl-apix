package common

import (
	"context"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
	"qbq-open-platform/common/errorsEnums"
	"time"
)

func Validate(ctx context.Context, data interface{}) error {
	zh := zh.New()
	uni := ut.New(zh, zh)
	trans, _ := uni.GetTranslator("zh")
	validate := validator.New()

	if err := validate.RegisterValidation("checkDatetime", checkDatetime); err != nil {
		return errorsEnums.NewErrCode(errorsEnums.PARAM_VALID_ERROR)
	}

	zhTranslations.RegisterDefaultTranslations(validate, trans)

	//自定义required错误内容
	validate.RegisterTranslation("required", trans, func(ut ut.Translator) error {
		return ut.Add("required", "{0}为必填字段!", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Field())
		return t
	})

	validate.RegisterTranslation("checkDatetime", trans, func(ut ut.Translator) error {
		return ut.Add("checkDatetime", "{0}格式不正确!", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("checkDatetime", fe.Field())
		return t
	})

	err := validate.StructCtx(ctx, data)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		for _, e := range errs {
			return errorsEnums.NewErrCodeMsg(errorsEnums.PARAM_VALID_ERROR, e.Translate(trans))
		}
	}
	return nil
}

func checkDatetime(fl validator.FieldLevel) bool {
	_, err := time.ParseInLocation("2006-01-02 15:04:05", fl.Field().String(), time.Local)
	if err != nil {
		return false
	}
	return true
}
