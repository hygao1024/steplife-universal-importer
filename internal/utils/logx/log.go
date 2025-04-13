package logx

func Error(args ...interface{}) {
	sugar.Error(args...)
}

func ErrorF(template string, args ...interface{}) {
	sugar.Errorf(template, args...)
}

func Info(args ...interface{}) {
	sugar.Info(args...)
}

func InfoF(template string, args ...interface{}) {
	sugar.Infof(template, args...)
}
