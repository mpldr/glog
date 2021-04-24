package glog

import (
	"os"
	"testing"
)

func TestAddLogFile(t *testing.T) {
	file, err := AddLogFile("test_AddLogFile_Warning.txt", WARNING)
	if err != nil {
		t.Log(err)
		t.Fail()
	}

	t.Cleanup(func() {
		SetOutput(WARNING, os.Stdout)
		file.Close()
		os.Remove("test_AddLogFile_Warning.txt")
	})

	if len(outputs[WARNING]) != 2 {
		t.Fail()
	}
	if _, err = os.Stat("test_AddLogFile_Warning.txt"); os.IsNotExist(err) {
		t.Errorf("AddLogFile did not create file 'test_AddLogFile_Warning.txt'")
	}
}
