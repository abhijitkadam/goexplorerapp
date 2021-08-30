package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func HomeHandler(rw http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(rw, "This is go server")
}

type Customer struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	Config()

	http.HandleFunc("/home", HomeHandler)

	http.HandleFunc("/form", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Println("Method: ", r.Method)
		fmt.Println("URL: ", r.URL)

		//print query params:
		//fmt.Println("Query param: ", r.URL.Query())
		//get query param "choice"
		//fmt.Println("Query param: ", r.URL.Query().Get("choice"))

		r.ParseForm()
		fmt.Println("Form data: ", r.Form)
		fmt.Println("Skill: ", r.Form.Get("skills"))
		fmt.Println("Skills: ", r.Form["skills"])

		rw.WriteHeader(http.StatusAccepted)
		fmt.Fprintf(rw, "Form accepted")

	})

	http.HandleFunc("/body", func(rw http.ResponseWriter, r *http.Request) {
		bytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}
		defer r.Body.Close()
		fmt.Print(r.Header)
		fmt.Println(string(bytes))
		fmt.Fprintf(rw, "ok")
	})

	http.HandleFunc("/json", func(rw http.ResponseWriter, r *http.Request) {
		var cust Customer
		err := json.NewDecoder(r.Body).Decode(&cust)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusBadRequest)
			return
		}

		fmt.Println(cust)
	})

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
