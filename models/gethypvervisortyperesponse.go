package models

type GetHypvervisorTypeResponse struct {
	HypervisorServerType int        `json:"HypervisorServerType,omitempty"`
	HypervisorType       int        `json:"HypervisorType,omitempty"`
	Templates []Template
}

