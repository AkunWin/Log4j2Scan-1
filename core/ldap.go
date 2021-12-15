package core

import (
	"fmt"
	"github.com/KpLi0rn/Log4j2Scan/config"
	"github.com/KpLi0rn/Log4j2Scan/log"
	module "github.com/KpLi0rn/Log4j2Scan/model"
	"net"
)

var (
	ResultChan chan *module.Result
)

func StartFakeServer(resultChan *chan *module.Result) {
	ResultChan = *resultChan
	log.Info("start fake reverse server")
	listen, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", config.Port))
	if err != nil {
		log.Error("listen fail err: %s", err)
		return
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Error("listen accept fail err: %s", err)
			continue
		}
		go acceptProcess(&conn)
	}
}

func acceptProcess(conn *net.Conn) {
	buf := make([]byte, 1024)
	num, err := (*conn).Read(buf)
	if err != nil {
		log.Error("accept data reading err: %s", err)
		return
	}
	hexStr := fmt.Sprintf("%x", buf[:num])
	// LDAP Protocol
	if "300c020101600702010304008000" == hexStr {
		res := &module.Result{
			Host:   (*conn).RemoteAddr().String(),
			Name:   "LDAP",
			Finger: hexStr,
		}
		ResultChan <- res
		_ = (*conn).Close()
		return
	}
	// RMI Protocol
	if checkRMI(buf) {
		res := &module.Result{
			Host:   (*conn).RemoteAddr().String(),
			Name:   "RMI",
			Finger: fmt.Sprintf("%x", buf[0:7]),
		}
		ResultChan <- res
		_ = (*conn).Close()
		return
	}
}

// RMI Protocol Docs:
// https://docs.oracle.com/javase/9/docs/specs/rmi/protocol.html
func checkRMI(data []byte) bool {
	if data[0] == 0x4a &&
		data[1] == 0x52 &&
		data[2] == 0x4d &&
		data[3] == 0x49 {
		if data[4] != 0x00 &&
			data[4] != 0x01 {
			return false
		}
		if data[6] != 0x4b &&
			data[6] != 0x4c &&
			data[6] != 0x4d {
			return false
		}
		lastData := data[7:]
		for _, v := range lastData {
			if v != 0x00 {
				return false
			}
		}
		return true
	}
	return false
}
