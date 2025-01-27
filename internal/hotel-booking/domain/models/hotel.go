package models

// Hotel represents an enterprise-level hotel entity.
type Hotel struct {
	ID               string           `json:"id"`                // Unique identifier for the hotel.
	Name             string           `json:"name"`              // Hotel name.
	Brand            string           `json:"brand,omitempty"`   // Hotel brand or chain (optional).
	Location         Location         `json:"location"`          // Hotel location details.
	Rating           float64          `json:"rating"`            // Average rating of the hotel.
	Facilities       []string         `json:"facilities"`        // List of hotel facilities (e.g., Free WiFi, Pool, Gym).
	RoomTypes        []Room           `json:"room_types"`        // List of room types available at the hotel.
	Images           []string         `json:"images"`            // Hotel images (URLs).
	PriceRange       PriceRange       `json:"price_range"`       // Price range for rooms.
	Policies         Policies         `json:"policies"`          // Hotel policies.
	Availability     Availability     `json:"availability"`      // Room availability information.
	PaymentMethods   []string         `json:"payment_methods"`   // Accepted payment methods.
	ProviderMetadata ProviderMetadata `json:"provider_metadata"` // Metadata related to the external provider.
}

// Location represents hotel location details.
type Location struct {
	City       string  `json:"city"`        // City where the hotel is located.
	Country    string  `json:"country"`     // Country where the hotel is located.
	Latitude   float64 `json:"latitude"`    // Latitude of the hotel.
	Longitude  float64 `json:"longitude"`   // Longitude of the hotel.
	Address    string  `json:"address"`     // Physical address of the hotel.
	PostalCode string  `json:"postal_code"` // Postal code.
}

// Room represents the room details in a hotel.
type Room struct {
	ID           string   `json:"id"`           // Unique identifier for the room.
	Type         string   `json:"type"`         // Room type (e.g., Single, Double, Suite).
	Capacity     int      `json:"capacity"`     // Room capacity (number of people).
	Price        float64  `json:"price"`        // Price per night for the room.
	BedType      string   `json:"bed_type"`     // Type of bed (e.g., King, Queen).
	Availability bool     `json:"availability"` // Whether the room is available for booking.
	Images       []string `json:"images"`       // Room images (URLs).
	Facilities   []string `json:"facilities"`   // Room-specific facilities (e.g., Air Conditioning, TV, etc.).
}

// PriceRange represents the price range for the hotel.
type PriceRange struct {
	MinPrice float64 `json:"min_price"` // Minimum price for a room.
	MaxPrice float64 `json:"max_price"` // Maximum price for a room.
	Currency string  `json:"currency"`  // Currency for the price range.
}

// Policies represents hotel policies.
type Policies struct {
	Cancellation    string `json:"cancellation"`      // Cancellation policy.
	CheckInTime     string `json:"check_in_time"`     // Check-in time.
	CheckOutTime    string `json:"check_out_time"`    // Check-out time.
	SmokingPolicy   string `json:"smoking_policy"`    // Smoking policy (e.g., Smoking or Non-Smoking rooms).
	ChildPolicy     string `json:"child_policy"`      // Child policy (e.g., Children allowed).
	ExtraBedsPolicy string `json:"extra_beds_policy"` // Extra beds policy (e.g., Extra bed allowed).
}

// Availability represents the availability status of the hotel.
type Availability struct {
	AvailableRooms int `json:"available_rooms"` // Number of rooms available for booking.
	TotalRooms     int `json:"total_rooms"`     // Total number of rooms in the hotel.
}

// ProviderMetadata represents the metadata about the external hotel provider.
type ProviderMetadata struct {
	ProviderName   string  `json:"provider_name"`   // Name of the hotel provider (e.g., Booking.com, Expedia).
	ProviderID     string  `json:"provider_id"`     // Unique ID for the hotel in the provider's system.
	ProviderRating float64 `json:"provider_rating"` // Rating from the external provider.
	LastUpdated    string  `json:"last_updated"`    // Timestamp of the last update from the provider.
}
