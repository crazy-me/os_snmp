package main

import (
	_ "github.com/crazy-me/os_snmp/initialize"
	"github.com/crazy-me/os_snmp/msg"
	"github.com/crazy-me/os_snmp/service"
)

//go:generate protoc --go_out=plugins=grpc:. ./proto/snmpWalk.proto

func main() {

	requestInfo := msg.SnmpV2Request{
		Timeout:          3,
		Retries:          1,
		SecurityLevel:    "",
		Network:          "udp",
		Address:          "192.168.31.138:161",
		Community:        "public",
		UserName:         "",
		AuthPassword:     "",
		AuthProtocol:     "",
		PrivPassword:     "",
		PrivProtocol:     "",
		SecurityEngineId: "",
		ContextEngineId:  "",
		ContextName:      "",
		Version:          "2c",
		Oid:              "1.3.6.1.2.1.1.1 1.3.6.1.2.1.1.5 1.3.6.1.2.1.1.3",
	}

	service.SnmpWalk(&requestInfo)

}
