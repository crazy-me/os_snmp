package utils

import (
	"context"
	"fmt"
	"github.com/crazy-me/os_snmp/msg"
	"google.golang.org/grpc"
	"os"
	"testing"
)

func TestSnmp(t *testing.T) {
	conn, err := grpc.Dial("localhost:8888", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := msg.NewSnmpV2ServiceClient(conn)

	requestInfo := &msg.SnmpV2Request{
		Timeout:          2,
		Retries:          1,
		SecurityLevel:    "",
		Network:          "tcp",
		Address:          "172.16.4.49:161",
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
		Oid:              "1.3.6.1.2.1.1.1",
	}
	response, err := client.GetSnmpWalk(context.Background(), requestInfo)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(response.Result)

}
