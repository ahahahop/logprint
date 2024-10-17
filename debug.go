// 小写字母,且不要带下划线
package logprint

import (
	"fmt"
	"time"
)

func Debuglog(mes interface{}) {
	fmt.Printf("debug] %s : %s\n", time.Now().Format("2006-01-02 15:04:05.000"), mes)
}
