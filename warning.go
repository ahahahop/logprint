package logprint

import (
	"fmt"
	"time"
)

func Warninglog(mes interface{}) {
	fmt.Printf("[warning] %s : %s\n", time.Now().Format("2006-01-02 15:04:05.000"), mes)
}
