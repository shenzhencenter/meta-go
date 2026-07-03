package objects

type LeadGenAppointmentBookingInfo struct {
	AdvertiserTimezoneOffset *string                         `json:"advertiser_timezone_offset,omitempty"`
	AppointmentDurations     *[]string                       `json:"appointment_durations,omitempty"`
	AppointmentSlotsByDay    *[]LeadGenAppointmentSlotsByDay `json:"appointment_slots_by_day,omitempty"`
}

var LeadGenAppointmentBookingInfoFields = struct {
	AdvertiserTimezoneOffset string
	AppointmentDurations     string
	AppointmentSlotsByDay    string
}{
	AdvertiserTimezoneOffset: "advertiser_timezone_offset",
	AppointmentDurations:     "appointment_durations",
	AppointmentSlotsByDay:    "appointment_slots_by_day",
}
