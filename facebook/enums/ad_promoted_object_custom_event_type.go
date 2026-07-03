package enums

type AdpromotedobjectCustomEventType string

const (
	AdpromotedobjectCustomEventTypeAchievementUnlocked             AdpromotedobjectCustomEventType = "ACHIEVEMENT_UNLOCKED"
	AdpromotedobjectCustomEventTypeAddPaymentInfo                  AdpromotedobjectCustomEventType = "ADD_PAYMENT_INFO"
	AdpromotedobjectCustomEventTypeAddToCart                       AdpromotedobjectCustomEventType = "ADD_TO_CART"
	AdpromotedobjectCustomEventTypeAddToWishlist                   AdpromotedobjectCustomEventType = "ADD_TO_WISHLIST"
	AdpromotedobjectCustomEventTypeAdImpression                    AdpromotedobjectCustomEventType = "AD_IMPRESSION"
	AdpromotedobjectCustomEventTypeCompleteRegistration            AdpromotedobjectCustomEventType = "COMPLETE_REGISTRATION"
	AdpromotedobjectCustomEventTypeContact                         AdpromotedobjectCustomEventType = "CONTACT"
	AdpromotedobjectCustomEventTypeContentView                     AdpromotedobjectCustomEventType = "CONTENT_VIEW"
	AdpromotedobjectCustomEventTypeCustomizeProduct                AdpromotedobjectCustomEventType = "CUSTOMIZE_PRODUCT"
	AdpromotedobjectCustomEventTypeD2Retention                     AdpromotedobjectCustomEventType = "D2_RETENTION"
	AdpromotedobjectCustomEventTypeD7Retention                     AdpromotedobjectCustomEventType = "D7_RETENTION"
	AdpromotedobjectCustomEventTypeDonate                          AdpromotedobjectCustomEventType = "DONATE"
	AdpromotedobjectCustomEventTypeFindLocation                    AdpromotedobjectCustomEventType = "FIND_LOCATION"
	AdpromotedobjectCustomEventTypeInitiatedCheckout               AdpromotedobjectCustomEventType = "INITIATED_CHECKOUT"
	AdpromotedobjectCustomEventTypeLead                            AdpromotedobjectCustomEventType = "LEAD"
	AdpromotedobjectCustomEventTypeLevelAchieved                   AdpromotedobjectCustomEventType = "LEVEL_ACHIEVED"
	AdpromotedobjectCustomEventTypeListingInteraction              AdpromotedobjectCustomEventType = "LISTING_INTERACTION"
	AdpromotedobjectCustomEventTypeMessagingConversationStartedX7D AdpromotedobjectCustomEventType = "MESSAGING_CONVERSATION_STARTED_7D"
	AdpromotedobjectCustomEventTypeOther                           AdpromotedobjectCustomEventType = "OTHER"
	AdpromotedobjectCustomEventTypePurchase                        AdpromotedobjectCustomEventType = "PURCHASE"
	AdpromotedobjectCustomEventTypeRate                            AdpromotedobjectCustomEventType = "RATE"
	AdpromotedobjectCustomEventTypeSchedule                        AdpromotedobjectCustomEventType = "SCHEDULE"
	AdpromotedobjectCustomEventTypeSearch                          AdpromotedobjectCustomEventType = "SEARCH"
	AdpromotedobjectCustomEventTypeServiceBookingRequest           AdpromotedobjectCustomEventType = "SERVICE_BOOKING_REQUEST"
	AdpromotedobjectCustomEventTypeSpentCredits                    AdpromotedobjectCustomEventType = "SPENT_CREDITS"
	AdpromotedobjectCustomEventTypeStartTrial                      AdpromotedobjectCustomEventType = "START_TRIAL"
	AdpromotedobjectCustomEventTypeSubmitApplication               AdpromotedobjectCustomEventType = "SUBMIT_APPLICATION"
	AdpromotedobjectCustomEventTypeSubscribe                       AdpromotedobjectCustomEventType = "SUBSCRIBE"
	AdpromotedobjectCustomEventTypeTutorialCompletion              AdpromotedobjectCustomEventType = "TUTORIAL_COMPLETION"
)

func (value AdpromotedobjectCustomEventType) String() string {
	return string(value)
}
