package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func SetHotelsRoutes(hotelsRouter *mux.Router, store HotelRepo) {
	hotelsRouter.HandleFunc("", CreateGetHotelsHandler(store)).Methods("GET")

	hotelsRouter.HandleFunc("/{code}", GetHotelHandler).Methods("GET")
}

//Without common middlewares
// func SetHotelsRoutes(hotelsRouter *mux.Router, store HotelRepo) {
// 	hotelsRouter.HandleFunc("", BasicAuthMiddleware(LoggerMiddleware(CreateGetHotelsHandler(store)))).Methods("GET")

// 	hotelsRouter.HandleFunc("/{code}", BasicAuthMiddleware(LoggerMiddleware(GetHotelHandler))).Methods("GET")
// }

func CreateGetHotelsHandler(store HotelRepo) http.HandlerFunc {
	//GetHotelsHandler is using store abstraction
	return func(rw http.ResponseWriter, r *http.Request) {
		data, err := store.GetHotels()
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
		bytes, err := json.Marshal(data)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(rw, string(bytes))
	}
}

func GetHotelHandler(rw http.ResponseWriter, r *http.Request) {
	code, ok := mux.Vars(r)["code"]
	if !ok {
		http.Error(rw, "bad request, provide proper code", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(rw, "This is data for hotel %s ...", code)
}
