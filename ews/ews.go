package ews

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Conn struct {
	Username string
	Password string
	Host     string
	Debug    bool
}

func (c Conn) Auth() string {
	return c.Username + ":" + c.Password
}

func (c *Conn) Do(body interface{}, response interface{}) error {
	buf := &bytes.Buffer{}
	io.WriteString(buf, xml.Header)
	enc := xml.NewEncoder(buf)
	if c.Debug {
		enc.Indent("  ", "  ")
	}

	envelope := SoapEnvelope{
		XmlnsXsi:  "http://www.w3.org/2001/XMLSchema-instance",
		XmlnsXsd:  "http://www.w3.org/2001/XMLSchema",
		XmlnsSoap: "http://schemas.xmlsoap.org/soap/envelope/",
		XmlnsT:    "http://schemas.microsoft.com/exchange/services/2006/types",
		XmlnsM:    "http://schemas.microsoft.com/exchange/services/2006/messages",
	}
	envelope.Body.Request = body
	enc.Encode(envelope)

	if c.Debug {
		fmt.Println(buf.String())
	}
	req, _ := http.NewRequest(
		"POST",
		strings.TrimSpace(c.Host),
		buf,
	)
	req.Header.Add("Content-Type", "text/xml; charset=utf-8")
	req.SetBasicAuth(c.Username, c.Password)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	se := &SoapResponse{}
	err = xml.NewDecoder(resp.Body).Decode(se)
	if err != nil {
		return err
	}
	if response == nil || true {
		fmt.Println(string(se.Body.Data))
		return nil
	}
	return xml.Unmarshal(se.Body.Data, response)
}

type SoapEnvelope struct {
	XMLName   xml.Name `xml:"soap:Envelope"`
	XmlnsXsi  string   `xml:"xmlns:xsi,attr"`
	XmlnsXsd  string   `xml:"xmlns:xsd,attr"`
	XmlnsSoap string   `xml:"xmlns:soap,attr"`
	XmlnsT    string   `xml:"xmlns:t,attr"`
	XmlnsM    string   `xml:"xmlns:m,attr"`
	Body      struct {
		XMLName xml.Name `xml:"soap:Body"`
		Request interface{}
	}
}

type SoapResponse struct {
	XMLName   xml.Name `xml:"Envelope"`
	XmlnsXsi  string   `xml:"xmlns:xsi,attr"`
	XmlnsXsd  string   `xml:"xmlns:xsd,attr"`
	XmlnsSoap string   `xml:"xmlns:soap,attr"`
	XmlnsT    string   `xml:"xmlns:t,attr"`
	XmlnsM    string   `xml:"xmlns:m,attr"`
	Header    struct {
		XMLName           xml.Name `xml:"Header"`
		ServerVersionInfo struct {
			XMLName          xml.Name `xml:"ServerVersionInfo"`
			MajorVersion     string   `xml:"MajorVersion,attr"`
			MinorVersion     string   `xml:"MinorVersion,attr"`
			MajorBuildNumber string   `xml:"MajorBuildNumber,attr"`
			MinorBuildNumber string   `xml:"MinorBuildNumber,attr"`
			XmlnsH           string   `xml:"xmlns:h,attr"`
			XmlnsXsd         string   `xml:"xmlns:xsd,attr"`
			XmlnsXsi         string   `xml:"xmlns:xsi,attr"`
		}
	}
	Body struct {
		XMLName xml.Name `xml:"Body"`
		Data    []byte   `xml:",innerxml"`
	}
}

type EWSCalendarInfo struct {
	FolderId struct {
		XMLName   xml.Name `xml:"FolderId"`
		Id        string   `xml:"Id,attr"`
		ChangeKey string   `xml:"ChangeKey,attr"`
	}
	DisplayName      string `xml:"DisplayName"`
	TotalCount       int    `xml:"TotalCount"`
	ChildFolderCount int    `xml:"ChildFolderCount"`
}
