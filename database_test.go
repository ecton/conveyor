package main

import "testing"

func TestDestroy(*testing.T) {
	migrateDatabaseDown()
	migrateDatabaseUp()
}
