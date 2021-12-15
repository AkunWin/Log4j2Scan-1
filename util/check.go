package util

import (
	"github.com/KpLi0rn/Log4j2Scan/log"
)

func CheckPort(port int) bool {
	if port > 0 && port < 1024 {
		log.Warn("use system port: %d", port)
	}
	if port > 65535 || port < 1 {
		log.Error("use error port: %d", port)
		return false
	}
	return true
}
