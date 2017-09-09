package model

type Shop struct {
	Id                    string `json:"id" db:"id" desc:"id"`
	Name                  string `json:"name" db:"name" desc:"名称"`
	Logo                  string `json:"logo" db:"logo" desc:"商标"`
	PreImage              string `json:"preImage" db:"pre_image" desc:"预览图"`
	TotalCashCouponNumber int    `json:"totalCashCouponNumber" db:"total_cash_coupon_number" desc:"累计优惠次数"`
	TotalCashCouponPrice  int    `json:"totalCashCouponPrice" db:"total_cash_coupon_price" desc:"累计优惠金额"`
	Introduction          string `json:"introduction" db:"introduction" desc:"介绍"`
	Address               string `json:"address" db:"address" desc:"介绍"`
	Description           string `json:"description" db:"description" desc:"描述"`

	ShopIntroductionImageList []*ShopIntroductionImage `json:"shopIntroductionImageList" desc:"商店介绍图"`
	CashCouponList            []*CashCoupon            `json:"cashCouponList" desc:"代金券"`
}
