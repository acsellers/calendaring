package ews

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"os/exec"
	"strings"
)

type Conn struct {
	Username string
	Password string
	Host     string
}

func (c Conn) Auth() string {
	return c.Username + ":" + c.Password
}

func (c *Conn) Do() {
	cmd := exec.Command(
		"curl",
		"-u", strings.TrimSpace(string(c.Auth())),
		"--data",
		"@-",
		"-H",
		"Content-Type: text/xml; charset=utf-8",
		strings.TrimSpace(string(c.Host)),
	)

	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Fatal(err)
	}

	if err := cmd.Start(); err != nil {
		log.Fatal("Couldn't start command", err)
	}

	go func() {
		writeFindItem(stdin)
	}()

	buf := bytes.Buffer{}
	io.Copy(&buf, stdout)
	fmt.Println(buf.String())

	buf2 := bytes.Buffer{}
	io.Copy(&buf2, stderr)
	fmt.Println(buf2.String())

	if err := cmd.Wait(); err != nil {
		log.Fatal("Failed on wait", err)
	}
}

type Envelope struct {
	XMLName   xml.Name `xml:"soap:Envelope"`
	XmlnsXsi  string   `xml:"xmlns:xsi,attr"`
	XmlnsXsd  string   `xml:"xmlns:xsd,attr"`
	XmlnsSoap string   `xml:"xmlns:soap,attr"`
	XmlnsT    string   `xml:"xmlns:t,attr"`
	XmlnsM    string   `xml:"xmlns:m,attr"`
	Body      EnvelopeBody
}

type EnvelopeBody struct {
	XMLName xml.Name `xml:"soap:Body"`
	Request Request
}

type Request struct {
	XMLName   xml.Name `xml:"FindItem"`
	Xmlns     string   `xml:"xmlns,attr"`
	XmlnsT    string   `xml:"xmlns:t,attr"`
	Traversal string   `xml:"Traversal,attr"`

	Shape struct {
		XMLName   xml.Name `xml:"ItemShape"`
		BaseShape string   `xml:"t:BaseShape"`
	}
	ParentFolderIds struct {
		XMLName xml.Name `xml:"ParentFolderIds"`
		Folder  struct {
			XMLName xml.Name `xml:"t:DistinguishedFolderId"`
			Id      string   `xml:"Id,attr"`
		}
	}
	Properties string `xml:"m:UserConfigurationProperties"`
}

func writeFindItem(w io.WriteCloser) {
	defer w.Close()
	buf := bytes.Buffer{}
	io.WriteString(&buf, xml.Header)
	enc := xml.NewEncoder(&buf)
	enc.Indent("  ", "  ")
	r := Request{}
	r.Traversal = "Shallow"
	r.Xmlns = "http://schemas.microsoft.com/exchange/services/2006/messages"
	r.XmlnsT = "http://schemas.microsoft.com/exchange/services/2006/types"
	r.Shape.BaseShape = "IdOnly"
	r.ParentFolderIds.Folder.Id = "deleteditems"

	enc.Encode(
		Envelope{
			XmlnsXsi:  "http://www.w3.org/2001/XMLSchema-instance",
			XmlnsXsd:  "http://www.w3.org/2001/XMLSchema",
			XmlnsSoap: "http://schemas.xmlsoap.org/soap/envelope/",
			XmlnsT:    "http://schemas.microsoft.com/exchange/services/2006/types",
			XmlnsM:    "http://schemas.microsoft.com/exchange/services/2006/messages",
			Body: EnvelopeBody{
				Request: r,
			},
		},
	)
	fmt.Println(buf.String())
	io.Copy(w, &buf)
}
