package models

type TemplateDate struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{}
	CRFToken  string
	Flash     string
	Warning   string
	Error     string
}
