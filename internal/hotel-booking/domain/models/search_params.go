package models

type SearchParams struct {
	Location     string   `json:"location"`
	CheckIn      string   `json:"check_in"`
	CheckOut     string   `json:"check_out"`
	Guests       int      `json:"guests"`
	Rooms        int      `json:"rooms"`
	PriceMin     int      `json:"price_min"`
	PriceMax     int      `json:"price_max"`
	Rating       int      `json:"rating"`
	Amenities    []string `json:"amenities"`
	SortOrder    string   `json:"sort_order"`
	Currency     string   `json:"currency"`
	Language     string   `json:"language"`
	Children     int      `json:"children"`
	ChildrenAges []int    `json:"children_ages"`
	HotelName    string   `json:"hotel_name"`
	StarRating   int      `json:"star_rating"`
}
