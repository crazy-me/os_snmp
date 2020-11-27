package api

import (
	"context"
	"github.com/crazy-me/os_snmp/msg"
	"github.com/crazy-me/os_snmp/service"
)

type SnmpWalkImpl struct {
}

func (s *SnmpWalkImpl) GetSnmpWalk(ctx context.Context, request *msg.SnmpV2Request) (*msg.SnmpV2Response, error) {

	response := &msg.SnmpV2Response{}
	sliceList, err := service.SnmpWalk(request)
	if err != nil {
		return response, err
	}

	for _, val := range sliceList {
		response.Result = append(response.Result, &msg.Result{
			Oid:   val["oid"],
			Value: val["value"],
		})
	}

	return response, nil
}
