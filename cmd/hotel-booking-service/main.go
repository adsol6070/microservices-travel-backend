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

	// hotelIDs := []string{"ALBLR275", "ALBLR545"}
	// adults := 1

	client := hotels.NewAmadeusClient("ldK8AEKr1ryNBhfpEMNkux4CwjydYqrX", "8DJFOdD0t7pbUQSf")

	// cityCode := "BLR" // Example: Bangalore city code
	// hotelsList, err := client.HotelSearch(cityCode)
	// if err != nil {
	// 	log.Fatalf("❌ Error fetching hotel list: %v", err)
	// }

	// fmt.Println("\n📌 **Hotels in the City**")
	// for _, hotel := range hotelsList {
	// 	fmt.Printf("🏨 **Hotel Name:** %s\n", hotel.Name)
	// 	fmt.Printf("   🏙 Country Code: %s\n", hotel.Address.CountryCode)
	// 	fmt.Printf("   🏙 Chain Code: %s\n", hotel.ChainCode)
	// 	fmt.Printf("   📍 Location: (%.5f, %.5f)\n", hotel.GeoCode.Latitude, hotel.GeoCode.Longitude)
	// 	fmt.Println("   -----------------------------------")
	// }

	// offers, err := client.FetchHotelOffers(hotelIDs, adults)
	// if err != nil {
	// 	log.Fatalf("❌ Error fetching hotel offers: %v", err)
	// }

	// fmt.Println("\n📌 **Available Hotel Offers**")
	// for _, hotel := range offers {
	// 	fmt.Printf("🏨 **Hotel:** %s (%s)\n", hotel.Hotel.Name, hotel.Hotel.CityCode)
	// 	fmt.Printf("   📍 Location: (%.5f, %.5f)\n", hotel.Hotel.Latitude, hotel.Hotel.Longitude)
	// 	fmt.Println("   -----------------------------------")
	// 	for _, roomOffer := range hotel.Offers {
	// 		fmt.Printf("   ✅ **Room Type:** %s\n", roomOffer.Room.TypeEstimated.Category)
	// 		fmt.Printf("   🛏 Beds: %d (%s)\n", roomOffer.Room.TypeEstimated.Beds, roomOffer.Room.TypeEstimated.BedType)
	// 		fmt.Printf("   💰 Price: %s %s\n", roomOffer.Price.Currency, roomOffer.Price.Total)
	// 		fmt.Printf("   📅 Check-in: %s | Check-out: %s\n", roomOffer.CheckInDate, roomOffer.CheckOutDate)
	// 		fmt.Printf("   📝 Description: %s\n", roomOffer.Room.Description.Text)
	// 		fmt.Println("   -----------------------------------")
	// 	}
	// }

	fmt.Println("\n📌 **Testing Hotel Booking**")
	bookingRequest := `{
    "data": {
        "type": "hotel-order",
        "guests": [
            {
                "tid": 1,
                "title": "MR",
                "firstName": "BOB",
                "lastName": "SMITH",
                "phone": "+33679278416",
                "email": "bob.smith@email.com"
            }
        ],
        "travelAgent": {
            "contact": {
                "email": "bob.smith@email.com"
            }
        },
        "roomAssociations": [
            {
                "guestReferences": [
                    {
                        "guestReference": "1"
                    }
                ],
                "hotelOfferId": "IKWAU9X7IY"
            }
        ],
        "payment": {
            "method": "CREDIT_CARD",
            "paymentCard": {
                "paymentCardInfo": {
                    "vendorCode": "VI",
                    "cardNumber": "4151289722471370",
                    "expiryDate": "2026-08",
                    "holderName": "BOB SMITH"
                }
            }
        }
    }
}`

	bookingResponse, err := client.CreateHotelBooking([]byte(bookingRequest))
	if err != nil {
		log.Fatalf("❌ Error creating hotel booking: %v", err)
	} else {
		fmt.Printf("✅ Booking Successful! Confirmation: %v\n", bookingResponse)
	}

	router := mux.NewRouter()

	// Set Port and Start Server
	serverPort := "5100"
	logger.Printf("Starting server on port %s...\n", "5100")
	if err := http.ListenAndServe(":"+serverPort, router); err != nil {
		logger.Fatalf("Failed to start server: %v", err)
	}

}
