package model

import "fmt"

type StepLife struct {
	CSVHeader [][]string
	CSVData   [][]string
}

func NewStepLife() *StepLife {
	return &StepLife{
		CSVHeader: [][]string{
			{
				"dataTime", "locType", "longitude", "latitude", "heading", "accuracy", "speed", "distance",
				"isBackForeground", "stepType", "altitude",
			},
		},
		CSVData: [][]string{},
	}
}

func (this *StepLife) AddCSVRow(row Row) {
	this.CSVData = append(this.CSVData, []string{
		fmt.Sprintf("%d", row.DataTime),
		fmt.Sprintf("%d", row.LocType),
		fmt.Sprintf("%.8f", row.Longitude),
		fmt.Sprintf("%.8f", row.Latitude),
		fmt.Sprintf("%d", row.Heading),
		fmt.Sprintf("%d", row.Accuracy),
		fmt.Sprintf("%.2f", row.Speed),
		fmt.Sprintf("%d", row.Distance),
		fmt.Sprintf("%d", row.IsBackForeground),
		fmt.Sprintf("%d", row.StepType),
		fmt.Sprintf("%.2f", row.Altitude),
	})
}
