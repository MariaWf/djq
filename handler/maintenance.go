package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"strconv"
	"fmt"
	"mimi/djq/cache"
	"log"
	"mimi/djq/util"
	"mimi/djq/task"
)

func PromotionalPartnerActionMaintenanceStatusGet(c *gin.Context) {
	rate, err := cache.Get(cache.CacheNamePromotionalPartnerRate)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrUnknown.Error()))
		return
	}
	if rate == "" {
		rate = "3"
		err = cache.Set(cache.CacheNamePromotionalPartnerRate, rate, 0)
		if err != nil {
			log.Println(err)
			c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrUnknown.Error()))
			return
		}
	}
	counting, err := cache.Get(cache.CacheNamePromotionalPartnerCounting)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrUnknown.Error()))
		return
	}
	if counting == "" {
		counting = "false"
		err = cache.Set(cache.CacheNamePromotionalPartnerCounting, counting, 0)
		if err != nil {
			log.Println(err)
			c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrUnknown.Error()))
			return
		}
	}
	vMap := make(map[string]string)
	vMap["rate"] = rate
	vMap["counting"] = counting
	result := util.BuildSuccessResult(vMap)
	c.JSON(http.StatusOK, result)
}

func PromotionalPartnerActionMaintenanceStatusPost(c *gin.Context) {
	rate := c.PostForm("rate")
	rateN, err := strconv.Atoi(rate)
	if err != nil || rateN < 1 || rateN > 99 {
		if err != nil {
			log.Println(err)
		}
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(fmt.Sprintf("值只能从1到99:%v", rate)))
		return
	}

	err = cache.Set(cache.CacheNamePromotionalPartnerRate, rate, 0)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrUnknown.Error()))
		return
	}
	vMap := make(map[string]string)
	vMap["rate"] = rate
	result := util.BuildSuccessResult(vMap)
	c.JSON(http.StatusOK, result)
}

func PromotionalPartnerActionCountNow(c *gin.Context) {
	defer func(){
		if err:=recover();err!=nil{
			log.Println(err)
			c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrUnknown.Error()))
		}	
	}()
	task.CountForPromotionalPartnerAction()
	result := util.BuildSuccessResult("")
	c.JSON(http.StatusOK, result)
}

func ShopAccountActionMaintenanceStatusGet(c *gin.Context) {
	hide, err := cache.Get(cache.CacheNameShopRedPackHide)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrUnknown.Error()))
		return
	}
	if hide == "" {
		hide = "false"
		err = cache.Set(cache.CacheNameShopRedPackHide, hide, 0)
		if err != nil {
			log.Println(err)
			c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrUnknown.Error()))
			return
		}
	}
	vMap := make(map[string]string)
	vMap["redPackHide"] = hide
	result := util.BuildSuccessResult(vMap)
	c.JSON(http.StatusOK, result)
}

func ShopAccountActionMaintenanceStatusPost(c *gin.Context) {
	hide := c.PostForm("redPackHide")
	if hide != "true"  && hide != "false" {
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(fmt.Sprintf("值只能为true或false:%v", hide)))
		return
	}

	err := cache.Set(cache.CacheNameShopRedPackHide, hide, 0)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrUnknown.Error()))
		return
	}
	vMap := make(map[string]string)
	vMap["redPackHide"] = hide
	result := util.BuildSuccessResult(vMap)
	c.JSON(http.StatusOK, result)
}

func IndexContactWayActionMaintenanceStatusGet(c *gin.Context) {
	hide, err := cache.Get(cache.CacheNameIndexContactWayHide)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrUnknown.Error()))
		return
	}
	number, err := cache.Get(cache.CacheNameIndexContactWayNumber)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrUnknown.Error()))
		return
	}

	if !util.MatchMobile(number) {
		number = ""
		err = cache.Set(cache.CacheNameIndexContactWayNumber, number, 0)
		if err != nil {
			log.Println(err)
			c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrUnknown.Error()))
			return
		}
	}
	if hide == "" {
		hide = "false"
		err = cache.Set(cache.CacheNameIndexContactWayHide, hide, 0)
		if err != nil {
			log.Println(err)
			c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrUnknown.Error()))
			return
		}
	}
	vMap := make(map[string]string)
	vMap["indexContactWayHide"] = hide
	vMap["indexContactWayNumber"] = number
	result := util.BuildSuccessResult(vMap)
	c.JSON(http.StatusOK, result)
}

func IndexContactWayActionMaintenanceStatusPost(c *gin.Context) {
	hide := c.PostForm("indexContactWayHide")
	number := c.PostForm("indexContactWayNumber")
	if hide != "true"  && hide != "false" {
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(fmt.Sprintf("值只能为true或false:%v", hide)))
		return
	}
	if !util.MatchMobile(number) {
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(fmt.Sprintf(util.ErrMobileFormat.Error()+":%v", number)))
		return
	}

	err := cache.Set(cache.CacheNameIndexContactWayNumber, number, 0)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrUnknown.Error()))
		return
	}

	err = cache.Set(cache.CacheNameIndexContactWayHide, hide, 0)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrUnknown.Error()))
		return
	}
	vMap := make(map[string]string)
	vMap["indexContactWayHide"] = hide
	vMap["indexContactWayNumber"] = number
	result := util.BuildSuccessResult(vMap)
	c.JSON(http.StatusOK, result)
}

func CashCouponOrderActionMaintenanceStatusGet(c *gin.Context) {
	hide, err := cache.Get(cache.CacheNameGlobalTotalCashCouponPriceHide)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrUnknown.Error()))
		return
	}
	counting, err := cache.Get(cache.CacheNameCashCouponOrderCounting)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrUnknown.Error()))
		return
	}

	if counting=="" {
		counting = "false"
		err = cache.Set(cache.CacheNameCashCouponOrderCounting, counting, 0)
		if err != nil {
			log.Println(err)
			c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrUnknown.Error()))
			return
		}
	}
	if hide == "" {
		hide = "false"
		err = cache.Set(cache.CacheNameGlobalTotalCashCouponPriceHide, hide, 0)
		if err != nil {
			log.Println(err)
			c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrUnknown.Error()))
			return
		}
	}
	vMap := make(map[string]string)
	vMap["cashCouponOrderHide"] = hide
	vMap["cashCouponOrderCounting"] = counting
	result := util.BuildSuccessResult(vMap)
	c.JSON(http.StatusOK, result)
}

func CashCouponOrderActionMaintenanceStatusPost(c *gin.Context) {
	hide := c.PostForm("cashCouponOrderHide")
	if hide != "true"  && hide != "false" {
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(fmt.Sprintf("值只能为true或false:%v", hide)))
		return
	}

	err := cache.Set(cache.CacheNameGlobalTotalCashCouponPriceHide, hide, 0)
	if err != nil {
		log.Println(err)
		c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrUnknown.Error()))
		return
	}
	vMap := make(map[string]string)
	vMap["cashCouponOrderHide"] = hide
	result := util.BuildSuccessResult(vMap)
	c.JSON(http.StatusOK, result)
}

func CashCouponOrderActionCountNow(c *gin.Context) {
	defer func(){
		if err:=recover();err!=nil{
			log.Println(err)
			c.AbortWithStatusJSON(http.StatusOK, util.BuildFailResult(ErrUnknown.Error()))
		}
	}()
	task.CountCashCouponAction()
	result := util.BuildSuccessResult("")
	c.JSON(http.StatusOK, result)
}