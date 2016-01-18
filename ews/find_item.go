package ews

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"time"
)

func (c *Conn) FindItemsCalendar() error {
	r := FindItemCalendarReq{}
	r.Traversal = "Shallow"
	r.Xmlns = "http://schemas.microsoft.com/exchange/services/2006/messages"
	r.XmlnsT = "http://schemas.microsoft.com/exchange/services/2006/types"
	r.Shape.BaseShape = "IdOnly"
	r.ParentFolderIds.Folder.Id = "calendar"
	r.CalendarView.MaxEntriesReturned = 10
	r.CalendarView.StartDate = time.Now().AddDate(0, -1, 0).Format(time.RFC3339)
	r.CalendarView.EndDate = time.Now().AddDate(0, 3, 0).Format(time.RFC3339)

	body, err := c.Do(r)
	if err != nil {
		return err
	}
	b := &bytes.Buffer{}
	io.Copy(b, body)
	fmt.Println(b.String())
	return nil
}

type FindItemCalendarReq struct {
	XMLName   xml.Name `xml:"FindItem"`
	Xmlns     string   `xml:"xmlns,attr"`
	XmlnsT    string   `xml:"xmlns:t,attr"`
	Traversal string   `xml:"Traversal,attr"`

	Shape struct {
		XMLName   xml.Name `xml:"ItemShape"`
		BaseShape string   `xml:"t:BaseShape"`
	}
	CalendarView struct {
		XMLName            xml.Name `xml:"m:CalendarView"`
		MaxEntriesReturned int      `xml:"MaxEntriesReturned,attr"`
		StartDate          string   `xml:"StartDate,attr"`
		EndDate            string   `xml:"EndDate,attr"`
	}
	ParentFolderIds struct {
		XMLName xml.Name `xml:"ParentFolderIds"`
		Folder  struct {
			XMLName xml.Name `xml:"t:DistinguishedFolderId"`
			Id      string   `xml:"Id,attr"`
		}
	}
}
