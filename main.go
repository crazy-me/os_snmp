package main

import (
	"fmt"
	"github.com/crazy-me/os_snmp/api"
	"github.com/crazy-me/os_snmp/consul"
	"github.com/crazy-me/os_snmp/initialize"
	"github.com/crazy-me/os_snmp/utils"
	"github.com/crazy-me/os_snmp/utils/global"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
	"net"
	"os"
	"strconv"
)

// grpc 编译
//go:generate protoc --go_out=plugins=grpc:. ./proto/snmpWalk.proto

func main() {
	args := utils.ParseArgs()
	if args.Community != "conf" {
		// 命令模式
		api.SnmpDebug(args)
		os.Exit(1)
	}

	// 服务模式
	server := grpc.NewServer()
	// GRPC服务注册
	initialize.ServiceRegister(server)

	// CONSUL服务注册
	grpc_health_v1.RegisterHealthServer(server, &consul.HealthImpl{})
	register := &consul.Register{
		Id:      "os_snmp_server",
		Address: utils.GetOutboundIP(),
		Port:    global.APP.System.Port,
		Tag:     []string{"snmp", "rpc"},
		Name:    "os_snmp",
	}
	_, err := register.ServerRegister(global.APP.Consul.Address)
	if err != nil {
		fmt.Println("consul register error !")
		os.Exit(1)
	}

	listener, err := net.Listen("tcp", ":"+strconv.Itoa(global.APP.System.Port))
	if err != nil {
		global.LOGGER.Errorf("tcp server err", zap.Any("err", err))
		panic(err)
	}

	fmt.Printf("server run successful! http://%s:%s\n", utils.GetOutboundIP(), strconv.Itoa(global.APP.System.Port))
	_ = server.Serve(listener)

}
