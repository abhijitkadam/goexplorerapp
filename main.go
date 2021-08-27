package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func HomeHandler(rw http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(rw, "This is go server")
}

func main() {
	Config()

	http.HandleFunc("/home", HomeHandler)

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
