package objects

type LeadGenAppointmentSlotsByDay struct {
	AppointmentSlots *[]LeadGenAppointmentTimeSlot `json:"appointment_slots,omitempty"`
	Day              *string                       `json:"day,omitempty"`
}

var LeadGenAppointmentSlotsByDayFields = struct {
	AppointmentSlots string
	Day              string
}{
	AppointmentSlots: "appointment_slots",
	Day:              "day",
}
