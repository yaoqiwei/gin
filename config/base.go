package config

import (
	"gin/config/structs"
	"time"
)

var DebugMode string
var TimeZone *time.Location
var Http structs.HttpConf
