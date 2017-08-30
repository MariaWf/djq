package handler

import "mimi/uservice/util"

const (
	ResultStatusSuccess = 1
	ResultStatusFail = 0
)

type ResultVO struct {
	Status int `json:"status"`
	Msg    string `json:"msg"`
	Result interface{} `json:"result"`
}

func BuildSuccessResult(data interface{}) *ResultVO {
	result := &ResultVO{ResultStatusSuccess, "", data}
	return result
}

func BuildFailResult(msg string) *ResultVO {
	result := &ResultVO{ResultStatusFail, msg, nil}
	return result
}

func BuildSuccessPageResult(targetPage int, pageSize int, total int, datas interface{}) *ResultVO {
	result := &ResultVO{ResultStatusSuccess, "", nil}
	pageMap := make(map[string]interface{})
	if targetPage < util.BeginPage {
		targetPage = util.BeginPage
	}
	if pageSize < 1 {
		pageSize = util.DefaultPageSize
	}
	totalPage := total / pageSize
	if total % pageSize > 0 {
		totalPage++
	}
	pageMap["targetPage"] = targetPage
	pageMap["pageSize"] = pageSize
	pageMap["total"] = total
	pageMap["totalPage"] = totalPage
	pageMap["datas"] = datas
	result.Result = pageMap
	return result
}