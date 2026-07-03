package objects

type LeadGenAppointmentBookingInfo struct {
	AdvertiserTimezoneOffset *string                         `json:"advertiser_timezone_offset,omitempty"`
	AppointmentDurations     *[]string                       `json:"appointment_durations,omitempty"`
	AppointmentSlotsByDay    *[]LeadGenAppointmentSlotsByDay `json:"appointment_slots_by_day,omitempty"`
}
