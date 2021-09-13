package fn

import (
	"time"
)

func TimeSleep(seconds int64) error {
	time.Sleep(time.Duration(seconds) * time.Second)
	return nil
}
