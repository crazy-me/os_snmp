package api

import (
	"encoding/json"
	"fmt"
	"github.com/crazy-me/os_snmp/msg"
	"github.com/crazy-me/os_snmp/service"
)

type ResponseMessage struct {
	Code int                 `json:"code"` // 状态码
	Msg  string              `json:"msg"`  // 状态消息
	Data []map[string]string `json:"data"`
}

func SnmpDebug(request *msg.SnmpV2Request) {
	message := &ResponseMessage{}
	list, err := service.SnmpWalk(request)
	if err != nil {
		message.Code = 500
		message.Msg = err.Error()
	} else {
		sliceList := make([]map[string]string, 0)
		for _, val := range list {
			sliceList = append(sliceList, map[string]string{
				"oid":   val["oid"],
				"value": val["value"],
			})
		}
		message.Code = 200
		message.Msg = "successful"
		message.Data = sliceList
	}

	bytes, err := json.Marshal(message)
	fmt.Println(string(bytes))

}
