package errors

import (
	"github.com/asyoume/lib/pulic_type"
)

var Logger pulic_type.Logger = &pulic_type.DefalutLogger{}

func SetLogger(log pulic_type.Logger) {
	Logger = log
}
