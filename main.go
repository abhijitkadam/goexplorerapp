package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

//Middleware
func LoggerMiddleware(h http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		fmt.Println("Requesting ", r.URL)

		user := r.Context().Value("authuser")
		if user != nil {
			userName, ok := user.(string)
			if ok {
				fmt.Println("Authenticated & user is : ", userName)
			}
		}
		h(rw, r)
	}
}

func HomeHandler(rw http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(rw, "This is go server")
}

type Customer struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func JsonHandler(rw http.ResponseWriter, r *http.Request) {
	var cust Customer
	err := json.NewDecoder(r.Body).Decode(&cust)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Println(cust)
}

func main() {
	Config()

	store := NewHotelJsonFileRepo("data/golddata.json")

	//http.HandleFunc("/home", LoggerMiddleware(HomeHandler))
	//http.HandleFunc("/json", LoggerMiddleware(JsonHandler))

	r := mux.NewRouter()

	//without subrouting
	//SetRoutes(r)

	//Subrouter
	hotelsRouter := r.PathPrefix("/hotels").Subrouter()
	hotelsRouter.Use(MuxbasicAuthMiddleware, MuxloggingMiddleware)

	productsRouter := r.PathPrefix("/products").Subrouter()
	productsRouter.Use(MuxbasicAuthMiddleware, MuxloggingMiddleware)

	SetHotelsRoutes(hotelsRouter, &store)
	SetProductsRoutes(productsRouter)

	r.HandleFunc("/home", LoggerMiddleware(HomeHandler))
	r.HandleFunc("/json", LoggerMiddleware(JsonHandler))

	http.Handle("/", r)

	fmt.Println("Starting server : ", viper.Get("SERVER"), " on port ", viper.GetInt("port"))

	log.Fatal(http.ListenAndServe(":"+viper.GetString("port"), nil))

}

func SetRoutes(r *mux.Router) {
	r.HandleFunc("/hotels", LoggerMiddleware(func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw, "This is all the hotels data ...")
	})).Methods("GET")

	r.HandleFunc("/hotels/{code}", func(rw http.ResponseWriter, r *http.Request) {
		code, ok := mux.Vars(r)["code"]
		if !ok {
			http.Error(rw, "bad request, provide proper code", http.StatusBadRequest)
			return
		}

		fmt.Fprintf(rw, "This is data for hotel %s ...", code)
	}).Methods("GET")

	r.HandleFunc("/home", LoggerMiddleware(HomeHandler))
	r.HandleFunc("/json", LoggerMiddleware(JsonHandler))
}

func SetProductsRoutes(productsRouter *mux.Router) {
	productsRouter.HandleFunc("", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw, "This is all the products data ...")
	}).Methods("GET")

	productsRouter.HandleFunc("/{id}", func(rw http.ResponseWriter, r *http.Request) {
		id, ok := mux.Vars(r)["id"]
		if !ok {
			http.Error(rw, "bad request, provide proper code", http.StatusBadRequest)
			return
		}

		fmt.Fprintf(rw, "This is data for product %s ...", id)
	}).Methods("GET")
}

func Config() {

	flag.String("server", "appserver", "name of the server")

	flag.String("port", "8080", "port on which server runs")

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

	//override with env variables
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	//the env vars needs to be caps
	viper.AutomaticEnv()

}
