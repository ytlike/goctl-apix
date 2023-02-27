package utils

import (
	"github.com/satori/go.uuid"
)

func GetUUID() string {
	u := uuid.NewV1()
	return u.String()
}
