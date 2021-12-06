package formats

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"

	"git.sr.ht/~poldi1405/glog"
)

func TestCLF(t *testing.T) {
	req := &http.Request{
		Method: "GET",
		URL: &url.URL{
			Path: "/some-file.txt",
		},
		Proto:      "HTTP",
		ProtoMajor: 1,
		ProtoMinor: 0,
		RemoteAddr: "127.0.0.1",
		Response: &http.Response{
			StatusCode: 200,
		},
	}
	glog.LogLevel = glog.INFO
	glog.LogFormatter = CommonLogFormat
	fmt.Println(CLFormatFromRequest(req, 1234))
	glog.Info(CLFormatFromRequest(req, 1234))
}
