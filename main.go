package main

import "explorerapp/services"

func main() {

	//Want to send email

	email := services.NewEmail("Abhijitk@gmail.com", "Golang Training", "Details on golang .....")
	services.SendEmail(email)

}
