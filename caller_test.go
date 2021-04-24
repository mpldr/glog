package glog

import (
	"testing"
)

func caller(skipFrames int) string {
	return GetCaller(skipFrames)
}

func TestGetCaller(t *testing.T) {
	if caller(0) != "git.sr.ht/~poldi1405/glog.caller" {
		t.Fail()
	}
	if caller(1) != "git.sr.ht/~poldi1405/glog.TestGetCaller" {
		t.Fail()
	}
}
