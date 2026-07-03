package enums

type AdsnamingtemplateLevel string

const (
	AdsnamingtemplateLevelAdgroup           AdsnamingtemplateLevel = "ADGROUP"
	AdsnamingtemplateLevelAdAccount         AdsnamingtemplateLevel = "AD_ACCOUNT"
	AdsnamingtemplateLevelCampaign          AdsnamingtemplateLevel = "CAMPAIGN"
	AdsnamingtemplateLevelCampaignGroup     AdsnamingtemplateLevel = "CAMPAIGN_GROUP"
	AdsnamingtemplateLevelOpportunities     AdsnamingtemplateLevel = "OPPORTUNITIES"
	AdsnamingtemplateLevelPrivacyInfoCenter AdsnamingtemplateLevel = "PRIVACY_INFO_CENTER"
	AdsnamingtemplateLevelProduct           AdsnamingtemplateLevel = "PRODUCT"
	AdsnamingtemplateLevelTopline           AdsnamingtemplateLevel = "TOPLINE"
	AdsnamingtemplateLevelUniqueAdcreative  AdsnamingtemplateLevel = "UNIQUE_ADCREATIVE"
)

func (value AdsnamingtemplateLevel) String() string {
	return string(value)
}
