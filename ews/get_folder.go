package ews

import "encoding/xml"

func (c *Conn) GetFolder(name string) (*GetFolderResponse, error) {
	r := GetFolderRequest{}
	r.Xmlns = "http://schemas.microsoft.com/exchange/services/2006/messages"

	r.Shape.BaseShape = "Default"
	r.FolderIds.DistinguishedFolderIds = append(
		r.FolderIds.DistinguishedFolderIds,
		DistinguishedFolderId{Id: "calendar"},
	)

	resp := &GetFolderResponse{}
	err := c.Do(r, resp)
	return resp, err
}

type GetFolderRequest struct {
	XMLName xml.Name `xml:"m:GetFolder"`
	Xmlns   string   `xml:"xmlns,attr"`

	Shape struct {
		XMLName   xml.Name `xml:"m:FolderShape"`
		BaseShape string   `xml:"t:BaseShape"`
	}
	FolderIds struct {
		XMLName                xml.Name `xml:"m:FolderIds"`
		DistinguishedFolderIds []DistinguishedFolderId
	}
}

type DistinguishedFolderId struct {
	XMLName xml.Name `xml:"t:DistinguishedFolderId"`
	Id      string   `xml:"Id,attr"`
}

type GetFolderResponse struct {
	XMLName  xml.Name `xml:"GetFolderResponse"`
	Messages struct {
		XMLName       xml.Name `xml:"ResponseMessages"`
		ResponseClass string   `xml:"ResponseClass,attr"`
		Message       struct {
			XMLName       xml.Name          `xml:"GetFolderResponseMessage"`
			ResponseClass string            `xml:"ResponseClass,attr"`
			ResponseCode  string            `xml:"ResponseCode"`
			Calendars     []EWSCalendarInfo `xml:"Folders>CalendarFolder"`
		}
	}
}
