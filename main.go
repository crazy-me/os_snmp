package main

import (
	"fmt"
	"github.com/crazy-me/os_snmp/api"
	"github.com/crazy-me/os_snmp/initialize"
	"github.com/crazy-me/os_snmp/utils"
	"github.com/crazy-me/os_snmp/utils/global"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
	"os"
	"strconv"
)

//go:generate protoc --go_out=plugins=grpc:. ./proto/snmpWalk.proto

func main() {
	args := utils.ParseArgs()
	if args.Community != "config" {
		// 命令模式
		api.SnmpDebug(args)
		os.Exit(1)
	}

	// 服务模式
	server := grpc.NewServer()
	// GRPC服务注册
	initialize.ServiceRegister(server)

	// CONSUL服务注册
	// ...

	listener, err := net.Listen("tcp", ":"+strconv.Itoa(global.APP.System.Port))
	if err != nil {
		global.LOGGER.Errorf("tcp server err", zap.Any("err", err))
		panic(err)
	}
	fmt.Printf("run server success! %s\n", "http://127.0.0.1:"+strconv.Itoa(global.APP.System.Port))
	_ = server.Serve(listener)

}
