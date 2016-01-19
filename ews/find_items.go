package ews

import (
	"encoding/xml"
	"time"
)

func (c *Conn) FindCalendarItems() (*FindItemResponse, error) {
	r := FindItemCalendarRequest{}
	r.Traversal = "Shallow"
	r.Xmlns = "http://schemas.microsoft.com/exchange/services/2006/messages"
	r.XmlnsT = "http://schemas.microsoft.com/exchange/services/2006/types"
	r.Shape.BaseShape = "IdOnly"
	r.ParentFolderIds.Folder.Id = "calendar"
	r.CalendarView.MaxEntriesReturned = 10
	r.CalendarView.StartDate = time.Now().AddDate(0, -1, 0).Format(time.RFC3339)
	r.CalendarView.EndDate = time.Now().AddDate(0, 3, 0).Format(time.RFC3339)

	resp := &FindItemResponse{}
	err := c.Do(r, resp)
	return resp, err
}

type FindItemCalendarRequest struct {
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

type FindItemResponse struct {
	XMLName  xml.Name `xml:"FindItemResponse"`
	XmlnsT   string   `xml:"xmlns:t,attr"`
	Messages struct {
		XMLName xml.Name `xml:"ResponseMessages"`
		Message struct {
			XMLName      xml.Name `xml:"FindItemResponseMessage"`
			ResponseCode string   `xml:"ResponseCode"`
			RootFolder   struct {
				XMLName                 xml.Name    `xml:"RootFolder"`
				TotalItemsInView        string      `xml:"TotalItemsInView,attr"`
				IncludesLastItemInRange string      `xml:"IncludesLastItemInRange,attr"`
				Items                   []EWSItemId `xml:"Items>CalendarItem>ItemId"`
			}
		}
	}
}
