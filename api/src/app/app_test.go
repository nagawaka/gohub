package main

import (
	"testing"
	"os"
)

func TestMain(m *testing.M) {
	// runs
	os.Exit(m.Run())
}