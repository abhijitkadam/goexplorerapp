package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

//Middleware
func LoggerMiddleware(h http.HandlerFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		fmt.Println("Requesting ", r.URL)
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

	http.HandleFunc("/home", LoggerMiddleware(HomeHandler))

	http.HandleFunc("/json", LoggerMiddleware(JsonHandler))

	fmt.Println("Starting server : ", viper.Get("SERVER"), " on port ", viper.GetInt("port"))

	log.Fatal(http.ListenAndServe(":"+viper.GetString("port"), nil))

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
