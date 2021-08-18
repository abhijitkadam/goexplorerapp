package main

import "fmt"

//https://blog.golang.org/json
//https://pkg.go.dev/encoding/json#Marshal

//omitempty -> use when you don't want to include anyfield in json
//The problem with omitempty is it takes the default value I guess.
//The solution is to use pointer. So empty value of pointer will be nil.

func main() {
	hotels, err := FetchHotelsFromUrlUsingJsonDecoder("http://localhost:3000/inventory")

	if err != nil {
		panic(err)
	}

	fhotels := Hotels(hotels).FilterHotelByRoomType("LARGE")
	fmt.Println(fhotels)

	shotels := Hotels(hotels).FilterByFn(func(h *Hotel) bool {
		return h.RoomType == "SMALL"
	})

	fmt.Println(shotels)

	filterFn := func(h *Hotel) bool {
		return h.RoomType == "X-LARGE"
	}

	xlhotels := Hotels(hotels).FilterBy(filterFn)
	fmt.Println(xlhotels)

}

func MethodDemo() {
	hotel := NewHotel("gold", "HOTEL-DDF", "SMALL", 20.0)

	p := hotel.GetPrice()
	fmt.Println(p)

	//if use the func instead of method
	//SetPrice(hotel, 55)

	hotel.SetPrice(50)
	fmt.Println("New Price: ", hotel.GetPrice())

	//SetHotelPrice(&hotel, 65)

	hotel.SetHotelPrice(60)
	fmt.Println("After SetHotelPrice: ", hotel.GetPrice())
}

func SlicesDemo() {
	//GoSlices()
	//GoSlicesRange()

	//err in go
	//Always check for error first.
	hotels, err := FetchHotelsFromUrlUsingJsonDecoder("http://localhost:3000/inventory")

	if err != nil {
		panic(err)
	}

	fmt.Println(hotels)
}
