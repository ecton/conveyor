package main

import "flag"

func main() {
	serverMode := flag.Bool("server", false, "Launches the server")
	recreateDb := flag.Bool("destroy", false, "Runs all down migrations then up migrations")
	initializeConfiguration()
	flag.Parse()
	if *recreateDb {
		migrateDatabaseDown()
	}
	migrateDatabaseUp()

	if *serverMode {

	} else {
		logger().Fatal("Needs to be launched in either install or server mode.")
	}
}
