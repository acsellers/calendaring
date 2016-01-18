package main

import (
	"encoding/xml"
	"fmt"

	"github.com/acsellers/calendars/ews"
)

func main() {
	RunDecode()
	// RunConn()
}
func RunDecode() {
	xc := `<?xml version="1.0" encoding="utf-8"?>
	<s:Envelope xmlns:s="http://schemas.xmlsoap.org/soap/envelope/">
		<s:Header>
			<h:ServerVersionInfo MajorVersion="15" MinorVersion="1" MajorBuildNumber="365" MinorBuildNumber="23" xmlns:h="http://schemas.microsoft.com/exchange/services/2006/types" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"/>
		</s:Header>
		<s:Body>
			<m:GetFolderResponse xmlns:m="http://schemas.microsoft.com/exchange/services/2006/messages" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:t="http://schemas.microsoft.com/exchange/services/2006/types">
				<m:ResponseMessages>
					<m:GetFolderResponseMessage ResponseClass="Success">
						<m:ResponseCode>NoError</m:ResponseCode>
						<m:Folders>
							<t:CalendarFolder>
								<t:FolderId Id="AQAlAGFuZHJld0BzZWxsAGVyc3Rlc3Rpbmcub25taWNyb3NvZnQuY29tAC4AAAOU9jeFA+BAT4n/IFlBwlr6AQBNGAZEpZQuRJ/1KOxKSzREAAACAQ0AAAA=" ChangeKey="AgAAABYAAABNGAZEpZQuRJ/1KOxKSzREAAAAAAA3"/>
								<t:DisplayName>Calendar</t:DisplayName>
								<t:TotalCount>3</t:TotalCount>
								<t:ChildFolderCount>2</t:ChildFolderCount>
							</t:CalendarFolder>
						</m:Folders>
					</m:GetFolderResponseMessage>
				</m:ResponseMessages>
			</m:GetFolderResponse>
		</s:Body>
	</s:Envelope>
	`
	se := ews.SoapResponse{}
	ffr := ews.GetFolderResponse{}
	fmt.Println(xml.Unmarshal([]byte(xc), &se))
	fmt.Println(xml.Unmarshal(se.Body.Data, &ffr))
	fmt.Println(len(ffr.Messages.Message.Calendars))
	fmt.Println(ffr.Messages.Message.Calendars[0].DisplayName)
}

func RunConn() {
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
