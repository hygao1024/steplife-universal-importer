package pointcalc

import (
	"github.com/kellydunn/golang-geo"
	"math"
	"steplife-universal-importer/internal/model"
)

// Calculate
//
//	@Description: 			基于点之间的距离，计算出中间的点
//	@param previousPoint	前置点
//	@param currentPoint		当前点
//	@param spacing			间距
//	@return []model.Point
func Calculate(previousPoint model.Point, currentPoint model.Point, spacing int) []model.Point {
	p1 := geo.NewPoint(previousPoint.Latitude, previousPoint.Longitude)
	p2 := geo.NewPoint(currentPoint.Latitude, currentPoint.Longitude)
	dist := p1.GreatCircleDistance(p2) // 单位是千米
	// 100米之间生成一个点
	numPoints := int(math.Trunc(dist * 1000 / float64(spacing)))
	// 如果距离太小，则直接返回当前点
	if numPoints == 0 {
		return []model.Point{currentPoint}
	}

	var interpolatedPoints []model.Point
	for i := 0; i < numPoints; i++ {
		alpha := float64(i+1) / float64(numPoints+1)
		// 可能时钟回拨，重置
		if currentPoint.DataTime < previousPoint.DataTime {
			currentPoint.DataTime = previousPoint.DataTime + 1
		}
		newPoint := model.Point{
			DataTime:  previousPoint.DataTime + int64(alpha*(float64(currentPoint.DataTime-previousPoint.DataTime))),
			Altitude:  previousPoint.Altitude + alpha*(currentPoint.Altitude-previousPoint.Altitude),
			Speed:     previousPoint.Speed + alpha*(currentPoint.Speed-previousPoint.Speed),
			Latitude:  previousPoint.Latitude + alpha*(currentPoint.Latitude-previousPoint.Latitude),
			Longitude: previousPoint.Longitude + alpha*(currentPoint.Longitude-previousPoint.Longitude),
		}
		interpolatedPoints = append(interpolatedPoints, newPoint)
	}
	interpolatedPoints = append(interpolatedPoints, currentPoint)
	return interpolatedPoints
}
