package models

/*
"Value":{
	"ActiveJobs":[],
	"CPUQuantity":{
		"CompanyId":1,
		"ProductId":4,
		"ResourceId":330951,
		"ResourceType":2,
		"UserId":1856,
		"Quantity":2
	},
	"CompanyId":1,
	"ControlToolActivationDate":null,
	"ControlToolInstalled":false,
	"CreationDate":"\/Date(1446133084103+0100)\/",
	"DatacenterId":1,
	"EasyCloudIPAddress":null,
	"EasyCloudPackageID":null,
	"HypervisorServerType":2,
	"HypervisorType":2,
	"Name":"testkam",
	"NetworkAdapters":[{
			"IPAddresses":[{
				"CompanyId":1,
				"ProductId":20,
				"ResourceId":330950,
				"ResourceType":6,
				"UserId":1856,
				"Gateway":"95.110.165.1",
				"LoadBalancerID":null,
				"ServerId":56056,
				"SubNetMask":"255.255.255.0",
				"Value":"95.110.165.246"
			}],
			"Id":148263,
			"MacAddress":"00:50:56:93:06:b9",
			"NetworkAdapterType":0,
			"ServerId":56056,
			"VLan":null
			},{
			"IPAddresses":[],
			"Id":148264,
			"MacAddress":"00:50:56:93:22:d7",
			"NetworkAdapterType":1,
			"ServerId":56056,
			"VLan":null
			},{
			"IPAddresses":[],
			"Id":148265,
			"MacAddress":"00:50:56:93:7f:38",
			"NetworkAdapterType":2,
			"ServerId":56056,
			"VLan":null
	}],
	"Note":"<script>alert('attacked')<\/script>",
	"OSTemplate":{
		"CompanyId":1,
		"ProductId":393,
		"ResourceId":330954,
		"ResourceType":4,
		"UserId":1856,
		"Description":"Ubuntu Server 14.04 LTS 64bit",
		"Id":1187,
		"Name":"ubuntu1404_x64_1_0"},
		"Parameters":[
			{"Key":0,"Value":"10.11.10.18"},
			{"Key":1,"Value":"dc01vcenter502.intra.cloud.it"},
			{"Key":2,"Value":"50131eb0-f6ce-af98-f542-98dc1c69681b"},
			{"Key":3,"Value":"vm-21365"}
		],
		"RAMQuantity":{
			"CompanyId":1,
			"ProductId":5,
			"ResourceId":330952,
			"ResourceType":1,
			"UserId":1856,
			"Quantity":4
		},
		"RenewDateSmart":null,
		"ScheduledOperations":[],
		"ServerId":56056,
		"ServerStatus":3,
		"Snapshots":[],
		"ToolsAvailable":true,
		"UserId":1856,
		"VirtualDVDs":[],
		"VirtualDisks":[{
			"CompanyId":1,
			"ProductId":6,
			"ResourceId":330953,
			"ResourceType":3,
			"UserId":1856,
			"CreationDate":"\/Date(1446133084103+0100)\/",
			"Size":220
		}],
		"VncPort":23685
		}
	}
*/
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