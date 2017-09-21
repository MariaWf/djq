package task

//import (
//	"time"
//	"mimi/djq/cache"
//	"log"
//	"mimi/djq/wxpay"
//	"strings"
//	"mimi/djq/service"
//)

//func CloseCanceledOrder() {
//	keys, err := cache.FindKeys(cache.CacheNameWxpayPayOrderNumberCancel + "*")
//	if err != nil {
//		log.Println(err)
//		goto nextCycle
//	}
//	for _, key := range keys {
//		payOrderNumber := strings.TrimLeft( key,cache.CacheNameWxpayPayOrderNumberCancel)
//		paid, err := wxpay.CloseOrderResult(payOrderNumber)
//		if err != nil {
//			log.Println(err)
//			delayCache(key)
//			continue
//		}
//		if paid {
//			serviceRefund := &service.Refund{}
//			serviceRefund.ConfirmByRefundOrderNumber()
//			//value, err := cache.Get(key)
//			//if err != nil {
//			//	log.Println(err)
//			//	delayCache(key)
//			//	continue
//			//}
//			//mimi_todo退款
//			continue
//			delayCache(key)
//		}
//		_, err = cache.Del(key)
//		if err != nil {
//			log.Println(err)
//		}
//	}
//	nextCycle:
//	time.Sleep(time.Minute * 3)
//	go CloseCanceledOrder()
//}
//
//func delayCache(name string) {
//	_, err := cache.Expire(name, time.Hour * 24 * 7)
//	if err != nil {
//		log.Println(err)
//	}
//}