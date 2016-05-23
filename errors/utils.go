package errors

import (
	"fmt"
)

func FormatValue(err *ErrStruct, lang string, values []string) string {
	valueLen := len(values)
	Format, defaultValue := selectErrLang(err, lang)
	defaultValueLen := err.ValueLen

	switch {
	case valueLen == defaultValueLen:
		break
	case valueLen == 0 && values[0] == "":
		break
	case valueLen < defaultValueLen:
		values = append(values, defaultValue[valueLen:]...)
		break
	case valueLen > defaultValueLen:
		values = values[:defaultValueLen]
		break
	}
	return fmt.Sprintf(Format, getInterfaceValue(values)...)
}

func selectErrLang(err *ErrStruct, lang string) (string, []string) {
	Format, ok := err.Format[lang]
	if ok {
		return Format[0], Format[1:]
	}
	Format = err.Format["zh"]
	return Format[0], Format[1:]
}

func getInterfaceValue(values []string) []interface{} {
	interfaceValues := make([]interface{}, len(values))
	for i, v := range values {
		interfaceValues[i] = v
	}
	return interfaceValues
}
