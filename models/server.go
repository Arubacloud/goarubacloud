package models


type Server struct {
	Busy			bool	`json:"Busy,omitempty"`
	//CPUQuantity		int	`json:"CPUQuantity,omitempty"`
	//CompanyId		int	`json:"CompanyId,omitempty"`
	DatacenterId		int	`json:"DatacenterId,omitempty"`
	//HDQuantity		int	`json:"HDQuantity,omitempty"`
	//HDTotalSize		int	`json:"HDTotalSize,omitempty"`
	//HypervisorServerType	int	`json:"HypervisorServerType,omitempty"`
	//HypervisorType		int	`json:"HypervisorType,omitempty"`
	Name			string	`json:"Name,omitempty"`
	//OSTemplateId		int	`json:"OSTemplateId,omitempty"`
	//RAMQuantity		int	`json:"RAMQuantity,omitempty"`
	ServerId		int	`json:"ServerId,omitempty"`
	ServerStatus		int	`json:"ServerStatus,omitempty"`
	UserId			int	`json:"UserId,omitempty"`
	EasyCloudIPAddress	struct { Value string `json:"Value,omitempty"`} // We only get the IP Address
}