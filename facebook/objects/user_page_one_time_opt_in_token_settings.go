package objects

import (
	core "github.com/shenzhencenter/facebook-go-sdk/facebook"
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
