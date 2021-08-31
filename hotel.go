package main

/*
{
	source: "gold",
	name: "Hotel-AB44",
	room-type: "LARGE",
	price: "200"
}
*/
//json tags tell the json Marshal & unmarshall methods how to convert the fields to & from json
//like RoomType would be Marshalled to room-type
//and Price would be marshalled to price and since it is string in json and float32 in code it needs conversion
type Hotel struct {
	Source   string  `json:"source"`
	Name     string  `json:"name"`
	RoomType string  `json:"room-type"`
	Price    float32 `json:"price,string"`
}

type HotelData struct {
	Inventory []Hotel `json:"inventory"`
}

type HotelPrice float32

type UUID string

type Product struct {
	ID   UUID
	Name string
}

func NewHotel(src, name, roomtype string, price float32) Hotel {
	return Hotel{
		Source:   src,
		Name:     name,
		RoomType: roomtype,
		Price:    price,
	}
}

//function
func GetPrice(h Hotel) HotelPrice {
	return HotelPrice(h.Price)
}

//method with value receiver
func (h Hotel) GetPrice() HotelPrice {
	return HotelPrice(h.Price)
}

//function
func SetPrice(h Hotel, newPrice HotelPrice) {
	h.Price = float32(newPrice)
}

//method with value receiver
func (h Hotel) SetPrice(newPrice HotelPrice) {
	h.Price = float32(newPrice)
}

//function
func SetHotelPrice(h *Hotel, newPrice HotelPrice) {
	h.Price = float32(newPrice)
}

//method with pointer receiver
func (h *Hotel) SetHotelPrice(newPrice HotelPrice) {
	h.Price = float32(newPrice)
}

//[]Hotel ---->filter---> []Hotel

func FilterHotelByRoomType(hotelList []Hotel) []Hotel {
	return []Hotel{}
}

type Hotels []Hotel

type HotelFilterFn func(*Hotel) bool

func (hotels Hotels) FilterHotelByRoomTypeWithCopy(roomtype string) Hotels {
	var filteredList []Hotel

	//copy semantics
	for _, h := range hotels {
		if h.RoomType == roomtype {
			filteredList = append(filteredList, h)
		}
	}

	return filteredList
}

func (hotels Hotels) FilterHotelByRoomType(roomtype string) Hotels {
	var filteredList []Hotel

	//reference semantics: avoid copy
	for i := range hotels {
		if hotels[i].RoomType == roomtype {
			filteredList = append(filteredList, hotels[i])
		}
	}

	return filteredList
}

//higher order function as it is accepting func as a value
func (hotels Hotels) FilterByFn(fn func(h *Hotel) bool) Hotels {
	var filteredList []Hotel
	//reference semantics: avoid copy
	for i := range hotels {
		if fn(&hotels[i]) {
			filteredList = append(filteredList, hotels[i])
		}
	}
	return filteredList
}

//higher order function as it is accepting func as a value
func (hotels Hotels) FilterBy(fn HotelFilterFn) Hotels {
	var filteredList []Hotel
	//reference semantics: avoid copy
	for i := range hotels {
		if fn(&hotels[i]) {
			filteredList = append(filteredList, hotels[i])
		}
	}
	return filteredList
}
