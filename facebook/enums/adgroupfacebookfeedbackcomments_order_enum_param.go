package enums

type AdgroupfacebookfeedbackcommentsOrderEnumParam string

const (
	AdgroupfacebookfeedbackcommentsOrderEnumParamChronological        AdgroupfacebookfeedbackcommentsOrderEnumParam = "chronological"
	AdgroupfacebookfeedbackcommentsOrderEnumParamReverseChronological AdgroupfacebookfeedbackcommentsOrderEnumParam = "reverse_chronological"
)

func (value AdgroupfacebookfeedbackcommentsOrderEnumParam) String() string {
	return string(value)
}
