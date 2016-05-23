package errors

const (
	MISSING = iota
	INVALID
	SYSTEM_ERR
)

var defaultCodes = ErrCodes{
	SYSTEM_ERR: &ErrStruct{
		Code: 500,
		Format: map[string][]string{
			"zh": []string{
				"系统出现 %s 错误: %s", "未知", "请联系系统维护者",
			},
			"en": []string{
				"system %s error: %s", "unknown", "Please contact the system administrator",
			},
		},
		Level:    "error",
		ValueLen: 2,
		Type:     "system_error",
	}, MISSING: &ErrStruct{
		Code: 404,
		Format: map[string][]string{
			"zh": []string{
				"无法找到 %s: %s", "数据", "请联系系统维护者",
			},
			"en": []string{
				"%s not found: %s", "data", "Please contact the system administrator",
			},
		},
		Level:    "info",
		ValueLen: 2,
		Type:     "miss_error",
	}, INVALID: &ErrStruct{
		Code: 404,
		Format: map[string][]string{
			"zh": []string{
				"无效%s: %s", "数据", "请联系系统维护者",
			},
			"en": []string{
				"invalid parameter %s: %s", "data", "Please contact the system administrator",
			},
		},
		Level:    "info",
		ValueLen: 2,
		Type:     "invalid_request_error",
	},
}
