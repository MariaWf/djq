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
	PresentOrderStatusWaiting2Receive int = iota
	PresentOrderStatusReceived
)

var (
	ErrUpload = errors.New("上传失败")
	ErrUploadUnknownType = errors.New("未知文件类型")
	ErrUploadImageSupport = errors.New("只支持jpg;png;gif;jpeg格式")
)

var UploadImageSupport = []string{".jpg", ".png", ".gif", ".jpeg"}