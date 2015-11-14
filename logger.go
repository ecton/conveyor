package main

import "github.com/nilobject/conveyor/Godeps/_workspace/src/github.com/op/go-logging"

func logger() *logging.Logger {
	return logging.MustGetLogger("main")
}
