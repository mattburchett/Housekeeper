package util

import "time"

func SubtractedEpoch(days int) int64 {
	now := time.Now()
	unix := now.Unix()
	seconds := int64(days * 86400)
	epoch := int64(unix - seconds)
	return epoch
}
