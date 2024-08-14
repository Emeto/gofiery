package gofiery

import (
	"net/http"
	"time"
)

type Info struct {
	DiskAvailable   int64     `json:"disk_available"`
	DiskTotal       int64     `json:"disk_total"`
	EpochTime       time.Time `json:"epoch_time"`
	FieryLocale     string    `json:"fiery_locale"`
	Locale          string    `json:"locale"`
	MemoryAvailable int64     `json:"memory_available"`
	MemoryTotal     int64     `json:"memory_total"`
	Name            string    `json:"name"`
	OSLocale        string    `json:"os_locale"`
	SerialNumber    string    `json:"serial_number"`
	Timezone        string    `json:"timezone"`
	Uptime          int64     `json:"uptime"`
	Version         string    `json:"version"`
	Host            string    `json:"host"`
	Username        string    `json:"username"`
	AppId           string    `json:"app_id"`
}

func GetInfo(fc *FieryClient) *Info {
	var info Info
	response := fc.Run(fc.Endpoint("info"), http.MethodGet)
	info = response.data.item.(Info)
	return &info
}
