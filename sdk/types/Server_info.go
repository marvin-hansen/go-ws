package types

type ServerInfo struct {
	Time            *string `json:"time,omitempty"`
	ExchangeId      *string `json:"exchange_id,omitempty"`
	InstanceGuid    *string `json:"instance_guid,omitempty"`
	ServerVersion   *string `json:"server_version,omitempty"`
	ServerName      *string `json:"server_name,omitempty"`
	DnsName         *string `json:"dns_name,omitempty"`
	IsRunning       *bool   `json:"is_running,omitempty"`
	ServerStartTime *string `json:"time_server_start,omitempty"`
}

func NewServerInfo() (this *ServerInfo) {
	this = new(ServerInfo)
	return this
}
