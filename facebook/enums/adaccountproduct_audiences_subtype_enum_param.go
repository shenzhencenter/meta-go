package enums

type AdaccountproductAudiencesSubtypeEnumParam string

const (
	AdaccountproductAudiencesSubtypeEnumParamApp                         AdaccountproductAudiencesSubtypeEnumParam = "APP"
	AdaccountproductAudiencesSubtypeEnumParamBagOfAccounts               AdaccountproductAudiencesSubtypeEnumParam = "BAG_OF_ACCOUNTS"
	AdaccountproductAudiencesSubtypeEnumParamBidding                     AdaccountproductAudiencesSubtypeEnumParam = "BIDDING"
	AdaccountproductAudiencesSubtypeEnumParamClaim                       AdaccountproductAudiencesSubtypeEnumParam = "CLAIM"
	AdaccountproductAudiencesSubtypeEnumParamCustom                      AdaccountproductAudiencesSubtypeEnumParam = "CUSTOM"
	AdaccountproductAudiencesSubtypeEnumParamEngagement                  AdaccountproductAudiencesSubtypeEnumParam = "ENGAGEMENT"
	AdaccountproductAudiencesSubtypeEnumParamExclusion                   AdaccountproductAudiencesSubtypeEnumParam = "EXCLUSION"
	AdaccountproductAudiencesSubtypeEnumParamFox                         AdaccountproductAudiencesSubtypeEnumParam = "FOX"
	AdaccountproductAudiencesSubtypeEnumParamLookalike                   AdaccountproductAudiencesSubtypeEnumParam = "LOOKALIKE"
	AdaccountproductAudiencesSubtypeEnumParamManaged                     AdaccountproductAudiencesSubtypeEnumParam = "MANAGED"
	AdaccountproductAudiencesSubtypeEnumParamMeasurement                 AdaccountproductAudiencesSubtypeEnumParam = "MEASUREMENT"
	AdaccountproductAudiencesSubtypeEnumParamMessengerSubscriberList     AdaccountproductAudiencesSubtypeEnumParam = "MESSENGER_SUBSCRIBER_LIST"
	AdaccountproductAudiencesSubtypeEnumParamOfflineConversion           AdaccountproductAudiencesSubtypeEnumParam = "OFFLINE_CONVERSION"
	AdaccountproductAudiencesSubtypeEnumParamPartner                     AdaccountproductAudiencesSubtypeEnumParam = "PARTNER"
	AdaccountproductAudiencesSubtypeEnumParamPrimary                     AdaccountproductAudiencesSubtypeEnumParam = "PRIMARY"
	AdaccountproductAudiencesSubtypeEnumParamRegulatedCategoriesAudience AdaccountproductAudiencesSubtypeEnumParam = "REGULATED_CATEGORIES_AUDIENCE"
	AdaccountproductAudiencesSubtypeEnumParamStudyRuleAudience           AdaccountproductAudiencesSubtypeEnumParam = "STUDY_RULE_AUDIENCE"
	AdaccountproductAudiencesSubtypeEnumParamVideo                       AdaccountproductAudiencesSubtypeEnumParam = "VIDEO"
	AdaccountproductAudiencesSubtypeEnumParamWebsite                     AdaccountproductAudiencesSubtypeEnumParam = "WEBSITE"
)

func (value AdaccountproductAudiencesSubtypeEnumParam) String() string {
	return string(value)
}
