package places

import "testing"

func TestShouldReturnPlaces(t *testing.T) {
	mock_a_model := map[int]Place{0: {"Malo", "Mexican", 3, 3.4, 1.2, "imageUrl"}}

	actual := GetPlaces()

	if actual[0].Name != mock_a_model[0].Name {
		t.Errorf("Invalid Name property. Expected: %v Actual: %v", actual[0].Name, mock_a_model[0].Name)
	}

	if actual[0].FoodType != mock_a_model[0].FoodType {
		t.Errorf("Invalid FoodType property. Expected: %v Actual: %v", actual[0].FoodType, mock_a_model[0].FoodType)
	}

	if actual[0].Price != mock_a_model[0].Price {
		t.Errorf("Invalid Price property. Expected: %v Actual: %v", actual[0].Price, mock_a_model[0].Price)
	}

	if actual[0].Rating != mock_a_model[0].Rating {
		t.Errorf("Invalid Rating property. Expected: %v Actual: %v", actual[0].Rating, mock_a_model[0].Rating)
	}

	if actual[0].Distance != mock_a_model[0].Distance {
		t.Errorf("Invalid Distance property. Expected: %v Actual: %v", actual[0].Distance, mock_a_model[0].Distance)
	}

	if actual[0].ImageUrl != mock_a_model[0].ImageUrl {
		t.Errorf("Invalid ImageUrl property. Expected: %v Actual: %v", actual[0].ImageUrl, mock_a_model[0].ImageUrl)
	}
}