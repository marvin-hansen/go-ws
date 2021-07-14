package types

import "fmt"

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

func (s ServerInfo) String() string {
	return fmt.Sprintf("<ServerInfo> Time: %v, ExchangeId: %v, InstanceGuid: %v, ServerVersion: %v, ServerName: %v, DnsName:  %v, IsRunning  %v, ServerStartTime: %v",
		*s.Time,
		*s.ExchangeId,
		*s.InstanceGuid,
		*s.ServerVersion,
		*s.ServerName,
		*s.DnsName,
		*s.IsRunning,
		*s.ServerStartTime,
	)
}
