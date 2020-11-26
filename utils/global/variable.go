package global

import (
	"github.com/crazy-me/os_snmp/conf"
	"go.uber.org/zap"
)

var (
	APP    conf.App
	LOGGER *zap.SugaredLogger
)
