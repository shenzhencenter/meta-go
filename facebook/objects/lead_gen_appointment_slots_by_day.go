package objects

type LeadGenAppointmentSlotsByDay struct {
	AppointmentSlots *[]LeadGenAppointmentTimeSlot `json:"appointment_slots,omitempty"`
	Day              *string                       `json:"day,omitempty"`
}
