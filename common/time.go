package common

import "time"

func Timestamp() int64 {
	return time.Now().UnixNano() / int64(time.Second)
}
