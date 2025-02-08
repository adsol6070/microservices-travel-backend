package main

import (
	"fmt"
	"log"
	"microservices-travel-backend/internal/shared/api_provider/amadeus/hotels"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	logger := log.New(os.Stdout, "hotel-service: ", log.LstdFlags|log.Lshortfile)

	hotelIDs := []string{"ALBLR275", "ALBLR545"}
	adults := 1

	client := hotels.NewAmadeusClient("", "")
	offers, err := client.FetchHotelOffers(hotelIDs, adults)
	if err != nil {
		log.Fatalf("❌ Error fetching hotel offers: %v", err)
	}

	fmt.Println("\n📌 **Available Hotel Offers**\n")
	for _, hotel := range offers {
		fmt.Printf("🏨 **Hotel:** %s (%s)\n", hotel.Hotel.Name, hotel.Hotel.CityCode)
		fmt.Printf("   📍 Location: (%.5f, %.5f)\n", hotel.Hotel.Latitude, hotel.Hotel.Longitude)
		fmt.Println("   -----------------------------------")
		for _, roomOffer := range hotel.Offers {
			fmt.Printf("   ✅ **Room Type:** %s\n", roomOffer.Room.TypeEstimated.Category)
			fmt.Printf("   🛏 Beds: %d (%s)\n", roomOffer.Room.TypeEstimated.Beds, roomOffer.Room.TypeEstimated.BedType)
			fmt.Printf("   💰 Price: %s %s\n", roomOffer.Price.Currency, roomOffer.Price.Total)
			fmt.Printf("   📅 Check-in: %s | Check-out: %s\n", roomOffer.CheckInDate, roomOffer.CheckOutDate)
			fmt.Printf("   📝 Description: %s\n", roomOffer.Room.Description.Text)
			fmt.Println("   -----------------------------------")
		}
	}

	router := mux.NewRouter()

	// Set Port and Start Server
	serverPort := "5100"
	logger.Printf("Starting server on port %s...\n", "5100")
	if err := http.ListenAndServe(":"+serverPort, router); err != nil {
		logger.Fatalf("Failed to start server: %v", err)
	}

}
