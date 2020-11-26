package service

import (
	"fmt"
	"github.com/crazy-me/os_snmp/msg"
	"github.com/k-sone/snmpgo"
	"strings"
	"time"
)

func SnmpWalk(request *msg.SnmpV2Request) {

	args := &snmpgo.SNMPArguments{
		Network:          request.Network,
		Address:          request.Address,
		Timeout:          time.Duration(request.Timeout) * time.Second,
		Retries:          uint(request.Retries),
		Community:        request.Community,
		UserName:         request.UserName,
		AuthPassword:     request.AuthPassword,
		AuthProtocol:     snmpgo.AuthProtocol(request.AuthProtocol),
		PrivPassword:     request.PrivPassword,
		PrivProtocol:     snmpgo.PrivProtocol(request.PrivProtocol),
		SecurityEngineId: request.SecurityEngineId,
		ContextEngineId:  request.ContextEngineId,
		ContextName:      request.ContextName,
	}

	// snmp 版本
	switch request.Version {
	case "2c":
		args.Version = snmpgo.V2c
	case "3":
		args.Version = snmpgo.V3
	default:

	}

	switch request.SecurityLevel {
	case "NoAuthNoPriv":
		args.SecurityLevel = snmpgo.NoAuthNoPriv
	case "AuthNoPriv":
		args.SecurityLevel = snmpgo.AuthNoPriv
	case "AuthPriv":
		args.SecurityLevel = snmpgo.AuthPriv
	default:

	}

	oidSlice, err := snmpgo.NewOids(strings.Fields(request.Oid))
	if err != nil {

	}

	snmp, err := snmpgo.NewSNMP(*args)
	if err != nil {

	}

	if err = snmp.Open(); err != nil {

	}

	defer snmp.Close()

	result, err := snmp.GetBulkWalk(oidSlice, 0, 10)

	if result.ErrorStatus() != snmpgo.NoError {

	}

	for _, val := range result.VarBinds() {
		fmt.Println(val)
	}

}
