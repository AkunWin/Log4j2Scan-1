package config

import "time"

var (
	Port     int
	HttpPort int
)

const (
	DefaultChannelSize = 100
	DefaultHttpTimeout = time.Second * 3
	DefaultHttpPath    = "/"
)
