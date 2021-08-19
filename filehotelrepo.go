package main

import (
	"encoding/json"
	"os"
)

type HotelRepoFile struct {
	fileName string
}

func NewHotelRepoFile(path string) HotelRepoFile {
	return HotelRepoFile{fileName: path}
}

func (hr *HotelRepoFile) GetAllHotels() []Hotel {
	file, err := os.Open(hr.fileName)
	if err != nil {
		return []Hotel{}
	}
	defer file.Close()

	var data HotelData
	err = json.NewDecoder(file).Decode(&data)
	if err != nil {
		return []Hotel{}
	}

	return data.Inventory
}
