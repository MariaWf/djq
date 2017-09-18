package service

import (
	"testing"
	"mimi/djq/dao/arg"
)

func TestCashCoupon_Find(t *testing.T) {
	serviceCashCoupon := &CashCoupon{}
	argCashCoupon := &arg.CashCoupon{}
	argCashCoupon.ShopIdEqual = "382fdb0c2b7249048f758381574d8762"
	argCashCoupon.NotIncludeHide = true
	argCashCoupon.BeforeExpiryDate = true
	//argCashCoupon.OrderBy = "priority desc"
	argCashCoupon.DisplayNames = []string{"id", "name", "preImage", "expiryDate"}
	cashCouponList, err := Find(serviceCashCoupon, argCashCoupon)
	if err != nil {
		t.Error(err)
	}else{
		t.Log(len(cashCouponList))
	}
}
