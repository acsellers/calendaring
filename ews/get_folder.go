package ews

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
)

func (c *Conn) GetFolder(name string) error {
	r := GetFolderReq{}
	r.Xmlns = "http://schemas.microsoft.com/exchange/services/2006/messages"

	r.Shape.BaseShape = "Default"
	r.ParentFolderIds.Folder.Id = "calendar"

	body, err := c.Do(r)
	if err != nil {
		return err
	}
	b := &bytes.Buffer{}
	io.Copy(b, body)
	fmt.Println(b.String())
	return nil
}

type GetFolderReq struct {
	XMLName xml.Name `xml:"GetFolder"`
	Xmlns   string   `xml:"xmlns,attr"`

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
