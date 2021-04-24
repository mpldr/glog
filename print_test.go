package glog

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

func TestTrace(t *testing.T) {
	lvl := LogLevel
	var b bytes.Buffer
	SetLevel(TRACE)
	SetOutput(TRACE, &b)

	msg := "test 123 !#*"
	Trace(msg)
	if !strings.HasPrefix(b.String(), "TRACE") {
		t.Fail()
	}
	if !strings.HasSuffix(b.String(), GetCaller(0)+" – "+msg+"\n") {
		t.Fail()
	}

	b.Reset()
	SetShowCaller(TRACE, false)
	msg = "\thallo"
	Trace(msg)
	if !strings.HasPrefix(b.String(), "TRACE") {
		t.Fail()
	}
	if !strings.HasSuffix(b.String(), msg+"\n") {
		t.Fail()
	}
	if strings.HasSuffix(b.String(), GetCaller(0)+" – "+msg+"\n") {
		t.Fail()
	}

	// reset
	SetShowCaller(TRACE, true)
	SetOutput(TRACE, os.Stdout)
	SetLevel(lvl)
}

func TestTracef(t *testing.T) {
	lvl := LogLevel
	var b bytes.Buffer
	SetLevel(TRACE)
	SetOutput(TRACE, &b)

	msg := "test 123 !#*"
	Tracef("%s %d !#%c", "test", 123, '*')
	if !strings.HasPrefix(b.String(), "TRACE") {
		t.Fail()
	}
	if !strings.HasSuffix(b.String(), GetCaller(0)+" – "+msg+"\n") {
		t.Fail()
	}

	b.Reset()
	SetShowCaller(TRACE, false)
	msg = "\thallo"
	Tracef("\t%s", "hallo")
	if !strings.HasPrefix(b.String(), "TRACE") {
		t.Fail()
	}
	if !strings.HasSuffix(b.String(), msg+"\n") {
		t.Fail()
	}
	if strings.HasSuffix(b.String(), GetCaller(0)+" – "+msg+"\n") {
		t.Fail()
	}

	// reset
	SetShowCaller(TRACE, true)
	SetOutput(TRACE, os.Stdout)
	SetLevel(lvl)
}

func TestDebug(t *testing.T) {
	lvl := LogLevel
	var b bytes.Buffer
	SetLevel(DEBUG)
	SetOutput(DEBUG, &b)

	msg := "test 123 !#*"
	Debug(msg)
	if !strings.HasPrefix(b.String(), "DEBUG") {
		t.Fail()
	}
	if !strings.HasSuffix(b.String(), GetCaller(0)+" – "+msg+"\n") {
		t.Fail()
	}

	b.Reset()
	SetShowCaller(DEBUG, false)
	msg = "\thallo"
	Debug(msg)
	if !strings.HasPrefix(b.String(), "DEBUG") {
		t.Fail()
	}
	if !strings.HasSuffix(b.String(), msg+"\n") {
		t.Fail()
	}
	if strings.HasSuffix(b.String(), GetCaller(0)+" – "+msg+"\n") {
		t.Fail()
	}

	// reset
	SetShowCaller(DEBUG, true)
	SetOutput(DEBUG, os.Stdout)
	SetLevel(lvl)
}

func TestDebugf(t *testing.T) {
	lvl := LogLevel
	var b bytes.Buffer
	SetLevel(DEBUG)
	SetOutput(DEBUG, &b)

	msg := "test 123 !#*"
	Debugf("%s %d !#%c", "test", 123, '*')
	if !strings.HasPrefix(b.String(), "DEBUG") {
		t.Fail()
	}
	if !strings.HasSuffix(b.String(), GetCaller(0)+" – "+msg+"\n") {
		t.Fail()
	}

	b.Reset()
	SetShowCaller(DEBUG, false)
	msg = "\thallo"
	Debugf("\t%s", "hallo")
	if !strings.HasPrefix(b.String(), "DEBUG") {
		t.Fail()
	}
	if !strings.HasSuffix(b.String(), msg+"\n") {
		t.Fail()
	}
	if strings.HasSuffix(b.String(), GetCaller(0)+" – "+msg+"\n") {
		t.Fail()
	}

	// reset
	SetShowCaller(DEBUG, true)
	SetOutput(DEBUG, os.Stdout)
	SetLevel(lvl)
}

func TestInfo(t *testing.T) {
	lvl := LogLevel
	var b bytes.Buffer
	SetLevel(INFO)
	SetOutput(INFO, &b)
	SetShowCaller(INFO, true)

	msg := "test 123 !#*"
	Info(msg)
	if !strings.HasPrefix(b.String(), "INFO") {
		t.Fail()
	}
	if !strings.HasSuffix(b.String(), GetCaller(0)+" – "+msg+"\n") {
		t.Fail()
	}

	b.Reset()
	SetShowCaller(INFO, false)
	msg = "\thallo"
	Info(msg)
	if !strings.HasPrefix(b.String(), "INFO") {
		t.Fail()
	}
	if !strings.HasSuffix(b.String(), msg+"\n") {
		t.Fail()
	}
	if strings.HasSuffix(b.String(), GetCaller(0)+" – "+msg+"\n") {
		t.Fail()
	}

	// reset
	SetOutput(INFO, os.Stdout)
	SetLevel(lvl)
}

func TestInfof(t *testing.T) {
	lvl := LogLevel
	var b bytes.Buffer
	SetLevel(INFO)
	SetOutput(INFO, &b)
	SetShowCaller(INFO, true)

	msg := "test 123 !#*"
	Infof("%s %d !#%c", "test", 123, '*')
	if !strings.HasPrefix(b.String(), "INFO") {
		t.Fail()
	}
	if !strings.HasSuffix(b.String(), GetCaller(0)+" – "+msg+"\n") {
		t.Fail()
	}

	b.Reset()
	SetShowCaller(INFO, false)
	msg = "\thallo"
	Infof("\t%s", "hallo")
	if !strings.HasPrefix(b.String(), "INFO") {
		t.Fail()
	}
	if !strings.HasSuffix(b.String(), msg+"\n") {
		t.Fail()
	}
	if strings.HasSuffix(b.String(), GetCaller(0)+" – "+msg+"\n") {
		t.Fail()
	}

	// reset
	SetOutput(INFO, os.Stdout)
	SetLevel(lvl)
}

func TestWarn(t *testing.T) {
	lvl := LogLevel
	var b bytes.Buffer
	SetLevel(WARNING)
	SetOutput(WARNING, &b)
	SetShowCaller(WARNING, true)

	msg := "test 123 !#*"
	Warn(msg)
	if !strings.HasPrefix(b.String(), "WARN") {
		t.Fail()
	}
	if !strings.HasSuffix(b.String(), GetCaller(0)+" – "+msg+"\n") {
		t.Fail()
	}

	b.Reset()
	SetShowCaller(WARNING, false)
	msg = "\thallo"
	Warn(msg)
	if !strings.HasPrefix(b.String(), "WARN") {
		t.Fail()
	}
	if !strings.HasSuffix(b.String(), msg+"\n") {
		t.Fail()
	}
	if strings.HasSuffix(b.String(), GetCaller(0)+" – "+msg+"\n") {
		t.Fail()
	}

	// reset
	SetOutput(WARNING, os.Stdout)
	SetLevel(lvl)
}

func TestWarnf(t *testing.T) {
	lvl := LogLevel
	var b bytes.Buffer
	SetLevel(WARNING)
	SetOutput(WARNING, &b)
	SetShowCaller(WARNING, true)

	msg := "test 123 !#*"
	Warnf("%s %d !#%c", "test", 123, '*')
	if !strings.HasPrefix(b.String(), "WARN") {
		t.Fail()
	}
	if !strings.HasSuffix(b.String(), GetCaller(0)+" – "+msg+"\n") {
		t.Fail()
	}

	b.Reset()
	SetShowCaller(WARNING, false)
	msg = "\thallo"
	Warnf("\t%s", "hallo")
	if !strings.HasPrefix(b.String(), "WARN") {
		t.Fail()
	}
	if !strings.HasSuffix(b.String(), msg+"\n") {
		t.Fail()
	}
	if strings.HasSuffix(b.String(), GetCaller(0)+" – "+msg+"\n") {
		t.Fail()
	}

	// reset
	SetOutput(WARNING, os.Stdout)
	SetLevel(lvl)
}

func TestError(t *testing.T) {
	lvl := LogLevel
	var b bytes.Buffer
	SetLevel(ERROR)
	SetOutput(ERROR, &b)

	msg := "test 123 !#*"
	Error(msg)
	if !strings.HasPrefix(b.String(), "ERROR") {
		t.Fail()
	}
	if !strings.HasSuffix(b.String(), GetCaller(0)+" – "+msg+"\n") {
		t.Fail()
	}

	b.Reset()
	SetShowCaller(ERROR, false)
	msg = "\thallo"
	Error(msg)
	if !strings.HasPrefix(b.String(), "ERROR") {
		t.Fail()
	}
	if !strings.HasSuffix(b.String(), msg+"\n") {
		t.Fail()
	}
	if strings.HasSuffix(b.String(), GetCaller(0)+" – "+msg+"\n") {
		t.Fail()
	}

	// reset
	SetShowCaller(ERROR, true)
	SetOutput(ERROR, os.Stdout)
	SetLevel(lvl)
}

func TestErrorf(t *testing.T) {
	lvl := LogLevel
	var b bytes.Buffer
	SetLevel(ERROR)
	SetOutput(ERROR, &b)

	msg := "test 123 !#*"
	Errorf("%s %d !#%c", "test", 123, '*')
	if !strings.HasPrefix(b.String(), "ERROR") {
		t.Fail()
	}
	if !strings.HasSuffix(b.String(), GetCaller(0)+" – "+msg+"\n") {
		t.Fail()
	}

	b.Reset()
	SetShowCaller(ERROR, false)
	msg = "\thallo"
	Errorf("\t%s", "hallo")
	if !strings.HasPrefix(b.String(), "ERROR") {
		t.Fail()
	}
	if !strings.HasSuffix(b.String(), msg+"\n") {
		t.Fail()
	}
	if strings.HasSuffix(b.String(), GetCaller(0)+" – "+msg+"\n") {
		t.Fail()
	}

	// reset
	SetShowCaller(ERROR, true)
	SetOutput(ERROR, os.Stdout)
	SetLevel(lvl)
}

func TestFatal(t *testing.T) {
	lvl := LogLevel
	var b bytes.Buffer
	SetLevel(FATAL)
	SetOutput(FATAL, &b)

	msg := "test 123 !#*"
	Fatal(msg)
	if !strings.HasPrefix(b.String(), "FATAL") {
		t.Fail()
	}
	if !strings.HasSuffix(b.String(), GetCaller(0)+" – "+msg+"\n") {
		t.Fail()
	}

	b.Reset()
	SetShowCaller(FATAL, false)
	msg = "\thallo"
	Fatal(msg)
	if !strings.HasPrefix(b.String(), "FATAL") {
		t.Fail()
	}
	if !strings.HasSuffix(b.String(), msg+"\n") {
		t.Fail()
	}
	if strings.HasSuffix(b.String(), GetCaller(0)+" – "+msg+"\n") {
		t.Fail()
	}

	// reset
	SetShowCaller(FATAL, true)
	SetOutput(FATAL, os.Stdout)
	SetLevel(lvl)
}

func TestFatalf(t *testing.T) {
	lvl := LogLevel
	var b bytes.Buffer
	SetLevel(FATAL)
	SetOutput(FATAL, &b)

	msg := "test 123 !#*"
	Fatalf("%s %d !#%c", "test", 123, '*')
	if !strings.HasPrefix(b.String(), "FATAL") {
		t.Fail()
	}
	if !strings.HasSuffix(b.String(), GetCaller(0)+" – "+msg+"\n") {
		t.Fail()
	}

	b.Reset()
	SetShowCaller(FATAL, false)
	msg = "\thallo"
	Fatalf("\t%s", "hallo")
	if !strings.HasPrefix(b.String(), "FATAL") {
		t.Fail()
	}
	if !strings.HasSuffix(b.String(), msg+"\n") {
		t.Fail()
	}
	if strings.HasSuffix(b.String(), GetCaller(0)+" – "+msg+"\n") {
		t.Fail()
	}

	// reset
	SetShowCaller(FATAL, true)
	SetOutput(FATAL, os.Stdout)
	SetLevel(lvl)
}

func TestGetLogLine(t *testing.T) {
	tmp := showCaller

	logline := getLogLine(INFO, "test message")
	if !strings.HasPrefix(logline, styleFuncs[INFO]("INFO")) {
		t.Fail()
	}
	if !strings.HasSuffix(logline, "test message\n") {
		t.Fail()
	}

	logline = getLogLine(FATAL, "!# _123")
	if !strings.HasPrefix(logline, styleFuncs[FATAL]("FATAL")) {
		t.Fail()
	}
	if !strings.HasSuffix(logline, "!# _123\n") {
		t.Fail()
	}

	// reset
	showCaller = tmp
}
