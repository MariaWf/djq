package constant

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
