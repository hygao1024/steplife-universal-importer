package parser

import (
	"encoding/xml"
	"steplife-universal-importer/internal/model"
	"steplife-universal-importer/internal/utils/logx"
	timeUtils "steplife-universal-importer/internal/utils/time"
	"strings"
)

type GPX struct {
	XMLName xml.Name `xml:"gpx"`
	Tracks  []Track  `xml:"trk"`
}

type Track struct {
	Segments []TrackSegment `xml:"trkseg"`
}

type TrackSegment struct {
	Points []TrackPoint `xml:"trkpt"`
}

type TrackPoint struct {
	Lat   float64 `xml:"lat,attr"`
	Lon   float64 `xml:"lon,attr"`
	Ele   float64 `xml:"ele"`
	Time  string  `xml:"time"`
	Speed float64 `xml:"speed"`
}

type GpxAdaptor struct {
	BaseAdaptor
}

func NewGpxAdaptor() *GpxAdaptor {
	return &GpxAdaptor{}
}

func (this *GpxAdaptor) Parse(content []byte) ([]model.Point, error) {
	var points []model.Point
	var gpx GPX
	decoder := xml.NewDecoder(strings.NewReader(string(content)))
	err := decoder.Decode(&gpx)
	if err != nil {
		return nil, err
	}

	for _, track := range gpx.Tracks {
		for _, segment := range track.Segments {
			for _, pt := range segment.Points {
				timestamp, err := timeUtils.ToTimestamp(pt.Time)
				if err != nil {
					logx.ErrorF("时间解析失败：%s", err)
					return nil, err
				}
				points = append(points, model.Point{
					Latitude:  pt.Lat,
					Longitude: pt.Lon,
					Altitude:  pt.Ele,
					Speed:     pt.Speed,
					DataTime:  timestamp,
				})
			}
		}
	}

	return points, nil
}
