package validate

type Item struct {
	Key   string
	Value interface{}
	Rule  []string
}

type RoleFunc func(string, interface{}) error
