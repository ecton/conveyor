package main

import (
	"encoding/json"
	"os"
)

// PostgreSQLConfiguration contains the options to control how conveyor should
// connect to the database
type PostgreSQLConfiguration struct {
	host         string
	username     string
	password     string
	databaseName string
}

// Configuration structure for conveyor
// Contains all the settings needed to run the server
type Configuration struct {
	pg PostgreSQLConfiguration
}

func loadConfiguration(path string) {
	file, err := os.Open(path)
	if err != nil {
		logger().Fatalf("Could not open file %v, error %v", path, err)
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&configuration)
	if err != nil {
		logger().Fatalf("Could not parse %v, error %v", path, err)
	}
}

var configuration Configuration

func initializeConfiguration() {
	if _, err := os.Stat("/etc/conveyor.json"); err == nil {
		loadConfiguration("/etc/conveyor.json")
	} else if _, err := os.Stat("conveyor.json"); err == nil {
		loadConfiguration("conveyor.json")
	} else {
		// Default configuaration options
		configuration = Configuration{
			pg: PostgreSQLConfiguration{
				username: "testuser",
				password: "testing",
				host:     "localhost",
			},
		}
	}
}
