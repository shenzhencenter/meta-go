package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
)

type Event struct {
	AttendingCount           *int                          `json:"attending_count,omitempty"`
	CanGuestsInvite          *bool                         `json:"can_guests_invite,omitempty"`
	Category                 *enums.EventCategory          `json:"category,omitempty"`
	Cover                    *CoverPhoto                   `json:"cover,omitempty"`
	CreatedTime              *core.Time                    `json:"created_time,omitempty"`
	DeclinedCount            *int                          `json:"declined_count,omitempty"`
	Description              *string                       `json:"description,omitempty"`
	DiscountCodeEnabled      *bool                         `json:"discount_code_enabled,omitempty"`
	EndTime                  *string                       `json:"end_time,omitempty"`
	EventTimes               *[]ChildEvent                 `json:"event_times,omitempty"`
	GuestListEnabled         *bool                         `json:"guest_list_enabled,omitempty"`
	ID                       *core.ID                      `json:"id,omitempty"`
	InterestedCount          *int                          `json:"interested_count,omitempty"`
	IsCanceled               *bool                         `json:"is_canceled,omitempty"`
	IsDraft                  *bool                         `json:"is_draft,omitempty"`
	IsOnline                 *bool                         `json:"is_online,omitempty"`
	IsPageOwned              *bool                         `json:"is_page_owned,omitempty"`
	MaybeCount               *int                          `json:"maybe_count,omitempty"`
	Name                     *string                       `json:"name,omitempty"`
	NoreplyCount             *int                          `json:"noreply_count,omitempty"`
	OnlineEventFormat        *enums.EventOnlineEventFormat `json:"online_event_format,omitempty"`
	OnlineEventThirdPartyURL *string                       `json:"online_event_third_party_url,omitempty"`
	Owner                    *map[string]interface{}       `json:"owner,omitempty"`
	ParentGroup              *Group                        `json:"parent_group,omitempty"`
	Place                    *Place                        `json:"place,omitempty"`
	RegistrationSetting      *EventRegistrationSetting     `json:"registration_setting,omitempty"`
	ScheduledPublishTime     *string                       `json:"scheduled_publish_time,omitempty"`
	StartTime                *string                       `json:"start_time,omitempty"`
	TicketSetting            *EventTicketSetting           `json:"ticket_setting,omitempty"`
	TicketURI                *string                       `json:"ticket_uri,omitempty"`
	TicketURIStartSalesTime  *string                       `json:"ticket_uri_start_sales_time,omitempty"`
	TicketingPrivacyURI      *string                       `json:"ticketing_privacy_uri,omitempty"`
	TicketingTermsURI        *string                       `json:"ticketing_terms_uri,omitempty"`
	Timezone                 *string                       `json:"timezone,omitempty"`
	Type                     *enums.EventType              `json:"type,omitempty"`
	UpdatedTime              *core.Time                    `json:"updated_time,omitempty"`
}

var EventFields = struct {
	AttendingCount           string
	CanGuestsInvite          string
	Category                 string
	Cover                    string
	CreatedTime              string
	DeclinedCount            string
	Description              string
	DiscountCodeEnabled      string
	EndTime                  string
	EventTimes               string
	GuestListEnabled         string
	ID                       string
	InterestedCount          string
	IsCanceled               string
	IsDraft                  string
	IsOnline                 string
	IsPageOwned              string
	MaybeCount               string
	Name                     string
	NoreplyCount             string
	OnlineEventFormat        string
	OnlineEventThirdPartyURL string
	Owner                    string
	ParentGroup              string
	Place                    string
	RegistrationSetting      string
	ScheduledPublishTime     string
	StartTime                string
	TicketSetting            string
	TicketURI                string
	TicketURIStartSalesTime  string
	TicketingPrivacyURI      string
	TicketingTermsURI        string
	Timezone                 string
	Type                     string
	UpdatedTime              string
}{
	AttendingCount:           "attending_count",
	CanGuestsInvite:          "can_guests_invite",
	Category:                 "category",
	Cover:                    "cover",
	CreatedTime:              "created_time",
	DeclinedCount:            "declined_count",
	Description:              "description",
	DiscountCodeEnabled:      "discount_code_enabled",
	EndTime:                  "end_time",
	EventTimes:               "event_times",
	GuestListEnabled:         "guest_list_enabled",
	ID:                       "id",
	InterestedCount:          "interested_count",
	IsCanceled:               "is_canceled",
	IsDraft:                  "is_draft",
	IsOnline:                 "is_online",
	IsPageOwned:              "is_page_owned",
	MaybeCount:               "maybe_count",
	Name:                     "name",
	NoreplyCount:             "noreply_count",
	OnlineEventFormat:        "online_event_format",
	OnlineEventThirdPartyURL: "online_event_third_party_url",
	Owner:                    "owner",
	ParentGroup:              "parent_group",
	Place:                    "place",
	RegistrationSetting:      "registration_setting",
	ScheduledPublishTime:     "scheduled_publish_time",
	StartTime:                "start_time",
	TicketSetting:            "ticket_setting",
	TicketURI:                "ticket_uri",
	TicketURIStartSalesTime:  "ticket_uri_start_sales_time",
	TicketingPrivacyURI:      "ticketing_privacy_uri",
	TicketingTermsURI:        "ticketing_terms_uri",
	Timezone:                 "timezone",
	Type:                     "type",
	UpdatedTime:              "updated_time",
}
