package enums

type CustomaudiencedatasourceType string

const (
	CustomaudiencedatasourceTypeContactImporter    CustomaudiencedatasourceType = "CONTACT_IMPORTER"
	CustomaudiencedatasourceTypeCopyPaste          CustomaudiencedatasourceType = "COPY_PASTE"
	CustomaudiencedatasourceTypeEventBased         CustomaudiencedatasourceType = "EVENT_BASED"
	CustomaudiencedatasourceTypeFileImported       CustomaudiencedatasourceType = "FILE_IMPORTED"
	CustomaudiencedatasourceTypeHouseholdAudience  CustomaudiencedatasourceType = "HOUSEHOLD_AUDIENCE"
	CustomaudiencedatasourceTypeSeedBased          CustomaudiencedatasourceType = "SEED_BASED"
	CustomaudiencedatasourceTypeThirdPartyImported CustomaudiencedatasourceType = "THIRD_PARTY_IMPORTED"
	CustomaudiencedatasourceTypeUnknown            CustomaudiencedatasourceType = "UNKNOWN"
)

func (value CustomaudiencedatasourceType) String() string {
	return string(value)
}
