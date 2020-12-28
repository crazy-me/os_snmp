package consul

import (
	"fmt"
	"github.com/crazy-me/os_snmp/utils/global"
	"github.com/hashicorp/consul/api"
	"go.uber.org/zap"
)

type ServiceDiscovery struct {
	Address     string
	ServiceName string
}

func (d *ServiceDiscovery) GetServiceInfo() {
	config := api.DefaultConfig()
	config.Address = d.Address
	clientDiscovery, err := api.NewClient(config)
	if err != nil {
		global.LOGGER.Errorf("sNewClient:", zap.Any("err", err))

	}

	services, _, err := clientDiscovery.Health().Service(d.ServiceName, d.ServiceName, true, &api.QueryOptions{WaitIndex: 0})
	if err != nil {
		global.LOGGER.Errorf("instances from Consul:", zap.Any("err", err))
	}
	fmt.Println(services)

}
