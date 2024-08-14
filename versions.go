package gofiery

import (
	"net/http"
	"time"
)

type Version struct {
	Apache   ApacheVersion
	FieryAPI FieryAPIVersion `json:"fieryApi"`
	Gems     []interface{}
	OpenSSL  OpenSSLVersion `json:"openSsl"`
	Rhythm   RhythmVersion
	Ruby     RubyVersion
}

type ApacheVersion struct {
	Version   string
	Platform  string
	BuiltDate time.Time `json:"builtDate"`
}

type FieryAPIVersion struct {
	Version       string
	InstalledDate time.Time `json:"installedDate"`
}

type OpenSSLVersion struct {
	Version   string
	BuiltDate time.Time `json:"builtDate"`
}

type RhythmVersion struct {
	Version string
}

type RubyVersion struct {
	Version string
}

func GetVersions(fc *FieryClient) *Version {
	var versions Version
	response := fc.Run(fc.Endpoint("versions"), http.MethodGet)
	versions = response.data.item.(Version)
	return &versions
}

func GetVersion(of string, fc *FieryClient) any {
	var version any
	versions := GetVersions(fc)
	switch of {
	case "apache":
		version = versions.Apache
	case "fieryApi":
		version = versions.FieryAPI
	case "gems":
		version = versions.Gems
	case "openSsl":
		version = versions.OpenSSL
	case "rhythm":
		version = versions.Rhythm
	case "ruby":
		version = versions.Ruby
	}
	return &version
}
