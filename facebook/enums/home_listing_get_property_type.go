package enums

type HomelistinggetPropertyType string

const (
	HomelistinggetPropertyTypeApartment          HomelistinggetPropertyType = "APARTMENT"
	HomelistinggetPropertyTypeApartmentRoom      HomelistinggetPropertyType = "APARTMENT_ROOM"
	HomelistinggetPropertyTypeBuilderFloor       HomelistinggetPropertyType = "BUILDER_FLOOR"
	HomelistinggetPropertyTypeBungalow           HomelistinggetPropertyType = "BUNGALOW"
	HomelistinggetPropertyTypeCondo              HomelistinggetPropertyType = "CONDO"
	HomelistinggetPropertyTypeCondoRoom          HomelistinggetPropertyType = "CONDO_ROOM"
	HomelistinggetPropertyTypeEmptyValue         HomelistinggetPropertyType = "EMPTY_VALUE"
	HomelistinggetPropertyTypeHouse              HomelistinggetPropertyType = "HOUSE"
	HomelistinggetPropertyTypeHouseInCondominium HomelistinggetPropertyType = "HOUSE_IN_CONDOMINIUM"
	HomelistinggetPropertyTypeHouseInVilla       HomelistinggetPropertyType = "HOUSE_IN_VILLA"
	HomelistinggetPropertyTypeHouseRoom          HomelistinggetPropertyType = "HOUSE_ROOM"
	HomelistinggetPropertyTypeLand               HomelistinggetPropertyType = "LAND"
	HomelistinggetPropertyTypeLoft               HomelistinggetPropertyType = "LOFT"
	HomelistinggetPropertyTypeManufactured       HomelistinggetPropertyType = "MANUFACTURED"
	HomelistinggetPropertyTypeOther              HomelistinggetPropertyType = "OTHER"
	HomelistinggetPropertyTypeOtherRoom          HomelistinggetPropertyType = "OTHER_ROOM"
	HomelistinggetPropertyTypePenthouse          HomelistinggetPropertyType = "PENTHOUSE"
	HomelistinggetPropertyTypeSingleFamilyHome   HomelistinggetPropertyType = "SINGLE_FAMILY_HOME"
	HomelistinggetPropertyTypeStudio             HomelistinggetPropertyType = "STUDIO"
	HomelistinggetPropertyTypeTownhouse          HomelistinggetPropertyType = "TOWNHOUSE"
	HomelistinggetPropertyTypeTownhouseRoom      HomelistinggetPropertyType = "TOWNHOUSE_ROOM"
)

func (value HomelistinggetPropertyType) String() string {
	return string(value)
}
