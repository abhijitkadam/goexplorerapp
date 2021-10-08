package main

import (
	"flag"
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func HomeHandler(rw http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(rw, "This is go server")
}

func EchoHandler(c echo.Context) error {
	//1. Request
	//Binding is there: https://echo.labstack.com/guide/binding/
	//However if need be can ge the request like this:
	//c.Request()

	//2. ResponseWriter
	//There are helper methods
	//c.String(...)
	//c.JSON(..)
	//However if need be can ge the response writer
	//c.Response().Header().Add(..)
	//c.Response().WriteHeader(...)
	//fmt.Fprint(c.Response().Writer, "write it")

	//3 Context
	//c.Set("userinfo", user)
	//user, ok := c.Get("userinfo").(*User)

	//4. http.Error()
	//For echo in case of error there is a helper
	//return echo.NewHTTPError(http.StatusBadRequest, "Please provide valid input")
	//return echo.NewHTTPError(http.StatusUnauthorized, "Please provide valid credentials")

	return c.String(http.StatusOK, "Hello, World!")
}

///hotels/:name

type HotelCode struct {
	ID string `param:"id" query:"id" header:"id" form:"id" json:"id" xml:"id"`
}

// User
type User struct {
	Name  string `json:"name" form:"name" query:"name"`
	Email string `json:"email" form:"email" query:"email"`
}

func main() {
	Config()

	e := echo.New()

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			fmt.Println(c.Request().URL)
			return next(c)
		}
	})

	e.POST("/users", func(c echo.Context) (err error) {
		u := new(User)
		if err = c.Bind(u); err != nil {
			return
		}

		return c.JSON(http.StatusOK, u)
	})

	e.GET("/users", func(c echo.Context) (err error) {
		u := new(User)
		if err = c.Bind(u); err != nil {
			return
		}

		return c.JSON(http.StatusOK, u)
	})

	e.GET("/hotels/:id", func(c echo.Context) (err error) {
		//hc := new(HotelCode)
		// if err = c.Bind(hc); err != nil {
		// 	return
		// }

		code := c.Param("id")

		return c.JSON(http.StatusOK, code)
	})

	e.GET("/", EchoHandler)
	e.Logger.Fatal(e.Start(":" + viper.GetString("port")))

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
