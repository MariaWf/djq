package util

import (
	"strings"
	"github.com/satori/go.uuid"
)

func BuildUUID() string {
	return strings.Replace(uuid.NewV4().String(), "-", "", -1)
}
