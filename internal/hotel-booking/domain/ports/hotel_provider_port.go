package ports

type HotelProvider interface {
	GetHotels() ([]map[string]interface{}, error)
}
