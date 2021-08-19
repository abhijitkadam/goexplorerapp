package main

import (
	"encoding/json"
	"net/http"
)

type ApiHotelRepo struct {
	ApiUrl string
}

func NewApiHotelRepo(url string) ApiHotelRepo {
	return ApiHotelRepo{
		ApiUrl: url,
	}
}

func (hr *ApiHotelRepo) GetAllHotels() []Hotel {
	client := http.Client{}

	response, err := client.Get(hr.ApiUrl)

	//if not nil then there is some error
	if err != nil {
		//unhappy path, just return in this case. Should also log the error.
		return []Hotel{}
	}

	defer response.Body.Close()

	var hotels []Hotel

	json.NewDecoder(response.Body).Decode(&hotels)

	return hotels
}
