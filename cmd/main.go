package main

import (
	"fmt"
	"github.com/pkg/errors"
	"gopkg.in/ini.v1"
	consts "steplife-universal-importer/internal/const"
	"steplife-universal-importer/internal/model"
	"steplife-universal-importer/internal/server"
	xif "steplife-universal-importer/internal/utils/if"
	"steplife-universal-importer/internal/utils/logx"
	timeUtils "steplife-universal-importer/internal/utils/time"
	"time"
)

func initConfig() (model.Config, error) {
	var config model.Config

	cfg, err := ini.Load("config.ini")
	if err != nil {
		logx.ErrorF("Failed to load config: %v", err)
		return config, errors.Wrap(err, "Failed to load config")
	}

	err = cfg.MapTo(&config)
	if err != nil {
		logx.ErrorF("Failed to map config: %v", err)
		return config, errors.Wrap(err, "Failed to map config")
	}

	if config.PathStartTime != "" {
		config.PathStartTimestamp, err = timeUtils.ToTimestamp(config.PathStartTime)
		if err != nil {
			logx.ErrorF("时间解析失败：%s", err)
			return config, errors.Wrap(err, "时间解析失败")
		}
	} else {
		config.PathStartTimestamp = time.Now().Unix()
	}

	config.InsertPointDistance = xif.Int(
		config.InsertPointDistance < consts.MinInsertPointDistance,
		consts.DefaultInsertPointDistance,
		config.InsertPointDistance,
	)

	return config, nil
}

func main() {

	println("\n.---..---..---..---..-.   .-..---..---.   .-..-.-.-..---..----..---. .---..---..---. ")
	println(" \\ \\ `| |'| |- | |-'| |__ | || |- | |- ###| || | | || |-'| || || |-< `| |'| |- | |-< ")
	println("`---' `-' `---'`-'  `----'`-'`-'  `---'   `-'`-'-'-'`-'  `----'`-'`-' `-' `---'`-'`-'\n")

	logx.Info("执行中......")
	config, err := initConfig()
	if err != nil {
		logx.ErrorF("初始化配置失败：%v", err)
		panic(err)
	}

	err = server.Run(config)
	if err != nil {
		logx.ErrorF("Run error: %v", err)
		panic(err)
	}

	var input string
	logx.Info("按任意键结束......")
	_, _ = fmt.Scanln(&input)

}
