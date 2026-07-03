package enums

type AdstudyobjectiveBreakdowns string

const (
	AdstudyobjectiveBreakdownsAge     AdstudyobjectiveBreakdowns = "age"
	AdstudyobjectiveBreakdownsCellID  AdstudyobjectiveBreakdowns = "cell_id"
	AdstudyobjectiveBreakdownsCountry AdstudyobjectiveBreakdowns = "country"
	AdstudyobjectiveBreakdownsGender  AdstudyobjectiveBreakdowns = "gender"
)

func (value AdstudyobjectiveBreakdowns) String() string {
	return string(value)
}
