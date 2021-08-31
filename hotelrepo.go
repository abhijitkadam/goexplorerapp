package main

import (
	"encoding/json"
	"os"
	"sync"
)

type HotelRepo interface {
	GetHotels() ([]Hotel, error)
	//GetHotel(code string) (Hotel, error)
}

type HotelJsonFileRepo struct {
	FileName string
	rwm      sync.RWMutex
}

func NewHotelJsonFileRepo(file string) HotelJsonFileRepo {
	return HotelJsonFileRepo{FileName: file, rwm: sync.RWMutex{}}
}

func (hfr *HotelJsonFileRepo) GetHotels() ([]Hotel, error) {
	hfr.rwm.RLock()
	defer hfr.rwm.RUnlock()

	f, err := os.OpenFile(hfr.FileName, os.O_RDONLY, 0755)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var data HotelData

	err = json.NewDecoder(f).Decode(&data)
	if err != nil {
		return nil, err
	}

	return data.Inventory, nil
}

//unimplemented for demo
func (hfr *HotelJsonFileRepo) UpdateHotels(data []Hotel) error {
	hfr.rwm.Lock()
	defer hfr.rwm.Unlock()

	//Write operations on file
	return nil
}
