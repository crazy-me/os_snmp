package utils

import (
	"flag"
	"fmt"
	"github.com/crazy-me/os_snmp/msg"
	"os"
)

func ParseArgs() *msg.SnmpV2Request {
	commandArgs := os.Args
	if len(commandArgs) == 1 {
		helpInfo()
		os.Exit(1)
	}

	timeout := flag.Int("t", 5, "Request timeout (number of seconds)")
	retries := flag.Int("r", 1, "Number of retries")
	securityLevel := flag.String("l", "NoAuthNoPriv", "Security level (NoAuthNoPriv|AuthNoPriv|AuthPriv)")
	network := flag.String("p", "udp", "default udp (udp|udp6|tcp|tcp6)")
	address := flag.String("h", "", "Request host to port (127.0.0.1:161)")
	community := flag.String("c", "public", "set think Community")
	userName := flag.String("u", "", "Security name")
	authPassword := flag.String("A", "", "Authentication protocol pass phrase")
	authProtocol := flag.String("a", "MD5", "Authentication protocol (MD5|SHA)")
	privPassword := flag.String("P", "", "Privacy protocol pass phrase")
	privProtocol := flag.String("x", "DES", "Privacy protocol (DES|AES)")
	securityEngineId := flag.String("e", "", "Security engine ID")
	contextEngineId := flag.String("E", "", "Context engine ID")
	contextName := flag.String("n", "", "Context name")
	version := flag.String("v", "2c", "SNMP version to use (2c|3)")
	oid := flag.String("o", "", "set this oid value")
	flag.Parse()

	requestInfo := &msg.SnmpV2Request{
		Timeout:          int32(*timeout),
		Retries:          int32(*retries),
		SecurityLevel:    *securityLevel,
		Network:          *network,
		Address:          *address,
		Community:        *community,
		UserName:         *userName,
		AuthPassword:     *authPassword,
		AuthProtocol:     *authProtocol,
		PrivPassword:     *privPassword,
		PrivProtocol:     *privProtocol,
		SecurityEngineId: *securityEngineId,
		ContextEngineId:  *contextEngineId,
		ContextName:      *contextName,
		Version:          *version,
		Oid:              *oid,
	}

	return requestInfo
}

func helpInfo() {
	fmt.Print(`
请输入正确的运行参数!
os_snmp支持以下两种运行方式:
	*服务模式: ./os_snmp -c conf
	*命令模式: 
		* v2c: ./os_snmp -v 2c -c public 192.168.1.1 1.3.6.1.2.1
		*  v3: ./os_snmp -v 3  -u admin -l authPriv -a md5 -A abc@123 -x aes -X abc@123 192.168.1.1 1.3.6.1.2.1
	参数选项:
	SNMP Version 2c specific
		-v -v 1|2c|3      specifies SNMP version to use
		-c COMMUNITY      set the community string
	SNMP Version 3 specific
		-a PROTOCOL       set authentication protocol (MD5|SHA)
		-A PASSPHRASE     set authentication protocol pass phrase
		-e ENGINE-ID      set security engine ID (e.g. 800000020109840301)
		-E ENGINE-ID      set context engine ID (e.g. 800000020109840301)
		-l LEVEL          set security level (noAuthNoPriv|authNoPriv|authPriv)
		-n CONTEXT        set context name (e.g. bridge1)
		-u USER-NAME      set security name (e.g. bert)
		-x PROTOCOL       set privacy protocol (DES|AES)
		-X  PASSPHRASE    set privacy protocol pass phrase
 `)
}
