package main

import "flag"

import "github.com/nilobject/conveyor/Godeps/_workspace/src/github.com/op/go-logging"

var log = logging.MustGetLogger("conveyor")

func main() {
	installMode := flag.Bool("install", false, "Launches database installer")
	serverMode := flag.Bool("server", false, "Launches the server")
	flag.Parse()

	if *installMode {

	} else if *serverMode {

	} else {
		log.Fatal("Needs to be launched in either install or server mode.")
	}
}
