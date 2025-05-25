package parser

import (
	"encoding/json"
	"github.com/tidwall/gjson"
	"steplife-universal-importer/internal/model"
	"steplife-universal-importer/internal/utils/logx"
)

type Ovjsn struct {
	BaseAdaptor
}

func NewOvjsnAdaptor() *Ovjsn {
	return &Ovjsn{}
}

func (this *Ovjsn) Parse(content []byte) ([]model.Point, error) {
	var points []model.Point

	// 检查是否有 BOM
	if len(content) >= 3 && content[0] == 0xEF && content[1] == 0xBB && content[2] == 0xBF {
		content = content[3:]
	}

	result := gjson.ParseBytes(content)
	objItems := result.Get("ObjItems").Array()
	for _, objItem := range objItems {
		objChildrenPoints := this.parseObjChildren(objItem)
		points = append(points, objChildrenPoints...)
	}
	return points, nil
}

func (this *Ovjsn) parseObjChildren(objItems gjson.Result) []model.Point {
	var points []model.Point
	for _, itme := range objItems.Array() {
		name := itme.Get("Object.Name").String()

		objectDetail := itme.Get("Object.ObjectDetail")
		objChildren := objectDetail.Get("ObjChildren")
		if objChildren.Exists() {
			logx.InfoF("开始解析ovjsn子文件夹（%s）", name)
			objChildrenPoints := this.parseObjChildren(objChildren)
			logx.InfoF("ovjsn子文件夹（%s）解析完成", name)
			points = append(points, objChildrenPoints...)
		} else {
			objDetailPoints := this.parseObjDetail(objectDetail)

			logx.InfoF("ovjsn子文件（%s）解析完成，坐标点数：%d", name, len(objDetailPoints))
			points = append(points, objDetailPoints...)
		}
	}
	return points
}

func (this *Ovjsn) parseObjDetail(objDetail gjson.Result) []model.Point {
	var points []model.Point
	var latLngArr []float64
	latLngData := objDetail.Get("Latlng").Raw
	_ = json.Unmarshal([]byte(latLngData), &latLngArr)

	for i := 0; i < len(latLngArr); i += 2 {
		points = append(points, model.Point{
			Latitude:  latLngArr[i],
			Longitude: latLngArr[i+1],
		})
	}

	return points
}
