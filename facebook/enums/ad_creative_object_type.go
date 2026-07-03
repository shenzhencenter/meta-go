package enums

type AdcreativeObjectType string

const (
	AdcreativeObjectTypeApplication      AdcreativeObjectType = "APPLICATION"
	AdcreativeObjectTypeDomain           AdcreativeObjectType = "DOMAIN"
	AdcreativeObjectTypeEvent            AdcreativeObjectType = "EVENT"
	AdcreativeObjectTypeInvalid          AdcreativeObjectType = "INVALID"
	AdcreativeObjectTypeOffer            AdcreativeObjectType = "OFFER"
	AdcreativeObjectTypePage             AdcreativeObjectType = "PAGE"
	AdcreativeObjectTypePhoto            AdcreativeObjectType = "PHOTO"
	AdcreativeObjectTypePostDeleted      AdcreativeObjectType = "POST_DELETED"
	AdcreativeObjectTypePrivacyCheckFail AdcreativeObjectType = "PRIVACY_CHECK_FAIL"
	AdcreativeObjectTypeShare            AdcreativeObjectType = "SHARE"
	AdcreativeObjectTypeStatus           AdcreativeObjectType = "STATUS"
	AdcreativeObjectTypeStoreItem        AdcreativeObjectType = "STORE_ITEM"
	AdcreativeObjectTypeVideo            AdcreativeObjectType = "VIDEO"
)

func (value AdcreativeObjectType) String() string {
	return string(value)
}
