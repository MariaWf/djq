package util

const (
	BeginPage = 1
	DefaultPageSize = 10
	DefaultBeginPage = BeginPage
)

type PageResult struct {
	Objects    []interface{}
	TargetPage int
	PageSize   int
	TotalPage  int
	TotalRows  int
	Rows       int
}

func ComputePageStart(targetPage, pageSize int) int {
	if targetPage > BeginPage {
		return (targetPage - BeginPage) * pageSize
	}
	return 0
}