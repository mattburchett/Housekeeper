package util

import "time"

// SubtractedEpoch will take the current time in epoch and subtract (days) worth of seconds from the current epoch time.
func SubtractedEpoch(days int) int64 {
	now := time.Now()
	unix := now.Unix()
	seconds := int64(days * 86400)
	epoch := int64(unix - seconds)
	return epoch
}
