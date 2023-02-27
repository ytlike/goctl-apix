package utils

import (
	"encoding/json"
	"github.com/pkg/errors"
	"qbq-open-platform/common/errorsEnums"
)

// Struct2json struct 转 json
func Struct2json(s any) (string, error) {
	v, err := json.Marshal(s)
	if err != nil {
		return "", errors.Wrapf(errorsEnums.NewErrCode(errorsEnums.FAILURE), "json序列化失败, %v", err)
	}
	return string(v), err
}

// Json2struct json 转 struct
func Json2struct(s string, i interface{}) error {
	err := json.Unmarshal([]byte(s), i)
	if err != nil {
		return errors.Wrapf(errorsEnums.NewErrCode(errorsEnums.FAILURE), "json反序列化失败, %v", err)
	}
	return nil
}
