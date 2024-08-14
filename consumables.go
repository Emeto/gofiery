package gofiery

import "net/http"

type Consumables struct {
	Colorants []Colorant `json:"colorants"`
	Trays     []Tray     `json:"trays"`
}

type Colorant struct {
	Name  string `json:"name"`
	I18n  string `json:"i18n"`
	Color string `json:"color"`
	Level int    `json:"level"`
}

type Tray struct {
	Name                string         `json:"name"`
	TrayID              int            `json:"trayid"`
	I18n                string         `json:"i18n"`
	Dimensions          []float64      `json:"dimensions"`
	Lef                 bool           `json:"lef"`
	Level               int            `json:"level"`
	Attributes          trayAttributes `json:"attributes"`
	PaperCatalogMediaID int            `json:"pcmid"`
}

type trayAttributes struct {
	MediaType   string `json:"EFMediaType"`
	MediaWeight string `json:"EFMediaWeight"`
	PrintSize   string `json:"EFPrintSize"`
	PageSize    string `json:"PageSize"`
}

// GetConsumables reports information about the supply of paper,
// tray and toner on the print engine
func GetConsumables(fc *FieryClient) *Consumables {
	var consumables Consumables
	response := fc.Run(fc.Endpoint("consumables"), http.MethodGet)
	consumables = response.data.item.(Consumables)
	return &consumables
}
