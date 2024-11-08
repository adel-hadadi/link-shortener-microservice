package setup

import (
	"time"
	_ "time/tzdata"
)

func init() {
	loc, _ := time.LoadLocation("Asia/Tehran")
	time.Local = loc
}
