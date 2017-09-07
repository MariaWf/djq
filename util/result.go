package util

const (
	ResultStatusNeedLogin = 2
	ResultStatusSuccess = 1
	ResultStatusFail = 0
)

type ResultVO struct {
	Status int `json:"status"`
	Msg    string `json:"msg"`
	Result interface{} `json:"result"`
}

type PageVO struct {
	TargetPage int
	PageSize   int
	Total      int
	TotalPage  int
	Datas      interface{}
}

func BuildSuccessResult(data interface{}) *ResultVO {
	result := &ResultVO{ResultStatusSuccess, "", data}
	return result
}

func BuildFailResult(msg string) *ResultVO {
	result := &ResultVO{ResultStatusFail, msg, nil}
	return result
}

func BuildNeedLoginResult() *ResultVO {
	result := &ResultVO{ResultStatusNeedLogin, "请登录再进行下一步操作", nil}
	return result
}

func BuildSuccessPageResult(targetPage int, pageSize int, total int, datas interface{}) *ResultVO {
	result := &ResultVO{ResultStatusSuccess, "", nil}
	//pageMap := make(map[string]interface{})
	//pageMap["targetPage"] = targetPage
	//pageMap["pageSize"] = pageSize
	//pageMap["total"] = total
	//pageMap["totalPage"] = totalPage
	//pageMap["datas"] = datas
	result.Result = BuildPageVO(targetPage, pageSize, total, datas)
	return result
}

func BuildPageVO(targetPage int, pageSize int, total int, datas interface{}) *PageVO {
	if targetPage < BeginPage {
		targetPage = BeginPage
	}
	if pageSize < 1 {
		pageSize = DefaultPageSize
	}
	totalPage := total / pageSize
	if total % pageSize > 0 {
		totalPage++
	}
	return &PageVO{targetPage, pageSize, total, totalPage, datas}
}
