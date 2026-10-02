package gofiery

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

type Licence struct {
	ID             string    `json:"id"`
	Name           string    `json:"name"`
	StartDate      time.Time `json:"startDate"`
	ExpirationDate time.Time `json:"expirationDate"`
	ActivationCode string    `json:"activationCode"`
	Status         string    `json:"status"`
}

func (fc *FieryClient) GetLicenses() []Licence {
	var licenses []Licence
	response := fc.Run(fc.Endpoint("licenses"), http.MethodGet)
	if items, ok := response.data.items.([]Licence); ok {
		licenses = items
	} else {
		_, err := fmt.Fprintf(os.Stderr, "gofiery: could not parse json response: %s\n", err)
		if err != nil {
			return nil
		}
	}
	return licenses
}

func (fc *FieryClient) ActivateLicense(ac string) {}

func (fc *FieryClient) GetLicense(id string) *Licence {
	var licence Licence
	response := fc.Run(fc.Endpoint("licences/"+id), http.MethodGet)
	licence = response.data.item.(Licence)
	return &licence
}

func (fc *FieryClient) DeactivateLicense(id string) {}
