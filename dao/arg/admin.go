package arg

import (
	"mimi/djq/model"
)

type Admin struct {
	IdEqual        string
	IncludeDeleted bool
	NameLike       string
	NameEqual      string
	OrderBy        string
	IdsIn          []string
	IdsNotIn       []string
	LockedEqual    string `form:"locked" json:"locked"`
	LockedOnly     bool
	UnlockedOnly   bool
	KeywordLike    string `form:"keyword" json:"keyword"`

	PasswordEqual string

	PageSize   int `form:"pageSize" json:"pageSize"`
	TargetPage int `form:"targetPage" json:"targetPage"`

	DisplayNames []string

	UpdateObject interface{}
	UpdateNames  []string
}

func (arg *Admin) GetDisplayNames() []string {
	return arg.DisplayNames
}

func (arg *Admin) GetModelInstance() model.BaseModelInterface {
	return &model.Admin{}
}

func (arg *Admin) GetIdsIn() []string {
	return arg.IdsIn
}

func (arg *Admin) SetIdsIn(idsIn []string) {
	arg.IdsIn = idsIn
}

func (arg *Admin) GetUpdateNames() []string {
	return arg.UpdateNames
}

func (arg *Admin) SetUpdateNames(updateNames []string) {
	arg.UpdateNames = updateNames
}

func (arg *Admin) GetOrderBy() string {
	return arg.OrderBy
}

func (arg *Admin) SetOrderBy(orderBy string) {
	arg.OrderBy = orderBy
}

func (arg *Admin) GetIncludeDeleted() bool {
	return arg.IncludeDeleted
}

func (arg *Admin) SetIncludeDeleted(includeDeleted bool) {
	arg.IncludeDeleted = includeDeleted
}

func (arg *Admin) GetIdEqual() string {
	return arg.IdEqual
}

func (arg *Admin) SetIdEqual(idEqual string) {
	arg.IdEqual = idEqual
}

func (arg *Admin) GetUpdateObject() interface{} {
	return arg.UpdateObject
}

func (arg *Admin) SetUpdateObject(updateObject interface{}) {
	arg.UpdateObject = updateObject
}

func (arg *Admin) GetTargetPage() int {
	return arg.TargetPage
}

func (arg *Admin) SetTargetPage(targetPage int) {
	arg.TargetPage = targetPage
}

func (arg *Admin) GetPageSize() int {
	return arg.PageSize
}

func (arg *Admin) SetPageSize(pageSize int) {
	arg.PageSize = pageSize
}

func (arg *Admin) getCountConditions() (string, []interface{}) {
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
	if arg.KeywordLike != "" {
		if sql != "" {
			sql += " and"
		}
		sql += " (name like ? or mobile like ?)"
		params = append(params, "%"+arg.KeywordLike+"%")
		params = append(params, "%"+arg.KeywordLike+"%")
	}
	if arg.PasswordEqual != "" {
		if sql != "" {
			sql += " and"
		}
		sql += " password = ?"
		params = append(params, arg.PasswordEqual)
	}
	if arg.LockedOnly || arg.UnlockedOnly {
		if sql != "" {
			sql += " and"
		}
		sql += " locked = ?"
		params = append(params, arg.LockedOnly)
	}
	if arg.LockedEqual != "" {
		if sql != "" {
			sql += " and"
		}
		sql += " locked = ?"
		params = append(params, arg.LockedEqual == "true")
	}
	if arg.IdsNotIn != nil && len(arg.IdsNotIn) != 0 {
		if sql != "" {
			sql += " and"
		}
		sql += " id not in ("
		for i, id := range arg.IdsNotIn {
			if i != 0 {
				sql += ","
			}
			sql += "?"
			params = append(params, id)
		}
		sql += ")"
	}
	if len(params) != 0 {
		sql = " where" + sql
	}
	return sql, params
}
