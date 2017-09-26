package cache

import (
	"mimi/djq/db/redis"
	"time"
	"strings"
	"github.com/pkg/errors"
)

var ErrNameIsEmpty = errors.New("缓存名称为空")

const (
	CacheNameGlobalTotalCashCouponNumber = "globalTotalCashCouponNumber"
	CacheNameGlobalTotalCashCouponPrice = "globalTotalCashCouponPrice"
	CacheNameGlobalTotalCashCouponPriceHide = "globalTotalCashCouponPriceHide"
	CacheNameCashCouponOrderCounting = "cashCouponOrderCounting"

	CacheNameShopRedPackHide = "shopRedPackHide"
	CacheNamePromotionalPartnerRate = "promotionalPartnerRate"

	CacheNamePromotionalPartnerCounting = "promotionalPartnerCounting"

	CacheNameIndexContactWayNumber = "contactIndexWayNumber"
	CacheNameIndexContactWayHide = "contactIndexWayHide"

	CacheNameCheckingRefundingOrder = "checkingRefundingOrder"
	CacheNameCheckingPayingOrder = "checkingPayingOrder"
	CacheNameAgreeingNotUsedRefunding = "agreeingNotUsedRefunding"
	CacheNameCheckingExpiredCashCoupon = "checkingExpiredCashCoupon"
	CacheNameCheckingExpiredPresent = "checkingExpiredPresent"

	CacheNameWithWaterMarkInShopIntroductionImage = "withWaterMarkInShopIntroductionImage"

	cacheHead = "cache:"
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
	return cacheHead + name
}

func FindKeys(pattern string) ([]string, error) {
	conn := redis.Get()
	keys, err := conn.Keys(GetKey(pattern)).Result()
	if err != nil {
		return keys, err
	}
	newKeys := make([]string, len(keys), len(keys))
	for i, key := range keys {
		newKeys[i] = strings.TrimLeft(key, cacheHead)
	}
	return newKeys, nil
}

func Expire(name string, expiration time.Duration) (bool, error) {
	conn := redis.Get()
	return conn.Expire(GetKey(name), expiration).Result()
}

func GetExpire(name string) (time.Duration, error) {
	conn := redis.Get()
	return conn.TTL(GetKey(name)).Result()
}

func Del(name string) (int64, error) {
	conn := redis.Get()
	return conn.Del(GetKey(name)).Result()
}