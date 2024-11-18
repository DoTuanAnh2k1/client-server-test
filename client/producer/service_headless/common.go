package serviceheadless

import "client/common"

var connectionList []*common.Connection

var (
	serverHeadlessSvc string
	serverPort        string
)

var indexConnectionGet int
var indexClientGet int

var isHeadlessSvc bool
