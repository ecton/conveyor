package main

import "flag"

func main() {
	installMode := flag.Bool("install", false, "Launches database installer")
	serverMode := flag.Bool("server", false, "Launches the server")
	flag.Parse()

	if *installMode {

	} else if *serverMode {

	} else {
		logger().Fatal("Needs to be launched in either install or server mode.")
	}
}
