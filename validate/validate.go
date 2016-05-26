package validate

import (
//"fmt"
//"github.com/asaskevich/govalidator"
)

func Check(item []*Item) error {
	var err error
	for _, v := range item {
		for _, v2 := range v.Rule {
			role, ok := RoleMap[v2]
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
