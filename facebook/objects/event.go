package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
	"github.com/shenzhencenter/meta-go/facebook/enums"
	"time"
)

type Event struct {
	AttendingCount           *int                          `json:"attending_count,omitempty"`
	CanGuestsInvite          *bool                         `json:"can_guests_invite,omitempty"`
	Category                 *enums.EventCategory          `json:"category,omitempty"`
	Cover                    *CoverPhoto                   `json:"cover,omitempty"`
	CreatedTime              *time.Time                    `json:"created_time,omitempty"`
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
	UpdatedTime              *time.Time                    `json:"updated_time,omitempty"`
}
