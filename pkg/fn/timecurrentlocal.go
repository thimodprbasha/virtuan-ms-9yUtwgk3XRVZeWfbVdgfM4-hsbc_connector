package fn

import (
	"time"
)

func TimeCurrentLocal(location string) (time.Time, error) {
	now := time.Now()
	loc, _ := time.LoadLocation(location)
	localTime := now.In(loc)
	return localTime, nil
}
