package errors

const (
	NULL = iota
)

var defaultCodes = ErrCodes{
	NULL: &ErrStruct{
		Code: 404,
		Format: map[string][]string{
			"zh": []string{
				"%s", "",
			},
			"en": []string{
				"%s", "",
			},
		},
		Level:    "INFO",
		ValueLen: 1,
		Type:     "invalid_request_error",
	},
}
