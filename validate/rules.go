package validate

import (
	"fmt"

	"github.com/asyou-me/lib.v1/errors"
)

var RuleMap map[string]RuleFunc = map[string]RuleFunc{
	"require":  Require,
	"not_zero": NotZero,
}

// 是否必须
func Require(key string, value interface{}) error {
	var isPass = true
	if value == nil {
		return errors.New(errors.NULL, key+"")
	}

	switch v := value.(type) {
	case int8:
		break
	case int16:
		break
	case int32:
		break
	case int64:
		break
	case uint8:
		break
	case uint16:
		break
	case uint32:
		break
	case uint64:
		break
	case float32:
		break
	case float64:
		break
	case string:
		isPass = len(v) > 0
		break
	default:
		if fmt.Sprint(v) == "<nil>" {
			isPass = false
		}
		break
	}
	if !isPass {
		return errors.New(errors.NULL, key+" 为必填参数")
	}
	return nil
}

// 不能为0
func NotZero(key string, value interface{}) error {
	var isPass bool
	switch v := value.(type) {
	case int8:
		isPass = v > 0
		break
	case int16:
		isPass = v > 0
		break
	case int32:
		isPass = v > 0
		break
	case int64:
		isPass = v > 0
		break
	case uint8:
		isPass = v > 0
		break
	case uint16:
		isPass = v > 0
		break
	case uint32:
		isPass = v > 0
		break
	case uint64:
		isPass = v > 0
		break
	case float32:
		isPass = v > 0
		break
	case float64:
		isPass = v > 0
		break
	default:
		isPass = false
		break
	}
	if !isPass {
		return errors.New(errors.NULL, key+" 必须为数字,且不能为0")
	}
	return nil
}
