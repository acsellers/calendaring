package main

import (
	"encoding/xml"
	"fmt"
	"time"

	"github.com/acsellers/calendars/ews"
)

func main() {
	RunDecode()
	// RunConn()
}
func RunDecode() {
	xc := `<CalendarItem>
	<t:ItemId Id="AQAlAGFuZHJld0BzZWxsAGVyc3Rlc3Rpbmcub25taWNyb3NvZnQuY29tAEYAAAOU9jeFA+BAT4n/IFlBwlr6BwBNGAZEpZQuRJ/1KOxKSzREAAACAQ0AAABNGAZEpZQuRJ/1KOxKSzREAAACFPIAAAA=" ChangeKey="DwAAABYAAABNGAZEpZQuRJ/1KOxKSzREAAAAABVl"/>
	<t:ParentFolderId Id="AQAlAGFuZHJld0BzZWxsAGVyc3Rlc3Rpbmcub25taWNyb3NvZnQuY29tAC4AAAOU9jeFA+BAT4n/IFlBwlr6AQBNGAZEpZQuRJ/1KOxKSzREAAACAQ0AAAA=" ChangeKey="AQAAAA=="/>
	<t:ItemClass>IPM.Appointment</t:ItemClass>
	<t:Subject>Blah</t:Subject>
	<t:Sensitivity>Normal</t:Sensitivity>
	<t:Body BodyType="HTML"/>
	<t:DateTimeReceived>2016-01-14T05:29:20Z</t:DateTimeReceived>
	<t:Size>5755</t:Size>
	<t:Importance>Normal</t:Importance>
	<t:IsSubmitted>false</t:IsSubmitted>
	<t:IsDraft>false</t:IsDraft>
	<t:IsFromMe>false</t:IsFromMe>
	<t:IsResend>false</t:IsResend>
	<t:IsUnmodified>false</t:IsUnmodified>
	<t:DateTimeSent>2016-01-14T05:29:20Z</t:DateTimeSent>
	<t:DateTimeCreated>2016-01-14T05:29:20Z</t:DateTimeCreated>
	<t:ResponseObjects><t:ForwardItem/></t:ResponseObjects>
	<t:ReminderDueBy>2016-01-14T14:00:00Z</t:ReminderDueBy>
	<t:ReminderIsSet>true</t:ReminderIsSet>
	<t:ReminderMinutesBeforeStart>15</t:ReminderMinutesBeforeStart>
	<t:DisplayCc/>
	<t:DisplayTo/>
	<t:HasAttachments>false</t:HasAttachments>
	<t:Culture>en-US</t:Culture>
	<t:Start>2016-01-14T14:00:00Z</t:Start>
	<t:End>2016-01-14T14:30:00Z</t:End>
	<t:IsAllDayEvent>false</t:IsAllDayEvent>
	<t:LegacyFreeBusyStatus>Busy</t:LegacyFreeBusyStatus>
	<t:Location/>
	<t:IsMeeting>false</t:IsMeeting>
	<t:IsCancelled>false</t:IsCancelled>
	<t:IsRecurring>false</t:IsRecurring>
	<t:MeetingRequestWasSent>false</t:MeetingRequestWasSent>
	<t:IsResponseRequested>true</t:IsResponseRequested>
	<t:CalendarItemType>Single</t:CalendarItemType>
	<t:MyResponseType>Organizer</t:MyResponseType>
	<t:Organizer>
		<t:Mailbox>
			<t:Name>Andrew Sellers</t:Name>
			<t:EmailAddress>andrew@sellerstesting.onmicrosoft.com</t:EmailAddress>
			<t:RoutingType>SMTP</t:RoutingType>
		</t:Mailbox>
	</t:Organizer>
	<t:Duration>PT30M</t:Duration>
	<t:TimeZone>(UTC-06:00) Central Time (US &amp; Canada)</t:TimeZone>
	<t:AppointmentSequenceNumber>0</t:AppointmentSequenceNumber>
	<t:AppointmentState>0</t:AppointmentState>
	<t:IsOnlineMeeting>false</t:IsOnlineMeeting>
	</CalendarItem>
	`
	ffr := ews.EWSCalendarItem{}
	//fmt.Println(xml.Unmarshal([]byte(xc), &se))
	// fmt.Println(xml.Unmarshal(se.Body.Data, &ffr))
	fmt.Println(xml.Unmarshal([]byte(xc), &ffr))
	fmt.Println(ffr.Start)
	loc, _ := time.LoadLocation("America/Chicago")
	st, _ := time.Parse(time.RFC3339, ffr.Start)
	fmt.Println(st.In(loc).Format(time.ANSIC))
	// item.StartTime = st.In(loc)

	// fmt.Println(ffr.Messages.Message.RootFolder.Items[0].Id)
}

func RunConn() {
	c := ews.Conn{
		Username: "user",
		Password: "password",
		Host:     "https://outlook.office365.com/EWS/Exchange.asmx",
		Debug:    true,
	}
	// fmt.Println(c.FindFolders())
	// fmt.Println(c.GetFolder("calendar"))
	// fmt.Println(c.FindItemsCalendar())
	ciresp, err := c.FindCalendarItems()
	if err != nil {
		panic(err)
	}

	fmt.Println(ciresp.Messages.Message.RootFolder.Items)
	fmt.Println(ciresp.Messages.Message.RootFolder.TotalItemsInView)
	if len(ciresp.Messages.Message.RootFolder.Items) == 0 {
		panic("No items decoded")
	}
	resp, err := c.GetCalendarItem(
		ciresp.Messages.Message.RootFolder.Items,
		[]string{},
	)
	fmt.Println(err)
	fmt.Println(len(resp.Messages.Message.Items))
	// fmt.Println(string(resp.Messages.Message.Items[0].Data))
}
