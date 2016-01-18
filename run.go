package main

import (
	"fmt"

	"github.com/acsellers/calendars/ews"
)

func main() {
	c := ews.Conn{
		Username: "user",
		Password: "password",
		Host:     "https://outlook.office365.com/EWS/Exchange.asmx",
		Debug:    true,
	}
	fmt.Println(c.FindFolders())
	fmt.Println(c.GetFolder("calendar"))
	fmt.Println(c.FindItemsCalendar())
}
