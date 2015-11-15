package main

import "flag"

func main() {
	serverMode := flag.Bool("server", false, "Launches the server")
	initializeConfiguration()
	migrateDatabaseUp()
	flag.Parse()

	if *serverMode {

	} else {
		logger().Fatal("Needs to be launched in either install or server mode.")
	}
}
