package models


type Server struct {
	Busy			bool	`json:"Busy,omitempty"`
	DatacenterId		int	`json:"DatacenterId,omitempty"`
	Name			string	`json:"Name,omitempty"`
	ServerId		int	`json:"ServerId,omitempty"`
	ServerStatus		int	`json:"ServerStatus,omitempty"`
	UserId			int	`json:"UserId,omitempty"`
	EasyCloudIPAddress	struct { Value string `json:"Value,omitempty"`} // We only get the IP Address
}