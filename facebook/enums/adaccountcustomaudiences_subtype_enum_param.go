package enums

type AdaccountcustomaudiencesSubtypeEnumParam string

const (
	AdaccountcustomaudiencesSubtypeEnumParamApp                         AdaccountcustomaudiencesSubtypeEnumParam = "APP"
	AdaccountcustomaudiencesSubtypeEnumParamBagOfAccounts               AdaccountcustomaudiencesSubtypeEnumParam = "BAG_OF_ACCOUNTS"
	AdaccountcustomaudiencesSubtypeEnumParamBidding                     AdaccountcustomaudiencesSubtypeEnumParam = "BIDDING"
	AdaccountcustomaudiencesSubtypeEnumParamClaim                       AdaccountcustomaudiencesSubtypeEnumParam = "CLAIM"
	AdaccountcustomaudiencesSubtypeEnumParamCustom                      AdaccountcustomaudiencesSubtypeEnumParam = "CUSTOM"
	AdaccountcustomaudiencesSubtypeEnumParamEngagement                  AdaccountcustomaudiencesSubtypeEnumParam = "ENGAGEMENT"
	AdaccountcustomaudiencesSubtypeEnumParamExclusion                   AdaccountcustomaudiencesSubtypeEnumParam = "EXCLUSION"
	AdaccountcustomaudiencesSubtypeEnumParamFox                         AdaccountcustomaudiencesSubtypeEnumParam = "FOX"
	AdaccountcustomaudiencesSubtypeEnumParamLookalike                   AdaccountcustomaudiencesSubtypeEnumParam = "LOOKALIKE"
	AdaccountcustomaudiencesSubtypeEnumParamManaged                     AdaccountcustomaudiencesSubtypeEnumParam = "MANAGED"
	AdaccountcustomaudiencesSubtypeEnumParamMeasurement                 AdaccountcustomaudiencesSubtypeEnumParam = "MEASUREMENT"
	AdaccountcustomaudiencesSubtypeEnumParamMessengerSubscriberList     AdaccountcustomaudiencesSubtypeEnumParam = "MESSENGER_SUBSCRIBER_LIST"
	AdaccountcustomaudiencesSubtypeEnumParamOfflineConversion           AdaccountcustomaudiencesSubtypeEnumParam = "OFFLINE_CONVERSION"
	AdaccountcustomaudiencesSubtypeEnumParamPartner                     AdaccountcustomaudiencesSubtypeEnumParam = "PARTNER"
	AdaccountcustomaudiencesSubtypeEnumParamPrimary                     AdaccountcustomaudiencesSubtypeEnumParam = "PRIMARY"
	AdaccountcustomaudiencesSubtypeEnumParamRegulatedCategoriesAudience AdaccountcustomaudiencesSubtypeEnumParam = "REGULATED_CATEGORIES_AUDIENCE"
	AdaccountcustomaudiencesSubtypeEnumParamStudyRuleAudience           AdaccountcustomaudiencesSubtypeEnumParam = "STUDY_RULE_AUDIENCE"
	AdaccountcustomaudiencesSubtypeEnumParamVideo                       AdaccountcustomaudiencesSubtypeEnumParam = "VIDEO"
	AdaccountcustomaudiencesSubtypeEnumParamWebsite                     AdaccountcustomaudiencesSubtypeEnumParam = "WEBSITE"
)

func (value AdaccountcustomaudiencesSubtypeEnumParam) String() string {
	return string(value)
}
