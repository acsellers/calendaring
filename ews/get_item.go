package ews

import (
	"encoding/xml"
	"time"
)

func (c *Conn) GetCalendarItem(items []EWSItemId, fields []string) (*GetCalendarItemResponse, error) {
	r := GetCalendarItemRequest{}
	r.Xmlns = "http://schemas.microsoft.com/exchange/services/2006/messages"

	r.ItemShape.BaseShape = "AllProperties"
	/*
		for _, field := range fields {
			r.ItemShape.AdditionalProperties.Properies = append(
				r.ItemShape.AdditionalProperties.Properies,
				ItemProperties{URI: field},
			)
		}
	*/

	for _, item := range items {
		r.ItemIds.Items = append(
			r.ItemIds.Items,
			item.ToRequestId(),
		)
	}

	resp := &GetCalendarItemResponse{}
	err := c.Do(r, resp)
	loc, _ := time.LoadLocation("America/Chicago")
	for i, item := range resp.Messages.Message.Items {
		st, _ := time.Parse(time.RFC3339, item.Start)
		item.StartTime = st.In(loc)
		et, _ := time.Parse(time.RFC3339, item.End)
		item.EndTime = et.In(loc)

		resp.Messages.Message.Items[i] = item
	}

	return resp, err
}

type GetCalendarItemRequest struct {
	XMLName xml.Name `xml:"m:GetItem"`
	Xmlns   string   `xml:"xmlns,attr"`

	ItemShape struct {
		XMLName   xml.Name `xml:"m:ItemShape"`
		BaseShape string   `xml:"t:BaseShape"`
		/*AdditionalProperties struct {
			XMLName   xml.Name `xml:"AdditionalProperties"`
			Properies []ItemProperties
		}*/
	}
	ItemIds struct {
		XMLName xml.Name `xml:"m:ItemIds"`
		Items   []EWSReqItemId
	}
}

type ItemProperties struct {
	XMLName xml.Name `xml:"FieldURI"`
	URI     string   `xml:"FieldURI,attr"`
}

type GetCalendarItemResponse struct {
	XMLName  xml.Name `xml:"GetItemResponse"`
	XmlnsT   string   `xml:"xmlns:t,attr"`
	Messages struct {
		XMLName       xml.Name `xml:"ResponseMessages"`
		ResponseClass string   `xml:"ResponseClass,attr"`
		Message       struct {
			XMLName       xml.Name          `xml:"GetItemResponseMessage"`
			ResponseClass string            `xml:"ResponseClass,attr"`
			ResponseCode  string            `xml:"ResponseCode"`
			Items         []EWSCalendarItem `xml:"Items>CalendarItem"`
		}
	}
}

type EWSCalendarItem struct {
	XMLName xml.Name `xml:"CalendarItem"`
	EWSItemId
	ParentFolder struct {
		XMLName   xml.Name `xml:"ParentFolderId"`
		Id        string   `xml:"Id,attr"`
		ChangeKey string   `xml:"ChangeKey,attr"`
	}
	ItemClass string `xml:"ItemClass"`
	Subject   string `xml:"Subject"`
	// Sensitivity string `xml:"Sensitivity"`
	// Body ???
	// DateTimeReceived time.Time
	// Size int
	// Importance string
	// IsSubmitted bool
	// IsDraft bool
	// IsFromMe bool
	// IsResend bool
	// IsUnmodified bool
	// DateTimeSent time.Time
	// DateTimeCreated time.Time
	// ResponseObjects ???
	// ReminderDueBy time.Time
	// ReminderIsSet bool
	// ReminderMinutesBeforeStart int
	// DisplayCC ???
	// DisplayTo ???
	// HasAttachments bool
	// Culture string
	Start                string `xml:"Start"`
	StartTime            time.Time
	End                  string `xml:"End"`
	EndTime              time.Time
	IsAllDayEvent        bool   `xml:"IsAllDayEvent"`
	LegacyFreeBusyStatus string `xml:"LegacyFreeBusyStatus"`
	// Location ??
	// IsMeeting bool
	// IsCancelled bool
	// IsRecurring bool
	// MeetingRequestWasSent bool
	// IsResponseRequested bool
	CalendarItemType string `xml:"CalendarItemType"`
	// MyResponseType string
	// Organizer struct { Mailbox struct {
	// - Name string
	// - EmailAddress string
	// - RoutingType string
	Duration string `xml:"Duration"`
	TimeZone string `xml:"TimeZone"`
	// AppointmentSequenceNumber int
	// AppointmentState int
	// IsOnlineMeeting bool
}
