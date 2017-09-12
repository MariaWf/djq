package model

type Permission struct {
	Code        string `json:"code" desc:"编码"`
	Name        string `json:"name" desc:"名称"`
	Description string `json:"description" desc:"描述"`
}

func GetPermissionList() []*Permission {
	return []*Permission{
		{"admin_c", "添加管理员信息", "添加管理员信息"},
		{"admin_r", "查看管理员信息", "查看管理员信息"},
		{"admin_u", "更新管理员信息", "更新管理员信息"},
		{"admin_d", "删除管理员信息", "删除管理员信息"},

		{"role_c", "添加角色信息", "添加角色信息"},
		{"role_r", "查看角色信息", "查看角色信息"},
		{"role_u", "更新角色信息", "更新角色信息"},
		{"role_d", "删除角色信息", "删除角色信息"},

		{"advertisement_c", "添加广告信息", "添加广告信息"},
		{"advertisement_r", "查看广告信息", "查看广告信息"},
		{"advertisement_u", "更新广告信息", "更新广告信息"},
		{"advertisement_d", "删除广告信息", "删除广告信息"},

		{"shop_c", "添加商家信息", "添加商家信息"},
		{"shop_r", "查看商家信息", "查看商家信息"},
		{"shop_u", "更新商家信息", "更新商家信息"},
		{"shop_d", "删除商家信息", "删除商家信息"},

		{"shopClassification_c", "添加商家分类信息", "添加商家分类信息"},
		{"shopClassification_r", "查看商家分类信息", "查看商家分类信息"},
		{"shopClassification_u", "更新商家分类信息", "更新商家分类信息"},
		{"shopClassification_d", "删除商家分类信息", "删除商家分类信息"},

		{"cashCoupon_c", "添加代金券信息", "添加代金券信息"},
		{"cashCoupon_r", "查看代金券信息", "查看代金券信息"},
		{"cashCoupon_u", "更新代金券信息", "更新代金券信息"},
		{"cashCoupon_d", "删除代金券信息", "删除代金券信息"},
		
		{"promotionalPartner_c", "添加推广伙伴信息", "添加推广伙伴信息"},
		{"promotionalPartner_r", "查看推广伙伴信息", "查看推广伙伴信息"},
		{"promotionalPartner_u", "更新推广伙伴信息", "更新推广伙伴信息"},
		{"promotionalPartner_d", "删除推广伙伴信息", "删除推广伙伴信息"},

		{"user_c", "添加用户信息", "添加用户信息"},
		{"user_r", "查看用户信息", "查看用户信息"},
		{"user_u", "更新用户信息", "更新用户信息"},
		{"user_d", "删除用户信息", "删除用户信息"},

		{"present_c", "添加礼品信息", "添加礼品信息"},
		{"present_r", "查看礼品信息", "查看礼品信息"},
		{"present_u", "更新礼品信息", "更新礼品信息"},
		{"present_d", "删除礼品信息", "删除礼品信息"},

		{"presentOrder_c", "添加礼品订单信息", "添加礼品订单信息"},
		{"presentOrder_r", "查看礼品订单信息", "查看礼品订单信息"},
		{"presentOrder_u", "更新礼品订单信息", "更新礼品订单信息"},
		{"presentOrder_d", "删除礼品订单信息", "删除礼品订单信息"},

		{"cashCouponOrder_c", "添加代金券订单信息", "添加代金券订单信息"},
		{"cashCouponOrder_r", "查看代金券订单信息", "查看代金券订单信息"},
		{"cashCouponOrder_u", "更新代金券订单信息", "更新代金券订单信息"},
		{"cashCouponOrder_d", "删除代金券订单信息", "删除代金券订单信息"},

		{"refund_c", "添加退款申请信息", "添加退款申请信息"},
		{"refund_r", "查看退款申请信息", "查看退款申请信息"},
		{"refund_u", "更新退款申请信息", "更新退款申请信息"},
		{"refund_d", "删除退款申请信息", "删除退款申请信息"},

		{"refundReason_c", "添加退款理由信息", "添加退款理由信息"},
		{"refundReason_r", "查看退款理由信息", "查看退款理由信息"},
		{"refundReason_u", "更新退款理由信息", "更新退款理由信息"},
		{"refundReason_d", "删除退款理由信息", "删除退款理由信息"}}
}
