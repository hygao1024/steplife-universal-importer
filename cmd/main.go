package main

import (
	"flag"
	"steplife-universal-importer/internal/model"
	"steplife-universal-importer/internal/server"
	"steplife-universal-importer/internal/utils/logx"
	"time"
)

func main() {
	isInterpolate := flag.Bool(
		"isInterpolate", true,
		"是否需要插值点，默认值true，开启。当两点之间距离大于100米，将按照每100米进行插点",
	)
	flag.Parse()

	config := model.Config{
		StartTimestamp: int(time.Now().Unix()),
		IsInterpolate:  *isInterpolate,
	}

	err := server.Run(config)
	if err != nil {
		logx.ErrorF("Run error: %v", err)
		panic(err)
	}
}
