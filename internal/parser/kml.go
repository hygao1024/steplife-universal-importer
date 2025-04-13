package parser

import (
	"encoding/xml"
	"steplife-universal-importer/internal/model"
	"strconv"
	"strings"
)

type KML struct {
	Ovjsn
}

func NewKMLAdaptor() *KML {
	return &KML{}
}

func (this *KML) Parse(content []byte) ([]model.Point, error) {

	var points []model.Point

	decoder := xml.NewDecoder(strings.NewReader(string(content)))
	var inCoordinates bool
	var coordinates string
	for {
		tok, err := decoder.Token()
		if err != nil {
			break
		}

		switch el := tok.(type) {
		case xml.StartElement:
			if el.Name.Local == "coordinates" {
				inCoordinates = true
			}
		case xml.CharData:
			if inCoordinates {
				coordinates = strings.TrimSpace(string(el))
				inCoordinates = false
			}
		}
	}
	coordinatesArr := strings.Split(coordinates, " ")
	for _, point := range coordinatesArr {
		point = strings.TrimSpace(point)
		if point == "" {
			continue
		}
		pointData := strings.Split(point, ",")
		if len(pointData) != 2 {
			continue
		}
		lat, _ := strconv.ParseFloat(pointData[1], 64)
		lng, _ := strconv.ParseFloat(pointData[0], 64)
		points = append(points, model.Point{
			Latitude:  lat,
			Longitude: lng,
		})
	}
	return points, nil
}

func (this *KML) Convert2StepLife(config model.Config, points []model.Point) (*model.StepLife, error) {
	return this.BaseAdaptor.Convert2StepLife(config, points)
}
