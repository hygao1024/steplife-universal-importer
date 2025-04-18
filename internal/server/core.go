package server

import (
	"errors"
	"fmt"
	"path"
	consts "steplife-universal-importer/internal/const"
	"steplife-universal-importer/internal/model"
	"steplife-universal-importer/internal/parser"
	"steplife-universal-importer/internal/utils"
	"steplife-universal-importer/internal/utils/logx"
)

// Run
//
//	@Description: 	执行
//	@return error
func Run(config model.Config) error {
	directory := "./source_data"
	csvFilePath := "./output.csv"

	csvExisted, err := utils.CreateCSVFile(csvFilePath)
	if err != nil {
		return err
	}
	filePathMap, err := utils.GetAllFilePath(directory)
	if err != nil {
		return err
	}

	// 如果文件曾经不存在，则写入CSV文件头
	if !csvExisted {
		sl := model.NewStepLife()
		err = utils.WriteCSV(csvFilePath, sl.CSVHeader)
		if err != nil {
			logx.ErrorF("写入CSV文件头失败：%s", csvFilePath)
			return err
		}
	}

	for fileType, paths := range filePathMap {
		for i, filePath := range paths {
			logx.InfoF("处理第%d个文件（%s）", i, filePath)

			sl, err := parseOne(fileType, filePath, config)
			if err != nil {
				logx.ErrorF("处理第%d个文件（%s）失败：%s", i, filePath, err)
				return err
			}

			err = utils.WriteCSV(csvFilePath, sl.CSVData)
			if err != nil {
				logx.ErrorF("写入CSV文件失败：%s", csvFilePath)
				return err
			}

			// 更新起始时间戳
			config.PathStartTimestamp += int64(len(sl.CSVData))
		}
	}

	return nil
}

func parseOne(fileType, filePath string, config model.Config) (*model.StepLife, error) {

	var adaptor parser.FileAdaptor

	if fileType == consts.FileTypeCommon {
		adaptor = parser.CreateAdaptor(path.Ext(filePath))
	} else if fileType == consts.FileTypeVariFlight {
		// TODO
		logx.ErrorF("飞常准数据后续支持......")
		return nil, nil
	} else {
		logx.ErrorF("不支持的文件类型：%s", fileType)
		return nil, errors.New(fmt.Sprintf("不支持的文件类型：%s", fileType))
	}

	if adaptor == nil {
		return nil, errors.New(fmt.Sprintf("不支持的结构解析（%s）", fileType))
	}

	content, err := utils.ReadFile(filePath)
	if err != nil {
		logx.ErrorF("读取文件失败：%s", filePath)
		return nil, err
	}

	latLngData, err := adaptor.Parse(content)
	if err != nil {
		logx.ErrorF("解析文件失败：%s", filePath)
		return nil, err
	}

	sl, err := adaptor.Convert2StepLife(config, latLngData)
	if err != nil {
		logx.ErrorF("转换文件失败：%s", filePath)
		return nil, err
	}

	return sl, nil
}
