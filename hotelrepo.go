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

type HotelRepo interface {
	GetAllHotels() []Hotel
}
