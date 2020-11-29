package consul

import (
	"fmt"
	"github.com/crazy-me/os_snmp/utils/global"
	"github.com/hashicorp/consul/api"
	"time"
)

type Register struct {
	Id      string
	Address string
	Port    int
	Tag     []string
	Name    string
}

func (r *Register) ServerRegister(centerAddress string) (bool, error) {
	consulConfig := api.DefaultConfig()
	consulConfig.Address = centerAddress
	client, err := api.NewClient(consulConfig)
	if err != nil {
		return false, err
	}
	agent := client.Agent()
	interval := time.Duration(global.APP.Consul.Interval) * time.Second
	deregister := time.Duration(global.APP.Consul.ExpireTime) * time.Minute

	reg := &api.AgentServiceRegistration{
		ID:      r.Id,      // 服务唯一ID
		Name:    r.Name,    // 服务逻辑名称
		Tags:    r.Tag,     // 标签用于过滤
		Port:    r.Port,    // 服务端口
		Address: r.Address, // 服务地址
		Check: &api.AgentServiceCheck{ // 健康检查
			Interval: interval.String(), // 健康检查间隔
			//HTTP: "http://"+r.Address+":"+strconv.Itoa(r.Port)+r.Name, // http 健康检测
			GRPC: fmt.Sprintf("%v:%v/%v", r.Address, r.Port, r.Name), // grpc 支持，执行健康检查的地址，service 会传到 Health.Check 函数中
			//Method: "POST", //HTTP检查的方式 默认GET
			DeregisterCriticalServiceAfter: deregister.String(), // 注销时间，相当于过期时间
		},
	}

	err1 := agent.ServiceRegister(reg)
	if err1 != nil {
		return false, err
	}
	return true, nil
}
