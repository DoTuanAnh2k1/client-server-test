package service

import "net/http"

var clientList []*http.Client

var (
	serverSvc  string
	serverPort string
)

var indexClientGet int

var isSvc bool
