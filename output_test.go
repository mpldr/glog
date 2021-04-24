package glog

import (
	"bytes"
	"log"
	"os"
	"testing"
)

var (
	tmp bool = EnableMetaLogging
	b   bytes.Buffer
)

// enable metalogging to enable testing methods by checking if they printed something using metalog()
func setupForInvalidValues() {
	b = bytes.Buffer{}
	log.SetOutput(&b)
	tmp = EnableMetaLogging
	EnableMetaLogging = true
}

// reset EnableMetaLogging to the initial value and reset outputs from all levels
func reset() {
	EnableMetaLogging = tmp
	log.SetOutput(os.Stdout)
	SetOutputBetween(TRACE, FATAL, os.Stdout)
}

func TestSetupForInvalidValues(t *testing.T) {
	setupForInvalidValues()
	t.Cleanup(reset)

	if !EnableMetaLogging {
		t.Fail()
	}
}

func TestReset(t *testing.T) {
	reset()

	if EnableMetaLogging != tmp {
		t.Fail()
	}

	for lvl := TRACE; lvl <= FATAL; lvl++ {
		if outputs[lvl][0] != os.Stdout || len(outputs[lvl]) != 1 {
			t.Errorf("Resetting outputs of level %s failed", lvl.String())
		}
	}
}

func TestSetOutput(t *testing.T) {
	SetOutput(TRACE, os.Stderr)
	if outputs[TRACE][0] != os.Stderr {
		t.Fail()
	}

	SetOutput(INFO, os.Stderr)
	if outputs[INFO][0] != os.Stderr {
		t.Fail()
	}

	setupForInvalidValues()
	t.Cleanup(reset)

	SetOutput(23, os.Stdout)
	if len(b.String()) > 0 {
		t.Errorf("Did not return early from SetOutput() although the given level (23) is not valid ")
	}
}

func TestAddOutput(t *testing.T) {
	AddOutput(INFO, os.Stderr)
	if len(outputs[INFO]) != 2 {
		t.Fail()
	}
	if outputs[INFO][1] != os.Stderr {
		t.Fail()
	}

	AddOutput(WARNING, os.Stderr)
	if len(outputs[WARNING]) != 2 {
		t.Fail()
	}
	if outputs[WARNING][1] != os.Stderr {
		t.Fail()
	}

	setupForInvalidValues()
	t.Cleanup(reset)

	AddOutput(23, os.Stdout)
	if len(b.String()) > 0 {
		t.Errorf("Did not return early from AddOutput() although the given level (23) is not valid ")
	}
}

func TestSetOutputBetween(t *testing.T) {
	SetOutputBetween(TRACE, INFO, os.Stderr)
	for lvl := TRACE; lvl <= INFO; lvl++ {
		if outputs[lvl][0] != os.Stderr {
			t.Fail()
		}
	}

	SetOutputBetween(DEBUG, FATAL, os.Stderr)
	for lvl := DEBUG; lvl <= FATAL; lvl++ {
		if outputs[lvl][0] != os.Stderr {
			t.Fail()
		}
	}

	setupForInvalidValues()
	t.Cleanup(reset)

	SetOutputBetween(ERROR, 23, os.Stdout)
	if len(b.String()) > 0 {
		t.Errorf("Did not return early from SetOutputBetween() although the given upper level (23) is not valid ")
	}

	b.Reset()
	SetOutputBetween(14, TRACE, os.Stdout)
	if len(b.String()) > 0 {
		t.Errorf("Did not return early from SetOutputBetween() although the given lower level (14) is not valid ")
	}

	b.Reset()
	SetOutputBetween(14, 69, os.Stdout)
	if len(b.String()) > 0 {
		t.Errorf("Did not return early from SetOutputBetween() although both given levels (14, 69) are not valid ")
	}
}

func TestAddOutputBetween(t *testing.T) {
	AddOutputBetween(INFO, ERROR, os.Stderr)
	for lvl := INFO; lvl <= ERROR; lvl++ {
		if len(outputs[lvl]) != 2 {
			t.Fail()
			continue
		}
		if outputs[lvl][1] != os.Stderr {
			t.Fail()
		}
	}

	AddOutputBetween(TRACE, DEBUG, os.Stderr)
	for lvl := TRACE; lvl <= DEBUG; lvl++ {
		if len(outputs[lvl]) != 2 {
			t.Fail()
			continue
		}
		if outputs[lvl][1] != os.Stderr {
			t.Fail()
		}
	}

	setupForInvalidValues()
	t.Cleanup(reset)

	AddOutputBetween(ERROR, 23, os.Stdout)
	if len(b.String()) > 0 {
		t.Errorf("Did not return early from AddOutputBetween() although the given upper level (23) is not valid ")
	}

	b.Reset()
	AddOutputBetween(14, TRACE, os.Stdout)
	if len(b.String()) > 0 {
		t.Errorf("Did not return early from AddOutputBetween() although the given lower level (14) is not valid ")
	}

	b.Reset()
	AddOutputBetween(14, 69, os.Stdout)
	if len(b.String()) > 0 {
		t.Errorf("Did not return early from AddOutputBetween() although both given levels (14, 69) are not valid ")
	}
}

func TestSetOutputAbove(t *testing.T) {
	SetOutputAbove(TRACE, os.Stderr)
	for lvl := TRACE; lvl <= FATAL; lvl++ {
		if len(outputs[lvl]) != 1 {
			t.Fail()
			continue
		}
		if outputs[lvl][0] != os.Stderr {
			t.Fail()
		}
	}

	SetOutputAbove(WARNING, os.Stdout)
	for lvl := WARNING; lvl <= FATAL; lvl++ {
		if len(outputs[lvl]) != 1 {
			t.Fail()
			continue
		}
		if outputs[lvl][0] != os.Stdout || len(outputs[lvl]) != 1 {
			t.Fail()
		}
	}

	setupForInvalidValues()
	t.Cleanup(reset)

	SetOutputAbove(23, os.Stdout)
	if len(b.String()) > 0 {
		t.Errorf("Did not return early from SetOutputAbove() although the given level (23) is not valid ")
	}
}

func TestAddOutputAbove(t *testing.T) {
	AddOutputAbove(TRACE, os.Stderr)
	for lvl := TRACE; lvl <= FATAL; lvl++ {
		if len(outputs[lvl]) != 2 {
			t.Fail()
			continue
		}
		if outputs[lvl][1] != os.Stderr {
			t.Fail()
		}
	}

	AddOutputAbove(ERROR, os.Stdout)
	for lvl := ERROR; lvl <= FATAL; lvl++ {
		if len(outputs[lvl]) != 3 {
			t.Fail()
			continue
		}
		if outputs[lvl][2] != os.Stdout {
			t.Fail()
		}
	}

	setupForInvalidValues()
	t.Cleanup(reset)

	AddOutputAbove(23, os.Stdout)
	if len(b.String()) > 0 {
		t.Errorf("Did not return early from AddOutputAbove() although the given level (23) is not valid ")
	}
}

func TestSetOutputBelow(t *testing.T) {
	SetOutputBelow(DEBUG, os.Stderr)
	for lvl := TRACE; lvl <= DEBUG; lvl++ {
		if len(outputs[lvl]) != 1 {
			t.Fail()
			continue
		}
		if outputs[lvl][0] != os.Stderr {
			t.Fail()
		}
	}

	SetOutputBelow(FATAL, os.Stderr)
	for lvl := TRACE; lvl <= FATAL; lvl++ {
		if len(outputs[lvl]) != 1 {
			t.Fail()
			continue
		}
		if outputs[lvl][0] != os.Stderr {
			t.Fail()
		}
	}

	setupForInvalidValues()
	t.Cleanup(reset)

	SetOutputBelow(23, os.Stdout)
	if len(b.String()) > 0 {
		t.Errorf("Did not return early from SetOutputBelow() although the given level (23) is not valid ")
	}
}

func TestAddOutputBelow(t *testing.T) {
	AddOutputBelow(ERROR, os.Stderr)
	for lvl := TRACE; lvl <= ERROR; lvl++ {
		if len(outputs[lvl]) != 2 {
			t.Fail()
			continue
		}
		if outputs[lvl][1] != os.Stderr {
			t.Fail()
		}
	}

	AddOutputBelow(DEBUG, os.Stdout)
	for lvl := TRACE; lvl <= DEBUG; lvl++ {
		if len(outputs[lvl]) != 3 {
			t.Fail()
			continue
		}
		if outputs[lvl][2] != os.Stdout {
			t.Fail()
		}
	}

	setupForInvalidValues()
	t.Cleanup(reset)

	AddOutputBelow(23, os.Stdout)
	if len(b.String()) > 0 {
		t.Errorf("Did not return early from AddOutputBelow() although the given level (23) is not valid ")
	}
}

func TestIsTerminal(t *testing.T) {
	if !isTerminal(outputs[INFO][0]) {
		t.Fail()
	}

	SetOutput(INFO, os.Stderr)
	if !isTerminal(outputs[INFO][0]) {
		t.Fail()
	}

	file, err := AddLogFile("test_isTerminal.txt", INFO)
	if err != nil {
		t.Log(err)
		t.Skip()
	}

	// delete test file in cleanup
	t.Cleanup(func() {
		file.Close()
		SetOutput(INFO, os.Stdout)
		os.Remove("test_isTerminal.txt")
	})

	if isTerminal(outputs[INFO][1]) {
		t.Fail()
	}
}
