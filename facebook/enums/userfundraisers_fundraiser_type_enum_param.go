package enums

type UserfundraisersFundraiserTypeEnumParam string

const (
	UserfundraisersFundraiserTypeEnumParamPersonForCharity UserfundraisersFundraiserTypeEnumParam = "person_for_charity"
)

func (value UserfundraisersFundraiserTypeEnumParam) String() string {
	return string(value)
}
