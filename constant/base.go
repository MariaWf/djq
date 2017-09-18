package constant

import "github.com/pkg/errors"

const (
	Split4Permission = ","
	Split4Id = ","
)

var (
	AdminId string
	AdminRoleId string
)

type ApiType int

const (
	ApiTypeMi ApiType = iota
	ApiTypeUi
	ApiTypeSi
	ApiTypeOpen
)

const (
	//待取
	PresentOrderStatusWaiting2Receive int = iota
	//已取
	PresentOrderStatusReceived
)

const (
	//被选中（购物车）
	CashCouponOrderStatusInCart int = iota

	//已购买未使用
	CashCouponOrderStatusPaying

	//已购买未使用
	CashCouponOrderStatusPaidNotUsed

	//已使用
	CashCouponOrderStatusUsed

	//未使用待退款
	CashCouponOrderStatusNotUsedRefunding

	//已使用待退款
	CashCouponOrderStatusUsedRefunding

	//未使用已退款
	CashCouponOrderStatusNotUsedRefunded

	//已使用已退款
	CashCouponOrderStatusUsedRefunded
)

const (
	//未使用待退款
	RefundStatusNotUsedRefunding int = iota

	//未使用退款成功（每天一个固定时间统一自动处理）
	RefundStatusNotUsedRefundSuccess
	//未使用退款撤销
	RefundStatusNotUsedRefundCancel
	//已使用待退款
	// （由平台审批，审批凭证之前说的下载表格找商家签名确认再上传感觉略麻烦，
	// 用户要手机下载文件再打印出来，直接平台打电话给商家确认如何。还有消费者只是部分退款如何处理，
	// 如一次性购买3000，代金券折了300元（折后2700），之后退了其中价值1000元的物品（如果让商家啃了，极端点不是1000而是2700呢），
	// 平台是否该支持部分退款，或者按比例退款而不是整张代金券地退）
	RefundStatusUsedRefunding
	//已使用退款成功
	RefundStatusUsedRefundSuccess
	//已使用退款失败
	RefundStatusUsedRefundFail
	//已使用退款撤销
	RefundStatusUsedRefundCancel
)

var (
	ErrUpload = errors.New("上传失败")
	ErrUploadUnknownType = errors.New("未知文件类型")
	ErrUploadImageSupport = errors.New("只支持jpg;png;gif;jpeg格式")
)

var(
	ErrWxpayConfirmIllegalOrderStatus = errors.New("支付确认，非法订单状态")
	ErrWxpayConfirmTotalFeeNotMatch = errors.New("支付确认，金额不符")

	ErrWxpayCancelIllegalOrderStatus = errors.New("支付取消，非法订单状态")
)

var UploadImageSupport = []string{".jpg", ".png", ".gif", ".jpeg"}