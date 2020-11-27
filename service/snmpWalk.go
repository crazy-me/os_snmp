package service

import (
	"github.com/crazy-me/os_snmp/msg"
	"github.com/crazy-me/os_snmp/utils/global"
	"github.com/k-sone/snmpgo"
	"go.uber.org/zap"
	"strings"
	"time"
)

// Todo: {"Oid": "1.3.6.1.2.1.1.1.0", "Variable": {"Type": "OctetString", "Value": "Linux walle.haoyue.com 3.10.0-1127.8.2.el7.x86_64 #1 SMP Tue May 12 16:57:42 UTC 2020 x86_64"}}
// Todo: {"Oid": "1.3.6.1.2.1.1.3.0", "Variable": {"Type": "TimeTicks", "Value": "137910603"}}
// Todo: {"Oid": "1.3.6.1.2.1.1.5.0", "Variable": {"Type": "OctetString", "Value": "walle.haoyue.com"}}

func SnmpWalk(request *msg.SnmpV2Request) ([]map[string]string, error) {
	varBindsList := make([]map[string]string, 0)
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
		return varBindsList, &snmpgo.ArgumentError{
			Value:   "params:" + request.Version,
			Message: "Version not supported",
		}
	}
	if request.Version == "3" {
		switch request.SecurityLevel {
		case "NoAuthNoPriv":
			args.SecurityLevel = snmpgo.NoAuthNoPriv
		case "AuthNoPriv":
			args.SecurityLevel = snmpgo.AuthNoPriv
		case "AuthPriv":
			args.SecurityLevel = snmpgo.AuthPriv
		default:
			return varBindsList, &snmpgo.ArgumentError{
				Value:   "params:" + request.SecurityLevel,
				Message: "SecurityLevel not supported",
			}
		}
	}

	oidSlice, err := snmpgo.NewOids(strings.Fields(request.Oid))
	if err != nil {
		global.LOGGER.Errorf("snmpgo.NewOids", zap.Any("err", err))
		return varBindsList, err
	}

	snmp, err := snmpgo.NewSNMP(*args)
	if err != nil {
		global.LOGGER.Errorf("snmpgo.NewSNMP", zap.Any("err", err))
		return varBindsList, err
	}

	if err = snmp.Open(); err != nil {
		global.LOGGER.Errorf("snmp.Open", zap.Any("err", err))
		return varBindsList, err
	}

	defer snmp.Close()

	result, err := snmp.GetBulkWalk(oidSlice, 0, 10)

	if result.ErrorStatus() != snmpgo.NoError {
		global.LOGGER.Errorf("result.ErrorStatus", zap.Any("err", err))
		return varBindsList, err
	}

	for _, val := range result.VarBinds() {
		varBindsMap := make(map[string]string)
		varBindsMap["oid"] = val.Oid.String()
		varBindsMap["value"] = val.Variable.String()
		varBindsList = append(varBindsList, varBindsMap)
	}

	return varBindsList, nil

}
