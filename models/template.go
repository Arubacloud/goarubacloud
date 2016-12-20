package models

type GetHypvervisorTypeResponse struct {
	HypervisorServerType int        `json:"HypervisorServerType,omitempty"`
	HypervisorType       int        `json:"HypervisorType,omitempty"`
	Templates []Template
}

type Template struct {
	ApplianceType               string        `json:"ApplianceType,omitempty"`
	ArchitectureType            int        `json:"ArchitectureType,omitempty"`
	CompanyId                   int        `json:"CompanyId,omitempty"`
	Description                 string        `json:"Description,omitempty"`
	Enabled                     bool        `json:"Enabled,omitempty"`
	ExportEnabled               bool        `json:"ExportEnabled,omitempty"`
	Id                          int        `json:"Id,omitempty"`
	IdentificationCode          string        `json:"IdentificationCode,omitempty"`
	Name                        string        `json:"Name,omitempty"`
	OSFamily                    int        `json:"OSFamily,omitempty"`
	OSVersion                   string        `json:"OSVersion,omitempty"`
	OwnerUserId                 int        `json:"OwnerUserId,omitempty"`
	ParentTemplateID            int        `json:"ParentTemplateID,omitempty"`
	ProductId                   int        `json:"ProductId,omitempty"`
	ResourceBounds              []struct {
		Default      int        `json:"Default,omitempty"`
		Max          int        `json:"Max,omitempty"`
		ResourceType int        `json:"ResourceType,omitempty"`
		Min          int        `json:"Min,omitempty"`
	}
	Revision                    string        `json:"Revision,omitempty"`
	TemplateExtendedDescription string        `json:"TemplateExtendedDescription,omitempty"`
	TemplateOwnershipType       int        `json:"TemplateOwnershipType,omitempty"`
	TemplatePassword            string        `json:"TemplatePassword,omitempty"`
	TemplateSellingStatus       int        `json:"TemplateSellingStatus,omitempty"`
	TemplateStatus              int        `json:"TemplateStatus,omitempty"`
	TemplateType                int        `json:"TemplateType,omitempty"`
	TemplateUsername            string        `json:"TemplateUsername,omitempty"`
	ToolsAvailable              bool        `json:"ToolsAvailable,omitempty"`
}
