package xif

func Any(condition bool, trueVal, falseVal interface{}) interface{} {
	if condition {
		return trueVal
	}
	return falseVal
}
func Int(condition bool, trueVal, falseVal int) int {
	return Any(condition, trueVal, falseVal).(int)
}
func Int32(condition bool, trueVal, falseVal int32) int32 {
	return Any(condition, trueVal, falseVal).(int32)
}
func Int64(condition bool, trueVal, falseVal int64) int64 {
	return Any(condition, trueVal, falseVal).(int64)
}
func Uint(condition bool, trueVal, falseVal uint) uint {
	return Any(condition, trueVal, falseVal).(uint)
}
func Uint32(condition bool, trueVal, falseVal uint32) uint32 {
	return Any(condition, trueVal, falseVal).(uint32)
}
func Uint64(condition bool, trueVal, falseVal uint64) uint64 {
	return Any(condition, trueVal, falseVal).(uint64)
}

func Float32(condition bool, trueVal, falseVal float32) float32 {
	return Any(condition, trueVal, falseVal).(float32)
}
func Float64(condition bool, trueVal, falseVal float64) float64 {
	return Any(condition, trueVal, falseVal).(float64)
}

func Bool(condition bool) bool {
	if condition {
		return true
	}
	return false
}

func Str(condition bool, trueVal, falseVal string) string {
	return Any(condition, trueVal, falseVal).(string)
}
