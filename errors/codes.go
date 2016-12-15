package errors

import (
	"fmt"
)

// ErrStruct 错误类型
type ErrStruct struct {
	// status code
	Code int
	// 错误等级
	Level string
	// 错误信息模板
	Format map[string][]string
	// 错误信息
	Data []string
	// 验参信息模板
	ErrCode int
	// 验参信息模板
	ErrData *map[string]string
	// 错误长度
	ValueLen int
	// 文字错误类型
	Type string

	// Log 信息
	Log *LogStruct
}

// LogStruct 日志类型
type LogStruct struct {
	Local  string
	LogMsg string
	Line   int
}

// 默认输出语言
var defaultLang = "zh"

// LogMessage 设定消息
func (e *ErrStruct) LogMessage(msg string) {
	if e.Log == nil {
		e.Log = &LogStruct{}
	}
	e.Log.LogMsg = msg
}

// Message 通过语言类型获取消息
func (e *ErrStruct) Message(lang string) string {
	return FormatValue(e, lang, e.Data)
}

// 获取错误信息
func (e *ErrStruct) Error() string {
	return `{"message":"` + e.Message(defaultLang) +
		`","type":"` + e.Type +
		`","code":"` + fmt.Sprint(e.Code) + `"}`
}

// ErrCodes 错误类型集合
type ErrCodes map[int]*ErrStruct

// Resiger 注册一个错误类型
func (e *ErrCodes) Resiger(code int, err *ErrStruct) {
	(*e)[code] = err
}

// New 创建一个错误类型
func (e *ErrCodes) New(code int, values ...string) *ErrStruct {
	// Err 为指针 不能直接使用需要复制部分  多线程才能安全
	var Err = ErrStruct{}
	ErrP, ok := (*e)[code]
	if !ok {
		Err = *NotFoundErr
		Err.Data = []string{fmt.Sprint(code)}
	} else {
		Err = *ErrP
	}

	Err.Data = values

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
