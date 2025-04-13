package server

import (
	"errors"
	"fmt"
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

	for fileType, paths := range filePathMap {
		adaptor := parser.CreateAdaptor(fileType)
		if adaptor == nil {
			return errors.New(fmt.Sprintf("不支持的结构解析（%s）", fileType))
		}

		for i, path := range paths {
			logx.InfoF("处理第%d个文件（%s）", i, path)
			content, err := utils.ReadFile(path)
			if err != nil {
				logx.ErrorF("读取文件失败：%s", path)
				return err
			}

			latLngData, err := adaptor.Parse(content)
			if err != nil {
				logx.ErrorF("解析文件失败：%s", path)
				return err
			}

			sl, err := adaptor.Convert2StepLife(config, latLngData)
			if err != nil {
				logx.ErrorF("转换文件失败：%s", path)
				return err
			}

			// 如果文件曾经不存在，则写入CSV文件头
			if !csvExisted {
				err = utils.WriteCSV(csvFilePath, sl.CSVHeader)
				if err != nil {
					logx.ErrorF("写入CSV文件头失败：%s", csvFilePath)
					return err
				}
			}

			err = utils.WriteCSV(csvFilePath, sl.CSVData)
			if err != nil {
				logx.ErrorF("写入CSV文件失败：%s", csvFilePath)
				return err
			}

			// 更新起始时间戳
			config.StartTimestamp += len(sl.CSVData)
			csvExisted = true
		}
	}

	return nil
}
