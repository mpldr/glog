package glog

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestLevelsBelow(t *testing.T) {
	levels := levelsBelow(FATAL)
	if len(levels) != 6 {
		t.Fail()
	}

	levels = levelsBelow(DEBUG)
	if len(levels) != 2 {
		t.Fail()
	}

	levels = levelsBelow(TRACE)
	if len(levels) != 1 {
		t.Fail()
	}

	levels = levelsBelow(42)
	if len(levels) != 0 {
		t.Fail()
	}
}

func TestLevelsAbove(t *testing.T) {
	levels := levelsAbove(FATAL)
	if len(levels) != 1 {
		t.Fail()
	}

	levels = levelsAbove(DEBUG)
	if len(levels) != 5 {
		t.Fail()
	}

	levels = levelsAbove(TRACE)
	if len(levels) != 6 {
		t.Fail()
	}

	levels = levelsAbove(42)
	if len(levels) != 0 {
		t.Fail()
	}
}

func TestLevelsBetween(t *testing.T) {
	levels := levelsBetween(DEBUG, FATAL)
	if len(levels) != 5 {
		t.Errorf("Expexted %d results, but got %d for input (%d,%d)", 6, len(levels), DEBUG, FATAL)
	}

	levels = levelsBetween(FATAL, TRACE)
	if len(levels) != 6 {
		t.Errorf("Expexted %d results, but got %d for input (%d,%d)", 6, len(levels), FATAL, TRACE)
	}

	levels = levelsBetween(DEBUG, WARNING)
	if len(levels) != 3 {
		t.Errorf("Expexted %d results, but got %d for input (%d,%d)", 3, len(levels), DEBUG, WARNING)
		t.Fail()
	}

	levels = levelsBetween(DEBUG, DEBUG)
	if len(levels) != 1 {
		t.Errorf("Expexted %d results, but got %d for input (%d,%d)", 1, len(levels), DEBUG, DEBUG)
	}

	levels = levelsBetween(42, FATAL)
	if len(levels) != 0 {
		t.Errorf("Expexted %d results, but got %d for input (%d,%d)", 0, len(levels), 42, FATAL)
	}

	levels = levelsBetween(DEBUG, 42)
	if len(levels) != 0 {
		t.Errorf("Expexted %d results, but got %d for input (%d,%d)", 0, len(levels), DEBUG, 42)
	}
}

func TestStringer(t *testing.T) {
	if fmt.Sprint(TRACE) != "TRACE" {
		t.Errorf("TRACE returned %s", fmt.Sprint(TRACE))
	}

	if fmt.Sprint(DEBUG) != "DEBUG" {
		t.Errorf("DEBUG returned %s", fmt.Sprint(DEBUG))
	}

	if fmt.Sprint(INFO) != "INFO" {
		t.Errorf("INFO returned %s", fmt.Sprint(INFO))
	}

	if fmt.Sprint(WARNING) != "WARNING" {
		t.Errorf("WARNING returned %s", fmt.Sprint(WARNING))
	}

	if fmt.Sprint(ERROR) != "ERROR" {
		t.Errorf("ERROR returned %s", fmt.Sprint(ERROR))
	}

	if fmt.Sprint(FATAL) != "FATAL" {
		t.Errorf("FATAL returned %s", fmt.Sprint(FATAL))
	}

	for i := 0; i < 100; i++ {
		rnd := uint8(rand.Intn(255))
		if !isValid(Level(rnd)) && fmt.Sprint(Level(rnd)) != "ID-10T" {
			t.Errorf("did not provide appropriate response on %d", rnd)
		}
	}
}
