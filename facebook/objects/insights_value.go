package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
	"time"
)

type InsightsValue struct {
	CampaignID                       *core.ID                `json:"campaign_id,omitempty"`
	EarningSource                    *string                 `json:"earning_source,omitempty"`
	EndTime                          *time.Time              `json:"end_time,omitempty"`
	EngagementSource                 *string                 `json:"engagement_source,omitempty"`
	IsFromAds                        *string                 `json:"is_from_ads,omitempty"`
	IsFromFollowers                  *string                 `json:"is_from_followers,omitempty"`
	MessageType                      *string                 `json:"message_type,omitempty"`
	MessagingChannel                 *string                 `json:"messaging_channel,omitempty"`
	MonetizationTool                 *string                 `json:"monetization_tool,omitempty"`
	RecurringNotificationsEntryPoint *string                 `json:"recurring_notifications_entry_point,omitempty"`
	RecurringNotificationsFrequency  *string                 `json:"recurring_notifications_frequency,omitempty"`
	RecurringNotificationsTopic      *string                 `json:"recurring_notifications_topic,omitempty"`
	StartTime                        *time.Time              `json:"start_time,omitempty"`
	Value                            *map[string]interface{} `json:"value,omitempty"`
}
