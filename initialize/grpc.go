package initialize

import (
	"github.com/crazy-me/os_snmp/api"
	"github.com/crazy-me/os_snmp/msg"
	"google.golang.org/grpc"
)

func ServiceRegister(srv *grpc.Server) {
	// snmpWalk
	msg.RegisterSnmpV2ServiceServer(srv, &api.SnmpWalkImpl{})
}
