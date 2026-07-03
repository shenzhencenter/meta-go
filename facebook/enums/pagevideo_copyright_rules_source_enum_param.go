package enums

type PagevideoCopyrightRulesSourceEnumParam string

const (
	PagevideoCopyrightRulesSourceEnumParamMatchSettingsDialog PagevideoCopyrightRulesSourceEnumParam = "MATCH_SETTINGS_DIALOG"
	PagevideoCopyrightRulesSourceEnumParamRulesSelector       PagevideoCopyrightRulesSourceEnumParam = "RULES_SELECTOR"
	PagevideoCopyrightRulesSourceEnumParamRulesTab            PagevideoCopyrightRulesSourceEnumParam = "RULES_TAB"
)

func (value PagevideoCopyrightRulesSourceEnumParam) String() string {
	return string(value)
}
