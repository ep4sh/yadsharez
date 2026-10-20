package yadsh

import (
	"net/http"
	"time"
)

const (
	timeout          = 30 * time.Second
	idleConn         = 100
	idleConnsPerHost = 10
	idleTimeout      = 90 * time.Second
)

var client = &http.Client{
	Timeout: timeout,
	Transport: &http.Transport{
		MaxIdleConns:        idleConn,
		MaxIdleConnsPerHost: idleConnsPerHost,
		IdleConnTimeout:     idleTimeout,
	},
}
