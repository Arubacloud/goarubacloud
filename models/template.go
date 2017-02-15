package models

type GetHypvervisorTypeResponse struct {
	HypervisorServerType int        `json:"HypervisorServerType,omitempty"`
	HypervisorType       int        `json:"HypervisorType,omitempty"`
	Templates []Template
}

type Template struct {
	Id                          int        `json:"Id,omitempty"`
	Name 						string `json:"Name,omitempty"`
	TemplateSellingStatus       int        `json:"TemplateSellingStatus,omitempty"`
}

type GetPreconfiguredPackagesResponse struct {
	CloudPackages []CloudPackage
}

type CloudPackage struct {
	PackageID                          int        `json:"PackageID,omitempty"`
	Descriptions []LocalizedText
}

type LocalizedText struct {
	LanguageID                          int        `json:"LanguageID,omitempty"`
	Text 	string `json:"Text,omitempty"`
}
