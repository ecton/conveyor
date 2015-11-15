package main

import (
	"flag"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	flag.Parse()
	initializeConfiguration()
	os.Exit(m.Run())
}
