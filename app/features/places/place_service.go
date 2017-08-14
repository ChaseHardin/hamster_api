package places

func GetPlacesService() (data []Place) {
	data = append(data, Place{
		"Malo",
		"Mexican",
		3,
		3.4,
		1.2,
		"imageUrl"})
	return
}