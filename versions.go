package gofiery

import (
	"net/http"
	"time"
)

type Version struct {
	Time time.Time   `json:"time"`
	Data VersionData `json:"data"`
}

type VersionData struct {
	Kind string
	Item VersionItem
}

type VersionItem struct {
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

type SpecVersion interface {
	ApacheVersion
	FieryAPIVersion
	OpenSSLVersion
	RhythmVersion
	RubyVersion
}

func GetVersions(fc *FieryClient) *Version {
	var versions Version
	fc.Run(fc.Endpoint("versions"), http.MethodGet, &versions)
	return &versions
}

func GetVersion[V SpecVersion](of string, fc *FieryClient) *V {
	var version V
	fc.Run(fc.Endpoint("versions?package="+of), http.MethodGet, &version)
	return &version
}
