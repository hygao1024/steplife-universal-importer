package parser

import (
	"steplife-universal-importer/internal/model"
	xif "steplife-universal-importer/internal/utils/if"
	"steplife-universal-importer/internal/utils/logx"
	"steplife-universal-importer/internal/utils/pointcalc"
)

type FileAdaptor interface {
	//
	// Parse
	//  @Description: 		文件解析
	//  @param content
	//  @return []float64	返回经纬度坐标
	//  @return error
	//
	Parse(content []byte) ([]model.Point, error)

	//
	// Convert2StepLife
	//  @Description: 			将经纬度坐标转换成一生足迹数据结构
	//  @param config 			路径转换配置信息
	//  @param points
	//  @return *model.StepLife
	//  @return error
	Convert2StepLife(config model.Config, points []model.Point) (*model.StepLife, error)
}

type BaseAdaptor struct{}

func (this *BaseAdaptor) Parse(content []byte) ([]model.Point, error) {
	panic("implement me")
}

func (this *BaseAdaptor) Convert2StepLife(config model.Config, points []model.Point) (*model.StepLife, error) {
	previousPoint := model.Point{}

	sl := model.NewStepLife()
	logx.Info("处理经纬度")
	for i, point := range points {

		// 第0个坐标或者不需要插入值，不需要计算中间点，直接写入
		if i == 0 || config.EnableInsertPointStrategy == 0 {
			row := model.NewRow()
			row.DataTime = xif.Int64(row.DataTime == 0, config.PathStartTimestamp, row.DataTime)
			row.Point = point
			sl.AddCSVRow(*row)
			config.PathStartTimestamp++
		} else {
			interpolatedPoints := pointcalc.Calculate(previousPoint, point)
			for _, interpolatedPoint := range interpolatedPoints {
				row := model.NewRow()
				row.Point = interpolatedPoint
				row.DataTime = xif.Int64(
					interpolatedPoint.DataTime == 0,
					config.PathStartTimestamp,
					interpolatedPoint.DataTime,
				)
				sl.AddCSVRow(*row)
				config.PathStartTimestamp++
			}
		}
		previousPoint = point
	}
	logx.InfoF("处理经纬度完成，原始坐标%d个，插点后坐标%d个", len(points), len(sl.CSVData))
	return sl, nil
}

func CreateAdaptor(parserType string) FileAdaptor {
	switch parserType {
	case ".kml":
		return NewKMLAdaptor()
	case ".ovjsn":
		return NewOvjsnAdaptor()
	case ".gpx":
		return NewGpxAdaptor()
	default:
		return nil
	}
}
