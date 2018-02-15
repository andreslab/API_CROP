package model

import (
	utils "github.com/andreslab/prj_api_crop/model/utils"
)

type ContentScanModel struct {
	Data []ScanModel `json:"data"`
}

type ScanModel struct {
	ID            int64
	Date          utils.DateModel   `json:"created"`
	PercentInfect string            `json:"percent_infect"`
	GalleryID     string            `json:"gallery_id"`
	Zone          string            `json:"zone"`
	Area          []LatLngModel     `json:"area"`
	AreaInfection []ScanInfectModel `json:"area_infect"`
}
