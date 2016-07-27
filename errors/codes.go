package errors

import (
	"fmt"
)

// 错误类型
type ErrStruct struct {
	// status code
	Code int
	// 错误等级
	Level string
	// 错误信息模板
	Format map[string][]string
	// 错误信息
	Data []string
	// 错误长度
	ValueLen int
	// 文字错误类型
	Type string
	// Log 信息
	Local string
	// Log 信息
	LogMsg string
}

// 默认输出语言
var defaultLang = "zh"

// 通过语言类型获取消息
func (this *ErrStruct) Message(lang string) string {
	return FormatValue(this, lang, this.Data)
}

// 获取错误信息
func (this *ErrStruct) Error() string {
	return `{"message":"` + this.Message(defaultLang) +
		`","type":"` + this.Type +
		`","code":"` + fmt.Sprint(this.Code) + `"}`
}

// 错误类型集合
type ErrCodes map[int]*ErrStruct

// 注册一个错误类型
func (this *ErrCodes) Resiger(code int, err *ErrStruct) {
	(*this)[code] = err
}

// 创建一个错误类型
func (this *ErrCodes) New(code int, local string, values ...string) *ErrStruct {
	// Err 为指针 不能直接使用需要复制部分  多线程才能安全
	var Err ErrStruct = ErrStruct{}
	ErrP, ok := (*this)[code]
	if !ok {
		Err = *NotFoundErr
		Err.Data = []string{fmt.Sprint(code)}
	} else {
		Err = *ErrP
	}

	Err.Data = values
	if local != "" {
		Err.Local = local
	}
	return &Err
}

var NotFoundErr = &ErrStruct{
	Code: 500,
	Format: map[string][]string{
		"zh": []string{
			"抛出了无效的错误码 :%s", "未知",
		},
		"en": []string{
			"Thrown invalid error code :%s", "unknown",
		},
	},
	Level:    "error",
	ValueLen: 1,
	Type:     "system_err",
}
