package model

type Config struct {
	EnableInsertPointStrategy int    `ini:"enableInsertPointStrategy"`
	InsertPointDistance       int    `ini:"insertPointDistance"`
	PathStartTime             string `ini:"pathStartTime"`
	PathStartTimestamp        int64
}
