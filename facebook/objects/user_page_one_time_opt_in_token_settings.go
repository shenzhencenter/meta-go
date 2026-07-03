package objects

import (
	core "github.com/shenzhencenter/meta-go/facebook"
)

type UserPageOneTimeOptInTokenSettings struct {
	CreationTimestamp                *int       `json:"creation_timestamp,omitempty"`
	CustomAudienceIds                *[]core.ID `json:"custom_audience_ids,omitempty"`
	NextEligibleTime                 *int       `json:"next_eligible_time,omitempty"`
	NextEligibleTimeForPaidMessaging *int       `json:"next_eligible_time_for_paid_messaging,omitempty"`
	NotificationMessagesFrequency    *string    `json:"notification_messages_frequency,omitempty"`
	NotificationMessagesReoptin      *string    `json:"notification_messages_reoptin,omitempty"`
	NotificationMessagesTimezone     *string    `json:"notification_messages_timezone,omitempty"`
	NotificationMessagesToken        *string    `json:"notification_messages_token,omitempty"`
	RecipientID                      *core.ID   `json:"recipient_id,omitempty"`
	TokenExpiryTimestamp             *int       `json:"token_expiry_timestamp,omitempty"`
	TopicTitle                       *string    `json:"topic_title,omitempty"`
	UserTokenStatus                  *string    `json:"user_token_status,omitempty"`
}

var UserPageOneTimeOptInTokenSettingsFields = struct {
	CreationTimestamp                string
	CustomAudienceIds                string
	NextEligibleTime                 string
	NextEligibleTimeForPaidMessaging string
	NotificationMessagesFrequency    string
	NotificationMessagesReoptin      string
	NotificationMessagesTimezone     string
	NotificationMessagesToken        string
	RecipientID                      string
	TokenExpiryTimestamp             string
	TopicTitle                       string
	UserTokenStatus                  string
}{
	CreationTimestamp:                "creation_timestamp",
	CustomAudienceIds:                "custom_audience_ids",
	NextEligibleTime:                 "next_eligible_time",
	NextEligibleTimeForPaidMessaging: "next_eligible_time_for_paid_messaging",
	NotificationMessagesFrequency:    "notification_messages_frequency",
	NotificationMessagesReoptin:      "notification_messages_reoptin",
	NotificationMessagesTimezone:     "notification_messages_timezone",
	NotificationMessagesToken:        "notification_messages_token",
	RecipientID:                      "recipient_id",
	TokenExpiryTimestamp:             "token_expiry_timestamp",
	TopicTitle:                       "topic_title",
	UserTokenStatus:                  "user_token_status",
}
