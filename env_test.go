package glog

import (
	"os"
	"testing"
)

func TestEnvEnableMetalogger(t *testing.T) {
	origval := os.Getenv("GLOG_METALOGGER")
	err := os.Setenv("GLOG_METALOGGER", "1")
	if err != nil {
		t.Logf("cannot setup env: %v", err)
		t.SkipNow()
	}
	t.Cleanup(func() { os.Setenv("GLOG_METALOGGER", origval) })

	setupEnv()
	t.Cleanup(func() { EnableMetaLogging = false })

	if !EnableMetaLogging {
		t.Errorf("metalogger was not enabled")
	}
}

func TestEnvSetColor(t *testing.T) {
	expectedValues := map[string]int8{
		"1": 1, "ON": 1, "ALWAYS": 1,
		"on": 1, "AlwAyS": 1,
		"-1": -1, "0": -1, "OFF": -1, "NEVER": -1,
		"off": -1, "NevER": -1,
		"ABRA": 0, "Bonk": 0,
	}

	for env, res := range expectedValues {
		t.Run("value_"+env, func(st *testing.T) {
			origval := os.Getenv("GLOG_COLOR")
			err := os.Setenv("GLOG_COLOR", env)
			if err != nil {
				t.Logf("cannot setup env: %v", err)
				t.SkipNow()
			}
			t.Cleanup(func() { os.Setenv("GLOG_COLOR", origval) })

			oc := OverrideColor
			t.Cleanup(func() { OverrideColor = oc })
			setupEnv()

			if OverrideColor != res {
				t.Errorf("did not set color to expected value. value '%s' lead to value '%d', but '%d' was expected", env, OverrideColor, res)
			}
		})
	}
}

func TestEnvSetLevel(t *testing.T) {
	expectedValues := map[string]Level{
		"TRACE": TRACE, "VERBOSE": TRACE,
		"DEBUG": DEBUG,
		"INFO":  INFO,
		"WARN":  WARNING, "WARNING": WARNING,
		"ERROR": ERROR,
		"FATAL": FATAL,
		"MUTE":  Level(42), "SILENT": Level(42),
	}

	for env, res := range expectedValues {
		t.Run("value_"+env, func(st *testing.T) {
			origval := os.Getenv("GLOG_LEVEL")
			err := os.Setenv("GLOG_LEVEL", env)
			if err != nil {
				t.Logf("cannot setup env: %v", err)
				t.SkipNow()
			}
			t.Cleanup(func() { os.Setenv("GLOG_LEVEL", origval) })

			lvl := LogLevel
			t.Cleanup(func() { LogLevel = lvl })
			setupEnv()
			t.Cleanup(func() { levelSetFromEnv = false })

			if LogLevel != res {
				t.Errorf("did not set level to expected value. value '%s' lead to value '%d', but '%d' was expected", env, LogLevel, res)
			}

			if !levelSetFromEnv {
				t.Error("loglevel was not marked as set from env.")
			}
		})
	}
}
