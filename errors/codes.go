package errors

import (
	"fmt"
)

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
}

var defaultLang = "zh"

func (this *ErrStruct) Message(lang string) string {
	return FormatValue(this, lang, this.Data)
}

func (this *ErrStruct) Error() string {
	return `{"message":"` + this.Message(defaultLang) +
		`","type":"` + this.Type +
		`","code":"` + fmt.Sprint(this.Code) + `"}`
}

type ErrCodes map[int]*ErrStruct

func (this *ErrCodes) Resiger(code int, err *ErrStruct) {
	(*this)[code] = err
}

func (this *ErrCodes) New(code int, local string, values ...string) *ErrStruct {
	Err, ok := (*this)[code]
	if !ok {
		Err = NotFoundErr
		Err.Data = []string{fmt.Sprint(code)}
	}
	Err.Data = values
	return Err
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
