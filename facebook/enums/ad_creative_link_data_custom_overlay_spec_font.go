package enums

type AdcreativelinkdatacustomoverlayspecFont string

const (
	AdcreativelinkdatacustomoverlayspecFontDroidSerifRegular      AdcreativelinkdatacustomoverlayspecFont = "droid_serif_regular"
	AdcreativelinkdatacustomoverlayspecFontLatoRegular            AdcreativelinkdatacustomoverlayspecFont = "lato_regular"
	AdcreativelinkdatacustomoverlayspecFontNotoSansRegular        AdcreativelinkdatacustomoverlayspecFont = "noto_sans_regular"
	AdcreativelinkdatacustomoverlayspecFontNunitoSansBold         AdcreativelinkdatacustomoverlayspecFont = "nunito_sans_bold"
	AdcreativelinkdatacustomoverlayspecFontOpenSansBold           AdcreativelinkdatacustomoverlayspecFont = "open_sans_bold"
	AdcreativelinkdatacustomoverlayspecFontPtSerifBold            AdcreativelinkdatacustomoverlayspecFont = "pt_serif_bold"
	AdcreativelinkdatacustomoverlayspecFontRobotoCondensedRegular AdcreativelinkdatacustomoverlayspecFont = "roboto_condensed_regular"
	AdcreativelinkdatacustomoverlayspecFontRobotoMedium           AdcreativelinkdatacustomoverlayspecFont = "roboto_medium"
)

func (value AdcreativelinkdatacustomoverlayspecFont) String() string {
	return string(value)
}
