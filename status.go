package gofiery

import "net/http"

type Status struct {
	Fiery               string `json:"fiery"`
	FieryExtendedStatus string `json:"fieryExtendedStatus"`
}

func GetStatus(fc *FieryClient) *Status {
	var status Status
	response := fc.Run(fc.Endpoint("status"), http.MethodGet)
	status = response.data.item.(Status)
	return &status
}
