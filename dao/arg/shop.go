package arg

import (
	"mimi/djq/model"
)

type Shop struct {
	IdEqual        string
	IncludeDeleted bool
	NameLike       string
	NameEqual      string
	OrderBy        string
	IdsIn          []string

	NotIncludeHide    bool

	PageSize   int `form:"pageSize" json:"pageSize"`
	TargetPage int `form:"targetPage" json:"targetPage"`

	ShowColumnNames []string

	UpdateObject      interface{}
	UpdateColumnNames []string
}

func (arg *Shop) GetIdsIn() []string {
	return arg.IdsIn
}

func (arg *Shop) SetIdsIn(idsIn []string) {
	arg.IdsIn = idsIn
}

func (arg *Shop) GetUpdateColumnNames() []string {
	return arg.UpdateColumnNames
}

func (arg *Shop) SetUpdateColumnNames(updateColumnNames []string) {
	arg.UpdateColumnNames = updateColumnNames
}

func (arg *Shop) GetOrderBy() string {
	return arg.OrderBy
}

func (arg *Shop) SetOrderBy(orderBy string) {
	arg.OrderBy = orderBy
}

func (arg *Shop) GetIncludeDeleted() bool {
	return arg.IncludeDeleted
}

func (arg *Shop) SetIncludeDeleted(includeDeleted bool) {
	arg.IncludeDeleted = includeDeleted
}

func (arg *Shop) GetIdEqual() string {
	return arg.IdEqual
}

func (arg *Shop) SetIdEqual(idEqual string) {
	arg.IdEqual = idEqual
}

func (arg *Shop) GetUpdateObject() interface{} {
	return arg.UpdateObject
}

func (arg *Shop) SetUpdateObject(updateObject interface{}) {
	arg.UpdateObject = updateObject
}

func (arg *Shop) GetTargetPage() int {
	return arg.TargetPage
}

func (arg *Shop) SetTargetPage(targetPage int) {
	arg.TargetPage = targetPage
}

func (arg *Shop) GetPageSize() int {
	return arg.PageSize
}

func (arg *Shop) SetPageSize(pageSize int) {
	arg.PageSize = pageSize
}

func (arg *Shop) getBaseSql(sql string) string {
	return bindTableName(sql, tableNameShop)
}

func (arg *Shop) getAllColumnNames() []string {
	return ColumnNamesShop
}

func (arg *Shop) getShowColumnNames() []string {
	if arg.ShowColumnNames == nil || len(arg.ShowColumnNames) == 0 {
		return arg.getAllColumnNames()
	}
	s := make([]string, 0, len(arg.ShowColumnNames))
	for _, v := range arg.ShowColumnNames {
		switch v {
		case "id":
			s = append(s, "id")
		case "name":
			s = append(s, "name")
		case "logo":
			s = append(s, "logo")
		case "preImage":
			s = append(s, "pre_image")
		case "totalCashCouponNumber":
			s = append(s, "total_cash_coupon_number")
		case "totalCashCouponPrice":
			s = append(s, "total_cash_coupon_price")
		case "introduction":
			s = append(s, "introduction")
		case "address":
			s = append(s, "address")
		case "priority":
			s = append(s, "priority")
		case "hide":
			s = append(s, "hide")
		}
	}
	if len(s) == 0 {
		return arg.getAllColumnNames()
	}
	return s
}

func (arg *Shop) getColumnNameValues() ([]string, []interface{}) {
	if arg.UpdateColumnNames == nil || len(arg.UpdateColumnNames) == 0 {
		arg.UpdateColumnNames = arg.getAllColumnNames()[1:]
	}
	s := make([]string, 0, len(arg.UpdateColumnNames))
	params := make([]interface{}, 0, 9)
	for _, v := range arg.UpdateColumnNames {
		switch v {
		case "name":
			s = append(s, "name = ?")
			params = append(params, arg.UpdateObject.(*model.Shop).Name)
		case "logo":
			s = append(s, "logo = ?")
			params = append(params, arg.UpdateObject.(*model.Shop).Logo)
		case "preImage":
			s = append(s, "pre_image = ?")
			params = append(params, arg.UpdateObject.(*model.Shop).PreImage)
		case "totalCashCouponNumber":
			s = append(s, "total_cash_coupon_number = ?")
			params = append(params, arg.UpdateObject.(*model.Shop).TotalCashCouponNumber)
		case "totalCashCouponPrice":
			s = append(s, "total_cash_coupon_price = ?")
			params = append(params, arg.UpdateObject.(*model.Shop).TotalCashCouponPrice)
		case "introduction":
			s = append(s, "introduction = ?")
			params = append(params, arg.UpdateObject.(*model.Shop).Introduction)
		case "address":
			s = append(s, "address = ?")
			params = append(params, arg.UpdateObject.(*model.Shop).Address)
		case "priority":
			s = append(s, "priority = ?")
			params = append(params, arg.UpdateObject.(*model.Shop).Priority)
		case "hide":
			s = append(s, "hide = ?")
			params = append(params, arg.UpdateObject.(*model.Shop).Hide)
		}
	}
	return s, params
}

func (arg *Shop) getCountConditions() (string, []interface{}) {
	sql := ""
	params := make([]interface{}, 0, 9)
	if arg.NameLike != "" {
		if sql != "" {
			sql += " and"
		}
		sql += " name like ?"
		params = append(params, "%"+arg.NameLike+"%")
	}
	if arg.NameEqual != "" {
		if sql != "" {
			sql += " and"
		}
		sql += " name = ?"
		params = append(params, arg.NameEqual)
	}
	if arg.NotIncludeHide {
		if sql != "" {
			sql += " and"
		}
		sql += " hide = ?"
		params = append(params, false)
	}
	if len(params) != 0 {
		sql = " where" + sql
	}
	return sql, params
}
