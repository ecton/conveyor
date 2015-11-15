package main

import "testing"
import "github.com/stretchr/testify/assert"

func TestInfrastructureBasics(t *testing.T) {
	original := Infrastructure{Name: "Home"}
	original.Save()
	assert.NotEqual(t, 0, original.ID, "ID was not returned after insertion")

	// Re-request the data to make sure GetInfrastructure works
	second, err := GetInfrastructure(original.ID)
	assert.NoError(t, err, "Error retrieving inserted infrastructure")
	assert.Equal(t, original, original, "Infrastructure returned was not the one asked for")

	// Update the name
	second.Name = "Office"
	err = second.Save()
	assert.NoError(t, err, "Error updating Infrastructure")
	third, err := GetInfrastructure(second.ID)
	assert.NoError(t, err, "Error retrieving updated infrastructure")
	assert.Equal(t, second, third, "Infrastructure returned was not as expected")

	// Delete the infrastructure
	err = third.Delete()
	assert.NoError(t, err, "Error deleting infrastructure")
	assert.Equal(t, int64(0), third.ID, "Deleting did not set ID to 0")
	err = third.Delete()
	assert.Error(t, err, "No error deleting non-existent infrastructure")
	_, err = GetInfrastructure(0)
	assert.Error(t, err, "No error querying for non-existent infrastructure")
}
