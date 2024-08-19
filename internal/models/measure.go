package models

type Measure struct {
	Value float64 `json:"value"`
	UOM   string  `json:"uom"`
}
