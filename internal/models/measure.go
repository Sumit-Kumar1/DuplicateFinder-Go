package models

type Measure struct {
	Value float64          `json:"value"`
	UOM   []TranslatedText `json:"uom"`
}
