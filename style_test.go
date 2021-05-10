package glog

import (
	"fmt"
	"testing"
	"time"
	"unsafe"

	"git.sr.ht/~poldi1405/go-ansi"
)

func TestDefaultLogFormat(t *testing.T) {
	timestamp := time.Now()

	message := "test message"
	formattedMsg := fmt.Sprintf(logFormat, styleFuncs[DEBUG](DEBUG.Short()), timestamp.Format(TimeFormat), GetCaller(0)+" – ", message)
	formattedMsg_test := defaultLogFormat(DEBUG, timestamp, GetCaller(0), message)
	if formattedMsg != formattedMsg_test {
		t.Fail()
	}

	message = "test 123 !#*"
	formattedMsg = fmt.Sprintf(logFormat, styleFuncs[DEBUG](DEBUG.Short()), timestamp.Format(TimeFormat), GetCaller(0)+" – ", message)
	formattedMsg_test = defaultLogFormat(DEBUG, timestamp, GetCaller(0), message)
	if formattedMsg != formattedMsg_test {
		t.Fail()
	}
}

func TestSetStyle(t *testing.T) {
	tmp := styleFuncs

	tmpFunc := ansi.Blue
	SetStyle(ERROR, tmpFunc)
	lvlFunc := styleFuncs[ERROR]
	if *(*uintptr)(unsafe.Pointer(&lvlFunc)) != *(*uintptr)(unsafe.Pointer(&tmpFunc)) {
		t.Fail()
	}

	tmpFunc = FatalStyle
	SetStyle(WARNING, tmpFunc)
	lvlFunc = styleFuncs[WARNING]
	if *(*uintptr)(unsafe.Pointer(&lvlFunc)) != *(*uintptr)(unsafe.Pointer(&tmpFunc)) {
		t.Fail()
	}

	// reset
	styleFuncs = tmp
}

func TestSetStyleInvalid(t *testing.T) {
	tmp := styleFuncs

	SetStyle(42, ansi.Blue)
	for k, v := range styleFuncs {
		if *(*uintptr)(unsafe.Pointer(&tmp[k])) != *(*uintptr)(unsafe.Pointer(&v)) {
			t.Fail()
		}
	}

	// reset
	styleFuncs = tmp
}

func TestFatalStyle(t *testing.T) {
	str := "test string"
	fatalStr := ansi.Blink(ansi.Bold(ansi.Red(str)))
	if fatalStr != FatalStyle(str) {
		t.Fail()
	}

	str = "# 123 !_Ä"
	fatalStr = ansi.Blink(ansi.Bold(ansi.Red(str)))
	if fatalStr != FatalStyle(str) {
		t.Fail()
	}
}

func TestPlainStyle(t *testing.T) {
	str := "test string"
	plainStr := fmt.Sprint(str)
	if plainStr != PlainStyle(str) {
		t.Fail()
	}

	str = "# 123 !_Ä"
	plainStr = fmt.Sprint(str)
	if plainStr != PlainStyle(str) {
		t.Fail()
	}
}
