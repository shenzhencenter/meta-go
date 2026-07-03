package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type InsightsValue struct {
	CampaignID                       *core.ID                `json:"campaign_id,omitempty"`
	EarningSource                    *string                 `json:"earning_source,omitempty"`
	EndTime                          *core.Time              `json:"end_time,omitempty"`
	EngagementSource                 *string                 `json:"engagement_source,omitempty"`
	IsFromAds                        *string                 `json:"is_from_ads,omitempty"`
	IsFromFollowers                  *string                 `json:"is_from_followers,omitempty"`
	MessageType                      *string                 `json:"message_type,omitempty"`
	MessagingChannel                 *string                 `json:"messaging_channel,omitempty"`
	MonetizationTool                 *string                 `json:"monetization_tool,omitempty"`
	RecurringNotificationsEntryPoint *string                 `json:"recurring_notifications_entry_point,omitempty"`
	RecurringNotificationsFrequency  *string                 `json:"recurring_notifications_frequency,omitempty"`
	RecurringNotificationsTopic      *string                 `json:"recurring_notifications_topic,omitempty"`
	StartTime                        *core.Time              `json:"start_time,omitempty"`
	Value                            *map[string]interface{} `json:"value,omitempty"`
}

var InsightsValueFields = struct {
	CampaignID                       string
	EarningSource                    string
	EndTime                          string
	EngagementSource                 string
	IsFromAds                        string
	IsFromFollowers                  string
	MessageType                      string
	MessagingChannel                 string
	MonetizationTool                 string
	RecurringNotificationsEntryPoint string
	RecurringNotificationsFrequency  string
	RecurringNotificationsTopic      string
	StartTime                        string
	Value                            string
}{
	CampaignID:                       "campaign_id",
	EarningSource:                    "earning_source",
	EndTime:                          "end_time",
	EngagementSource:                 "engagement_source",
	IsFromAds:                        "is_from_ads",
	IsFromFollowers:                  "is_from_followers",
	MessageType:                      "message_type",
	MessagingChannel:                 "messaging_channel",
	MonetizationTool:                 "monetization_tool",
	RecurringNotificationsEntryPoint: "recurring_notifications_entry_point",
	RecurringNotificationsFrequency:  "recurring_notifications_frequency",
	RecurringNotificationsTopic:      "recurring_notifications_topic",
	StartTime:                        "start_time",
	Value:                            "value",
}
