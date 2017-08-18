package places

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldReturnAnObject(t *testing.T) {
	assert := assert.New(t)

	actual, _ := GetPlacesService("41.584627,-93.637551")

	assert.NotNil(actual)
}

func TestShouldReturnAtLeastOneResult(t *testing.T) {
	assert := assert.New(t)

	actual, _ := GetPlacesService("41.584627,-93.637551")

	assert.NotNil(actual[0].Name)
	assert.NotNil(actual[0].FoodType)
	assert.NotNil(actual[0].Price)
	assert.NotNil(actual[0].Rating)
	assert.NotNil(actual[0].Distance)
	assert.NotNil(actual[0].ImageUrl)
}
