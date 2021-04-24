package glog

import (
	"bytes"
	"log"
	"os"
	"strings"
	"testing"
)

func TestMetalog(t *testing.T) {
	tmp := EnableMetaLogging

	var b bytes.Buffer
	log.SetOutput(&b)
	t.Cleanup(func() {
		log.SetOutput(os.Stdout)
		EnableMetaLogging = tmp
	})

	EnableMetaLogging = true
	metalog("test 123")
	if !strings.HasSuffix(b.String(), "test 123\n") {
		t.Fail()
	}

	b.Reset()
	EnableMetaLogging = false
	metalog("test 123")
	if len(b.String()) > 0 {
		t.Fail()
	}
}
