package yadsh

import (
	"io"
	"net/http"
	"strings"
	"testing"
)

func TestGetHrefErrResp(t *testing.T) {
	res := &http.Response{
		Body: io.NopCloser(strings.NewReader(`{"code":404,"message":"not found"}`)),
	}

	href, err := getHref(res)
	if err == nil {
		t.Fatalf("err exptected, but got: %v", err)
	}
	if href != "" {
		t.Fatalf("empty string exptected, but got: %q", href)
	}
}
