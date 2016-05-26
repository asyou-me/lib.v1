package validate

import (
	"fmt"

	"github.com/asyoume/lib/errors"
)

var RoleMap map[string]RoleFunc = map[string]RoleFunc{
	"require":  Require,
	"not_zero": NotZero,
}

func Require(key string, value interface{}) error {
	var isPass = true
	if value == nil {
		return errors.New(errors.NULL, key)
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
		return errors.New(errors.NULL, key)
	}
	return nil
}

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
	fmt.Println(isPass)
	return nil
}
