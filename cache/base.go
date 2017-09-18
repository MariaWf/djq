package cache

import (
	"mimi/djq/db/redis"
	"time"
	"strings"
	"github.com/pkg/errors"
)

var ErrNameIsEmpty = errors.New("缓存名称为空")

const(
	CacheNameWxpayPayOrderNumberCancel = "wxpay:payOrderNumber:Cancel:"
	CacheNameWxpayErrorPayOrderNumberCancel = "wxpay:error:payOrderNumber:Cancel:"

	CacheNameWxpayPayOrderNumberConfirm = "wxpay:payOrderNumber:Confirm:"
	CacheNameWxpayErrorPayOrderNumberConfirm = "wxpay:error:payOrderNumber:Confirm:"
)

func Get(name string) (string, error) {
	conn := redis.Get()
	if exist, err := conn.Exists(GetKey(name)).Result(); err != nil {
		return "", err
	} else if exist == 0 {
		return "", nil
	}
	return conn.Get(GetKey(name)).Result()
}

func Set(name string, value interface{}, expiration time.Duration) error {
	conn := redis.Get()
	return conn.Set(GetKey(name), value, expiration).Err()
}

func GetKey(name string) string {
	name = strings.TrimSpace(name)
	if name == "" {
		panic(ErrNameIsEmpty)
	}
	return "cache:" + name
}