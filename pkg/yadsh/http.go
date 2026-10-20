package yadsh

import (
	"net/http"
	"time"
)

const (
	clientTimeOut = 30 * time.Second
)

var client = http.Client{Timeout: clientTimeOut}
