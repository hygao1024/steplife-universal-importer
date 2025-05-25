package parser

import (
	"encoding/json"
	"github.com/tidwall/gjson"
	"steplife-universal-importer/internal/model"
)

type Ovjsn struct {
	BaseAdaptor
}

func NewOvjsnAdaptor() *Ovjsn {
	return &Ovjsn{}
}

func (this *Ovjsn) Parse(content []byte) ([]model.Point, error) {
	var points []model.Point
	var latLngArr []float64

	// 检查是否有 BOM
	if len(content) >= 3 && content[0] == 0xEF && content[1] == 0xBB && content[2] == 0xBF {
		content = content[3:]
	}

	result := gjson.ParseBytes(content)
	objItems := result.Get("ObjItems").Array()
	for _, objItem := range objItems {
		latLngData := objItem.Get("Object.ObjectDetail.Latlng").Raw
		_ = json.Unmarshal([]byte(latLngData), &latLngArr)

		for i := 0; i < len(latLngArr); i += 2 {
			points = append(points, model.Point{
				Latitude:  latLngArr[i],
				Longitude: latLngArr[i+1],
			})
		}
	}
	return points, nil
}
