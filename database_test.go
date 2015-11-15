package main

import "testing"
import "github.com/stretchr/testify/assert"

func TestDestroy(t *testing.T) {
	err := migrateDatabaseDown()
	assert.NoError(t, err, "Error migrating down")
	err = migrateDatabaseUp()
	assert.NoError(t, err, "Error migrating up")
}

func TestMigrationWithBadInfo(t *testing.T) {
	oldPort := config().pg.port
	config().pg.port = 1
	err := migrateDatabaseDown()
	assert.Error(t, err, "No error with bad connection string")
	config().pg.port = oldPort
}
