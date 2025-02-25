package googlePlaceModels

type PlacesResponse struct {
	Places []Place `json:"places"`
}

type Place struct {
	ID          string      `json:"id"`
	DisplayName DisplayName `json:"displayName"`
	Photos      []Photo     `json:"photos"`
	Rating      float64     `json:"rating,omitempty"`
}

type DisplayName struct {
	Text         string `json:"text"`
	LanguageCode string `json:"languageCode"`
}

type Photo struct {
	Name               string              `json:"name"`
	WidthPx            int                 `json:"widthPx"`
	HeightPx           int                 `json:"heightPx"`
	AuthorAttributions []AuthorAttribution `json:"authorAttributions"`
	FlagContentURI     string              `json:"flagContentUri"`
	GoogleMapsURI      string              `json:"googleMapsUri"`
}

type AuthorAttribution struct {
	DisplayName string `json:"displayName"`
	URI         string `json:"uri"`
	PhotoURI    string `json:"photoUri"`
}

type TextQueryRequest struct {
	TextQuery    string       `json:"textQuery"`
	LocationBias LocationBias `json:"locationBias"`
}

type LocationBias struct {
	Circle Circle `json:"circle"`
}

type Circle struct {
	Center Coordinates `json:"center"`
	Radius float64     `json:"radius"`
}

type Coordinates struct {
	Latitude  *float64 `json:"latitude"`
	Longitude *float64 `json:"longitude"`
}
