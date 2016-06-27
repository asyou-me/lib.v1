package validate

// 入口函数，验证规则
// 参数1：规则列表
func Check(item []*Item) error {
	var err error
	for _, v := range item {
		for _, v2 := range v.Rule {
			role, ok := RuleMap[v2]
			if !ok {
				continue
			}
			err = role(v.Key, v.Value)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
