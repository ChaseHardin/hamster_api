package places

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldReturnAnObject(t *testing.T) {
	assert := assert.New(t)

	actual, _ := GetPlacesService()

	assert.NotNil(actual)
}

func TestShouldReturnAtLeastOneResult(t *testing.T) {
	assert := assert.New(t)

	actual, _ := GetPlacesService()

	assert.NotNil(actual[0].Name)
	assert.NotNil(actual[0].FoodType)
	assert.NotNil(actual[0].Price)
	assert.NotNil(actual[0].Rating)
	assert.NotNil(actual[0].Distance)
	assert.NotNil(actual[0].ImageUrl)
}

// func TestShouldReturnAListOfTwoPlaces(t *testing.T) {
// 	assert := assert.New(t)
// 	mockAModel := map[int]Place{
// 		0: {"Malo", "Mexican", 3, 3.4, 1.2, "imageUrl"},
// 		1: {"China One", "Chinese", 2, 1.2, 1.2, "imageUrl"}}

// 	actual := GetPlacesService("00002")

// 	assert.Equal(2, len(actual))
// 	assert.Equal(actual[0].Name, mockAModel[0].Name)
// 	assert.Equal(actual[0].FoodType, mockAModel[0].FoodType)
// 	assert.Equal(actual[0].Price, mockAModel[0].Price)
// 	assert.Equal(actual[0].Rating, mockAModel[0].Rating)
// 	assert.Equal(actual[0].Distance, mockAModel[0].Distance)
// 	assert.Equal(actual[0].ImageUrl, mockAModel[0].ImageUrl)

// 	assert.Equal(actual[1].Name, mockAModel[1].Name)
// 	assert.Equal(actual[1].FoodType, mockAModel[1].FoodType)
// 	assert.Equal(actual[1].Price, mockAModel[1].Price)
// 	assert.Equal(actual[1].Rating, mockAModel[1].Rating)
// 	assert.Equal(actual[1].Distance, mockAModel[1].Distance)
// 	assert.Equal(actual[1].ImageUrl, mockAModel[1].ImageUrl)
// }
