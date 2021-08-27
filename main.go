package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

//Built in flags package for input command line parsing
//A external lib: https://github.com/spf13/cobra
//https://pkg.go.dev/flag#example-Value

//Configuration stuff: https://github.com/spf13/viper
/*
func main() {

	//fmt.Println(os.Args)

	server := flag.String("servers", "appserver", "name of the server")

	//var servername string
	//flag.StringVar(&servername, "servers", "appserver", "name of the server")

	//port := flag.Int("port", 8080, "port on which server runs")

	flag.Parse()
	for _, v := range strings.Split(*server, ",") {
		fmt.Println(v)
	}
}
*/

//https://github.com/spf13/viper

func main() {

	fmt.Println(os.Getenv("SERVER"))

	flag.String("server", "appserver", "name of the server")

	flag.Int("port", 8080, "port on which server runs")

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

	//override with env variables
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))
	//the env vars needs to be caps
	viper.AutomaticEnv()

	fmt.Println("Starting server : ", viper.Get("SERVER"), " on port ", viper.GetInt("port"))
}

/*

{
    // Use IntelliSense to learn about possible attributes.
    // Hover to view descriptions of existing attributes.
    // For more information, visit: https://go.microsoft.com/fwlink/?linkid=830387
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Launch Package",
            "type": "go",
            "request": "launch",
            "mode": "auto",
            "program": "${workspaceFolder}",
            "envFile": "${workspaceFolder}/.env"
        }
    ]
}

*/
