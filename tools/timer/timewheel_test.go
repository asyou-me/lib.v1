package timer

import (
	"fmt"
	"testing"
	"time"
)

var stop chan struct{}

func callback(datas []interface{}) {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	fmt.Println(datas)
}

func TestTimeWheel(t *testing.T) {
	timeWheel := NewTimeWhell(1, 3, callback)
	timeWheel.Start()
	timeWheel.Add(11, "11秒测试1")
	timeWheel.Add(11, "11秒测试2")
	timeWheel.Add(12, "12秒测试")
	<-stop
}
