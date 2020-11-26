package api

import (
	"context"
	"github.com/crazy-me/os_snmp/msg"
)

type SnmpV2Impl struct {
}

func (s *SnmpV2Impl) GetSnmpWalk(ctx context.Context, request *msg.SnmpV2Request) (*msg.SnmpV2Response, error) {

}
