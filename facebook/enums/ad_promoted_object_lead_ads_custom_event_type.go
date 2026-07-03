package enums

type AdpromotedobjectLeadAdsCustomEventType string

const (
	AdpromotedobjectLeadAdsCustomEventTypeAchievementUnlocked             AdpromotedobjectLeadAdsCustomEventType = "ACHIEVEMENT_UNLOCKED"
	AdpromotedobjectLeadAdsCustomEventTypeAddPaymentInfo                  AdpromotedobjectLeadAdsCustomEventType = "ADD_PAYMENT_INFO"
	AdpromotedobjectLeadAdsCustomEventTypeAddToCart                       AdpromotedobjectLeadAdsCustomEventType = "ADD_TO_CART"
	AdpromotedobjectLeadAdsCustomEventTypeAddToWishlist                   AdpromotedobjectLeadAdsCustomEventType = "ADD_TO_WISHLIST"
	AdpromotedobjectLeadAdsCustomEventTypeAdImpression                    AdpromotedobjectLeadAdsCustomEventType = "AD_IMPRESSION"
	AdpromotedobjectLeadAdsCustomEventTypeCompleteRegistration            AdpromotedobjectLeadAdsCustomEventType = "COMPLETE_REGISTRATION"
	AdpromotedobjectLeadAdsCustomEventTypeContact                         AdpromotedobjectLeadAdsCustomEventType = "CONTACT"
	AdpromotedobjectLeadAdsCustomEventTypeContentView                     AdpromotedobjectLeadAdsCustomEventType = "CONTENT_VIEW"
	AdpromotedobjectLeadAdsCustomEventTypeCustomizeProduct                AdpromotedobjectLeadAdsCustomEventType = "CUSTOMIZE_PRODUCT"
	AdpromotedobjectLeadAdsCustomEventTypeD2Retention                     AdpromotedobjectLeadAdsCustomEventType = "D2_RETENTION"
	AdpromotedobjectLeadAdsCustomEventTypeD7Retention                     AdpromotedobjectLeadAdsCustomEventType = "D7_RETENTION"
	AdpromotedobjectLeadAdsCustomEventTypeDonate                          AdpromotedobjectLeadAdsCustomEventType = "DONATE"
	AdpromotedobjectLeadAdsCustomEventTypeFindLocation                    AdpromotedobjectLeadAdsCustomEventType = "FIND_LOCATION"
	AdpromotedobjectLeadAdsCustomEventTypeInitiatedCheckout               AdpromotedobjectLeadAdsCustomEventType = "INITIATED_CHECKOUT"
	AdpromotedobjectLeadAdsCustomEventTypeLead                            AdpromotedobjectLeadAdsCustomEventType = "LEAD"
	AdpromotedobjectLeadAdsCustomEventTypeLevelAchieved                   AdpromotedobjectLeadAdsCustomEventType = "LEVEL_ACHIEVED"
	AdpromotedobjectLeadAdsCustomEventTypeListingInteraction              AdpromotedobjectLeadAdsCustomEventType = "LISTING_INTERACTION"
	AdpromotedobjectLeadAdsCustomEventTypeMessagingConversationStartedX7D AdpromotedobjectLeadAdsCustomEventType = "MESSAGING_CONVERSATION_STARTED_7D"
	AdpromotedobjectLeadAdsCustomEventTypeOther                           AdpromotedobjectLeadAdsCustomEventType = "OTHER"
	AdpromotedobjectLeadAdsCustomEventTypePurchase                        AdpromotedobjectLeadAdsCustomEventType = "PURCHASE"
	AdpromotedobjectLeadAdsCustomEventTypeRate                            AdpromotedobjectLeadAdsCustomEventType = "RATE"
	AdpromotedobjectLeadAdsCustomEventTypeSchedule                        AdpromotedobjectLeadAdsCustomEventType = "SCHEDULE"
	AdpromotedobjectLeadAdsCustomEventTypeSearch                          AdpromotedobjectLeadAdsCustomEventType = "SEARCH"
	AdpromotedobjectLeadAdsCustomEventTypeServiceBookingRequest           AdpromotedobjectLeadAdsCustomEventType = "SERVICE_BOOKING_REQUEST"
	AdpromotedobjectLeadAdsCustomEventTypeSpentCredits                    AdpromotedobjectLeadAdsCustomEventType = "SPENT_CREDITS"
	AdpromotedobjectLeadAdsCustomEventTypeStartTrial                      AdpromotedobjectLeadAdsCustomEventType = "START_TRIAL"
	AdpromotedobjectLeadAdsCustomEventTypeSubmitApplication               AdpromotedobjectLeadAdsCustomEventType = "SUBMIT_APPLICATION"
	AdpromotedobjectLeadAdsCustomEventTypeSubscribe                       AdpromotedobjectLeadAdsCustomEventType = "SUBSCRIBE"
	AdpromotedobjectLeadAdsCustomEventTypeTutorialCompletion              AdpromotedobjectLeadAdsCustomEventType = "TUTORIAL_COMPLETION"
)

func (value AdpromotedobjectLeadAdsCustomEventType) String() string {
	return string(value)
}
