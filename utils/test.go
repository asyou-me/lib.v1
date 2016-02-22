package utils

import (
	"fmt"
	"time"
)

// example
// ctime:=Time_Time(func() {
//   for i := 0; i < 10000; i++ {
//      fmt.Println(1111)
//   }
// })

func Time_Time(func_test func()) int64 {

	t1 := time.Now().UnixNano()
	func_test()
	t2 := time.Now().UnixNano()

	return t2 - t1
}
