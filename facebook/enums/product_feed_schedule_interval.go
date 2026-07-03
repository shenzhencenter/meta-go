package enums

type ProductfeedscheduleInterval string

const (
	ProductfeedscheduleIntervalDaily   ProductfeedscheduleInterval = "DAILY"
	ProductfeedscheduleIntervalHourly  ProductfeedscheduleInterval = "HOURLY"
	ProductfeedscheduleIntervalMonthly ProductfeedscheduleInterval = "MONTHLY"
	ProductfeedscheduleIntervalWeekly  ProductfeedscheduleInterval = "WEEKLY"
)

func (value ProductfeedscheduleInterval) String() string {
	return string(value)
}
