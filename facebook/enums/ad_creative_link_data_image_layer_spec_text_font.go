package enums

type AdcreativelinkdataimagelayerspecTextFont string

const (
	AdcreativelinkdataimagelayerspecTextFontDroidSerifRegular      AdcreativelinkdataimagelayerspecTextFont = "droid_serif_regular"
	AdcreativelinkdataimagelayerspecTextFontLatoRegular            AdcreativelinkdataimagelayerspecTextFont = "lato_regular"
	AdcreativelinkdataimagelayerspecTextFontNotoSansRegular        AdcreativelinkdataimagelayerspecTextFont = "noto_sans_regular"
	AdcreativelinkdataimagelayerspecTextFontNunitoSansBold         AdcreativelinkdataimagelayerspecTextFont = "nunito_sans_bold"
	AdcreativelinkdataimagelayerspecTextFontOpenSansBold           AdcreativelinkdataimagelayerspecTextFont = "open_sans_bold"
	AdcreativelinkdataimagelayerspecTextFontOpenSansCondensedBold  AdcreativelinkdataimagelayerspecTextFont = "open_sans_condensed_bold"
	AdcreativelinkdataimagelayerspecTextFontPtSerifBold            AdcreativelinkdataimagelayerspecTextFont = "pt_serif_bold"
	AdcreativelinkdataimagelayerspecTextFontRobotoCondensedRegular AdcreativelinkdataimagelayerspecTextFont = "roboto_condensed_regular"
	AdcreativelinkdataimagelayerspecTextFontRobotoMedium           AdcreativelinkdataimagelayerspecTextFont = "roboto_medium"
)

func (value AdcreativelinkdataimagelayerspecTextFont) String() string {
	return string(value)
}
