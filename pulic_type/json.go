package pulic_type

import (
	"database/sql/driver"
	"errors"
	"fmt"
	"time"
)

// 可为null字符串
type StringNull struct {
	/* 0 为初始化值 1为null 2为正常字符串*/
	status int8
	value  string
}

// 序列化时调用
func (t *StringNull) MarshalJSON() ([]byte, error) {
	if t.value == "" {
		return []byte("null"), nil
	} else {
		return []byte(`"` + t.value + `"`), nil
	}
}

// 反序列化时调用
func (t *StringNull) UnmarshalJSON(value []byte) error {
	var str = string(value)
	if str == "" || str == "null" {
		t.status = 1
		return nil
	}
	t.Set(str)
	return nil
}

// 获取值
func (t *StringNull) Get() string {
	return t.value
}

// 设定值
func (t *StringNull) Set(value string) {
	if value != "" {
		t.status = 2
	} else {
		t.status = 1
	}
	t.value = value
}

// 获取当前值的状态
func (t *StringNull) Status() int8 {
	return t.status
}

// sql 数据库插入时转换值
func (t StringNull) Value() (driver.Value, error) {
	return driver.Value(t.value), nil
}

// 输出格式可变的时间类型
type StampTime struct {
	/*值为1为格式化为数字 值为0格式化为StampTime*/
	JosnType int8
	Data     int64
}

// json 序列化时调用
func (t *StampTime) MarshalJSON() ([]byte, error) {
	if t.JosnType == 0 {
		return []byte(fmt.Sprint(t.Data)), nil
	} else {
		return []byte(`"` + time.Unix(t.Data, 0).Format("2006-01-02 15:04:05") + `"`), nil
	}
}

// 插入 sql 数据库时调用
func (t *StampTime) Value() (driver.Value, error) {
	return driver.Value(time.Unix(t.Data, 0).Format("2006-01-02 15:04:05")), nil
}

// 可为null字符串
type BoolNull struct {
	/* 0 为初始化值 1为null 2为正常bool*/
	status int8
	value  bool
}

// 反序列化时调用
func (t *BoolNull) UnmarshalJSON(value []byte) error {
	var str = string(value)

	switch str {
	case "null":
		t.status = 1
		break
	case "true":
		t.status = 2
		t.value = true
		break
	case "false":
		t.status = 2
		t.value = false
		break
	default:
		return errors.New("json: one value must be bool，not " + str)
	}

	return nil
}

// 序列化时调用
func (t *BoolNull) MarshalJSON() ([]byte, error) {
	if t.status == 2 {
		return []byte(fmt.Sprint(t.value)), nil
	} else {
		return []byte("null"), nil
	}
}

// 获取值
func (t *BoolNull) Get() bool {
	return t.value
}

// 设定值
func (t *BoolNull) Set(value bool) {
	t.status = 2
	t.value = value
}

// 获取当前值的状态
func (t *BoolNull) Status() int8 {
	return t.status
}

// 获取当前值的状态
func (t *BoolNull) Null() {
	t.status = 1
}

// sql 数据库插入时转换值
func (t *BoolNull) Value() (driver.Value, error) {
	return driver.Value(t.value), nil
}
