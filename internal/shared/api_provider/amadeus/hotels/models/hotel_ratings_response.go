package models

type HotelSentimentResponse struct {
	Meta     MetaResponse         `json:"meta"`
	Data     []HotelDataResponse  `json:"data"`
	Warnings []WarningResponse   `json:"warnings,omitempty"`
}

type HotelDataResponse struct {
	HotelID         string     `json:"hotelId"`
	OverallRating   int        `json:"overallRating"`
	NumberOfReviews int        `json:"numberOfReviews"`
	NumberOfRatings int        `json:"numberOfRatings"`
	Type            string     `json:"type"`
	Sentiments      Sentiments `json:"sentiments"`
}

type Sentiments struct {
	Staff          int `json:"staff"`
	Location       int `json:"location"`
	Service        int `json:"service"`
	RoomComforts   int `json:"roomComforts"`
	SleepQuality   int `json:"sleepQuality"`
	SwimmingPool   int `json:"swimmingPool"`
	ValueForMoney  int `json:"valueForMoney"`
	Facilities     int `json:"facilities"`
	Catering       int `json:"catering"`
	PointsOfInterest int `json:"pointsOfInterest"`
}

type MetaResponse struct {
	Count int   `json:"count"`
	Links LinksResponse `json:"links"`
}

type LinksResponse struct {
	Self string `json:"self"`
}

type WarningResponse struct {
	Code   int    `json:"code"`
	Title  string `json:"title"`
	Source Source `json:"source"`
	Detail string `json:"detail"`
}

type Source struct {
	Parameter string `json:"parameter"`
	Pointer   string `json:"pointer"`
}
