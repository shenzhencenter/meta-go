package enums

type AdstudyobjectiveType string

const (
	AdstudyobjectiveTypeBrand         AdstudyobjectiveType = "BRAND"
	AdstudyobjectiveTypeBrandlift     AdstudyobjectiveType = "BRANDLIFT"
	AdstudyobjectiveTypeConversions   AdstudyobjectiveType = "CONVERSIONS"
	AdstudyobjectiveTypeFtl           AdstudyobjectiveType = "FTL"
	AdstudyobjectiveTypeMae           AdstudyobjectiveType = "MAE"
	AdstudyobjectiveTypeMai           AdstudyobjectiveType = "MAI"
	AdstudyobjectiveTypeMpcConversion AdstudyobjectiveType = "MPC_CONVERSION"
	AdstudyobjectiveTypeNonsales      AdstudyobjectiveType = "NONSALES"
	AdstudyobjectiveTypePartner       AdstudyobjectiveType = "PARTNER"
	AdstudyobjectiveTypeSales         AdstudyobjectiveType = "SALES"
	AdstudyobjectiveTypeTelco         AdstudyobjectiveType = "TELCO"
)

func (value AdstudyobjectiveType) String() string {
	return string(value)
}
