package validate

// 规则定义
type Item struct {
	// 规则的字段，用于出错时增强返回信息
	Key string
	// 需要验证的值
	Value interface{}
	// 需要验证的规则
	Rule []string
}

// 规则函数定义
type RuleFunc func(string, interface{}) error
