package ews

import "encoding/xml"

func (c *Conn) FindFolders() (*FindFolderResponse, error) {
	r := FindFolderReq{}
	r.Xmlns = "http://schemas.microsoft.com/exchange/services/2006/messages"
	r.Traversal = "Shallow"

	r.Shape.BaseShape = "Default"
	r.ParentFolderIds.Folder.Id = "calendar"

	resp := &FindFolderResponse{}
	err := c.Do(r, resp)
	return resp, err
}

type FindFolderReq struct {
	XMLName   xml.Name `xml:"FindFolder"`
	Xmlns     string   `xml:"xmlns,attr"`
	Traversal string   `xml:"Traversal,attr"`

	Shape struct {
		XMLName   xml.Name `xml:"FolderShape"`
		BaseShape string   `xml:"t:BaseShape"`
	}
	ParentFolderIds struct {
		XMLName xml.Name `xml:"ParentFolderIds"`
		Folder  struct {
			XMLName xml.Name `xml:"t:DistinguishedFolderId"`
			Id      string   `xml:"Id,attr"`
		}
	}
}

type FindFolderResponse struct {
	XMLName  xml.Name `xml:"FindFolderResponse"`
	Messages struct {
		XMLName xml.Name `xml:"ResponseMessages"`
		Message struct {
			XMLName      xml.Name `xml:"FindFolderResponseMessage"`
			ResponseCode string   `xml:"ResponseCode"`
			RootFolder   struct {
				XMLName                 xml.Name          `xml:"RootFolder"`
				TotalItemsInView        string            `xml:"TotalItemsInView,attr"`
				IncludesLastItemInRange string            `xml:"IncludesLastItemInRange,attr"`
				Calendars               []EWSCalendarInfo `xml:"Folders>CalendarFolder"`
			}
		}
	}
}
