package model

type Row struct {
	LocType          int
	Heading          int
	Accuracy         int
	Distance         int
	IsBackForeground int
	StepType         int

	Point
}

func NewRow() *Row {
	return &Row{
		LocType:          1,
		Heading:          0,
		Accuracy:         14,
		Distance:         0,
		IsBackForeground: 0,
		StepType:         0,
		Point: Point{
			Altitude:  316,
			DataTime:  0,
			Speed:     0,
			Latitude:  0.0,
			Longitude: 0.0,
		},
	}
}
