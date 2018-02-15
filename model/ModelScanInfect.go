package model

type ScanInfectModel struct {
	Name          string        `json:"name"`
	Type          string        `json:"type"`
	PercentInfect string        `json:"percent_infect"`
	AreaInfect    []LatLngModel `json:"area_infect"`
}
