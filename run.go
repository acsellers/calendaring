package main

import "github.com/acsellers/calendars/ews"

func main() {
	c := ews.Conn{
		Username: "user",
		Password: "password",
		Host:     "https://outlook.office365.com/EWS/Exchange.asmx",
	}
	c.Do()
}
