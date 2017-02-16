package models

type CloudPackage struct {
	PackageID                          int        `json:"PackageID,omitempty"`
	Descriptions []LocalizedText
}

