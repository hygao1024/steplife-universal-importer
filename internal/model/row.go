package model

type Row struct {
	DataTime         int
	LocType          int
	Heading          int
	Accuracy         int
	Speed            int
	Distance         int
	IsBackForeground int
	StepType         int
	Altitude         int
	Point
}

func NewRow() *Row {
	return &Row{
		DataTime:         0,
		LocType:          1,
		Heading:          0,
		Accuracy:         14,
		Speed:            0,
		Distance:         0,
		IsBackForeground: 0,
		StepType:         0,
		Altitude:         316,
		Point: Point{
			Latitude:  0.0,
			Longitude: 0.0,
		},
	}
}
